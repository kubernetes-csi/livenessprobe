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
	"time"

	"k8s.io/klog"

	connlib "github.com/kubernetes-csi/csi-lib-utils/connection"
	"github.com/kubernetes-csi/csi-lib-utils/rpc"
	"google.golang.org/grpc"
)

// Command line flags
var (
	// kubeconfig        = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file. Required only when running out of cluster.")
	connectionTimeout = flag.Duration("connection-timeout", 0, "The --connection-timeout flag is deprecated")
	probeTimeout      = flag.Duration("probe-timeout", time.Second, "Probe timeout in seconds")
	csiAddress        = flag.String("csi-address", "/run/csi/socket", "Address of the CSI driver socket.")
	healthzPort       = flag.String("health-port", "9808", "TCP ports for listening healthz requests")
)

type healthProbe struct {
	conn       *grpc.ClientConn
	driverName string
}

func (h *healthProbe) checkProbe(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), *probeTimeout)
	defer cancel()

	klog.Infof("Sending probe request to CSI driver %q", h.driverName)
	ready, err := rpc.Probe(ctx, h.conn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		klog.Errorf("health check failed: %v", err)
		return
	}

	if !ready {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		klog.Error("driver responded but is not ready")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`ok`))
	klog.Infof("Health check succeeded")
}

func main() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "true")
	flag.Parse()

	if *connectionTimeout != 0 {
		klog.Warning("--connection-timeout is deprecated and will have no effect")
	}

	csiConn, err := connlib.Connect(*csiAddress)
	if err != nil {
		// connlib should retry forever so a returned error should mean
		// the grpc client is misconfigured rather than an error on the network
		klog.Fatalf("failed to establish connection to CSI driver: %v", err)
	}

	klog.Infof("calling CSI driver to discover driver name")
	csiDriverName, err := rpc.GetDriverName(context.Background(), csiConn)
	if err != nil {
		klog.Fatalf("failed to get CSI driver name: %v", err)
	}
	klog.Infof("CSI driver name: %q", csiDriverName)

	hp := &healthProbe{
		conn:       csiConn,
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
