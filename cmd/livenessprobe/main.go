/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"k8s.io/klog/v2"

	connlib "github.com/kubernetes-csi/csi-lib-utils/connection"
	"github.com/kubernetes-csi/csi-lib-utils/metrics"
	"github.com/kubernetes-csi/csi-lib-utils/rpc"
)

const (
	defaultHealthzPort = "9808"
)

// Command line flags
var (
	probeTimeout   = flag.Duration("probe-timeout", time.Second, "Probe timeout in seconds.")
	csiAddress     = flag.String("csi-address", "/run/csi/socket", "Address of the CSI driver socket.")
	healthzPort    = flag.String("health-port", defaultHealthzPort, fmt.Sprintf("(deprecated) TCP ports for listening healthz requests. The default is `%s`. If set, `--http-endpoint` cannot be set.", defaultHealthzPort))
	metricsAddress = flag.String("metrics-address", "", "(deprecated) The TCP network address where the prometheus metrics endpoint will listen (example: `:8080`). The default is empty string, which means metrics endpoint is disabled. If set, `--http-endpoint` cannot be set, and the address cannot resolve to localhost + the port from `--health-port`.")
	httpEndpoint   = flag.String("http-endpoint", "", "The TCP network address where the HTTP server for diagnostics, including CSI driver health check and metrics. The default is empty string, which means the server is disabled. If set, `--health-port` and `--metrics-address` cannot be explicitly set.")
	metricsPath    = flag.String("metrics-path", "/metrics", "The HTTP path where prometheus metrics will be exposed. Default is `/metrics`.")
)

type healthProbe struct {
	driverName     string
	metricsManager metrics.CSIMetricsManager
}

func (h *healthProbe) checkProbe(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), *probeTimeout)
	defer cancel()

	conn, err := acquireConnection(ctx, h.metricsManager)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		klog.Errorf("failed to establish connection to CSI driver: %v", err)
		return
	}
	defer conn.Close()

	klog.V(5).Infof("Sending probe request to CSI driver %q", h.driverName)
	ready, err := rpc.Probe(ctx, conn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		klog.Errorf("health check failed: %v", err)
		return
	}

	if !ready {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("driver responded but is not ready"))
		klog.Error("driver responded but is not ready")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`ok`))
	klog.V(5).Infof("Health check succeeded")
}

// acquireConnection wraps the connlib.Connect but adding support to context
// cancelation.
func acquireConnection(ctx context.Context, metricsManager metrics.CSIMetricsManager) (conn *grpc.ClientConn, err error) {

	var m sync.Mutex
	var canceled bool
	ready := make(chan bool)
	go func() {
		conn, err = connlib.Connect(*csiAddress, metricsManager)

		m.Lock()
		defer m.Unlock()
		if err != nil && canceled {
			conn.Close()
		}

		close(ready)
	}()

	select {
	case <-ctx.Done():
		m.Lock()
		defer m.Unlock()
		canceled = true
		return nil, ctx.Err()

	case <-ready:
		return conn, err
	}
}

func main() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "true")
	flag.Parse()

	if *healthzPort != defaultHealthzPort && *httpEndpoint != "" {
		klog.Error("only one of `--health-port` and `--http-endpoint` can be explicitly set.")
		os.Exit(1)
	}
	if *metricsAddress != "" && *httpEndpoint != "" {
		klog.Error("only one of `--metrics-address` and `--http-endpoint` can be explicitly set.")
		os.Exit(1)
	}
	var addr string
	if *httpEndpoint != "" {
		addr = *httpEndpoint
	} else {
		addr = net.JoinHostPort("0.0.0.0", *healthzPort)
	}

	metricsManager := metrics.NewCSIMetricsManager("" /* driverName */)
	csiConn, err := acquireConnection(context.Background(), metricsManager)
	if err != nil {
		// connlib should retry forever so a returned error should mean
		// the grpc client is misconfigured rather than an error on the network
		klog.Fatalf("failed to establish connection to CSI driver: %v", err)
	}

	klog.Infof("calling CSI driver to discover driver name")
	csiDriverName, err := rpc.GetDriverName(context.Background(), csiConn)
	csiConn.Close()
	if err != nil {
		klog.Fatalf("failed to get CSI driver name: %v", err)
	}
	klog.Infof("CSI driver name: %q", csiDriverName)

	hp := &healthProbe{
		driverName:     csiDriverName,
		metricsManager: metricsManager,
	}

	mux := http.NewServeMux()
	metricsManager.SetDriverName(csiDriverName)

	if *metricsAddress == "" {
		if *httpEndpoint != "" {
			metricsManager.RegisterToServer(mux, *metricsPath)
		}
	} else {
		// Remove once --metrics-address is removed
		metricsMux := http.NewServeMux()
		metricsManager.RegisterToServer(metricsMux, *metricsPath)
		go func() {
			klog.Infof("Separate metrics ServeMux listening at %q", *metricsAddress)
			err := http.ListenAndServe(*metricsAddress, metricsMux)
			if err != nil {
				klog.Fatalf("Failed to start prometheus metrics endpoint on specified address (%q) and path (%q): %s", *metricsAddress, *metricsPath, err)
			}
		}()
	}

	mux.HandleFunc("/healthz", hp.checkProbe)
	klog.Infof("ServeMux listening at %q", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		klog.Fatalf("failed to start http server with error: %v", err)
	}
}
