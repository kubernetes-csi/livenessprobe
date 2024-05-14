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
	"time"

	"k8s.io/klog/v2"

	"k8s.io/component-base/featuregate"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"
	_ "k8s.io/component-base/logs/json/register"

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
	logger := klog.FromContext(ctx)
	defer cancel()

	conn, err := connlib.Connect(ctx, *csiAddress, h.metricsManager, connlib.WithTimeout(*probeTimeout))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		logger.Error(err, "Failed to establish connection to CSI driver")
		return
	}
	defer conn.Close()

	logger.V(5).Info("Sending probe request to CSI driver", "driver", h.driverName)
	ready, err := rpc.Probe(ctx, conn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		logger.Error(err, "Health check failed")
		return
	}

	if !ready {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("driver responded but is not ready"))
		logger.Error(nil, "Driver responded but is not ready")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`ok`))
	logger.V(5).Info("Health check succeeded")
}

func main() {
	fg := featuregate.NewFeatureGate()
	logsapi.AddFeatureGates(fg)
	c := logsapi.NewLoggingConfiguration()
	logsapi.AddGoFlags(c, flag.CommandLine)
	logs.InitLogs()
	flag.Parse()
	logger := klog.Background()
	if err := logsapi.ValidateAndApply(c, fg); err != nil {
		logger.Error(err, "LoggingConfiguration is invalid")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	if *healthzPort != defaultHealthzPort && *httpEndpoint != "" {
		logger.Error(nil, "Only one of `--health-port` and `--http-endpoint` can be explicitly set")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	if *metricsAddress != "" && *httpEndpoint != "" {
		logger.Error(nil, "Only one of `--metrics-address` and `--http-endpoint` can be explicitly set")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	var addr string
	if *httpEndpoint != "" {
		addr = *httpEndpoint
	} else {
		addr = net.JoinHostPort("0.0.0.0", *healthzPort)
	}

	metricsManager := metrics.NewCSIMetricsManager("" /* driverName */)
	// Connect to the CSI driver without any timeout to avoid crashing the probe when the driver is not ready yet.
	// Goal: liveness probe never crashes, it only fails the probe when the driver is not available (yet).
	// Since a http server for the probe is not running at this point, Kubernetes liveness probe will fail immediately
	// with "connection refused", which is good enough to fail the probe.
	ctx := context.Background()
	csiConn, err := connlib.Connect(ctx, *csiAddress, metricsManager, connlib.WithTimeout(0))
	if err != nil {
		// connlib should retry forever so a returned error should mean
		// the grpc client is misconfigured rather than an error on the network or CSI driver.
		logger.Error(err, "Failed to establish connection to CSI driver")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	logger.Info("Calling CSI driver to discover driver name")
	csiDriverName, err := rpc.GetDriverName(context.Background(), csiConn)
	csiConn.Close()
	if err != nil {
		// The CSI driver does not support GetDriverName, which is serious enough to crash the probe.
		logger.Error(err, "Failed to get CSI driver name")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	logger.Info("CSI driver name", "driver", csiDriverName)

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
			logger.Info("Separate metrics ServeMux listening", "address", *metricsAddress)
			err := http.ListenAndServe(*metricsAddress, metricsMux)
			if err != nil {
				logger.Error(err, "Failed to start prometheus metrics endpoint on specified address and path", "addr", *metricsAddress, "path", *metricsPath)
				klog.FlushAndExit(klog.ExitFlushTimeout, 1)
			}
		}()
	}

	mux.HandleFunc("/healthz", hp.checkProbe)
	logger.Info("ServeMux listening", "address", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Error(err, "Failed to start http server")
	}
}
