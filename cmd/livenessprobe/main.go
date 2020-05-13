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
	"net"
	"net/http"
	"sync"
	"time"

	"k8s.io/klog"

	connlib "github.com/kubernetes-csi/csi-lib-utils/connection"
	"github.com/kubernetes-csi/csi-lib-utils/rpc"

	"google.golang.org/grpc"
)

// Command line flags
var (
	probeTimeout   = flag.Duration("probe-timeout", time.Second, "Probe timeout in seconds")
	connectTimeout = flag.Duration("connect-timeout", 0, "Limit time to establish a connection to CSI driver")
	csiAddress     = flag.String("csi-address", "/run/csi/socket", "Address of the CSI driver socket.")
	healthzPort    = flag.String("health-port", "9808", "TCP ports for listening healthz requests")
)

type healthProbe struct {
	driverName string
}

func (h *healthProbe) checkProbe(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), *probeTimeout)
	defer cancel()

	conn, err := acquireConnection(ctx)
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
func acquireConnection(ctx context.Context) (conn *grpc.ClientConn, err error) {
	var m sync.Mutex
	var canceled bool
	ready := make(chan bool)
	go func() {
		conn, err = connlib.Connect(*csiAddress)

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

	csiConn, err := acquireConnection(context.Background())
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
		driverName: csiDriverName,
	}

	addr := net.JoinHostPort("0.0.0.0", *healthzPort)
	http.HandleFunc("/healthz", hp.checkProbe)
	klog.Infof("Serving requests to /healthz on: %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		klog.Fatalf("failed to start http server with error: %v", err)
	}
}
