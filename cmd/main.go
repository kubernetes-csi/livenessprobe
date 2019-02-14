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

	"github.com/kubernetes-csi/livenessprobe/pkg/connection"
)

const (
	// Default timeout of short CSI calls like GetPluginInfo
	csiTimeout = time.Second
)

// Command line flags
var (
	// kubeconfig        = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file. Required only when running out of cluster.")
	connectionTimeout = flag.Duration("connection-timeout", 0, "The --connection-timeout flag is deprecated")
	csiAddress        = flag.String("csi-address", "/run/csi/socket", "Address of the CSI driver socket.")
	healthzPort       = flag.String("health-port", "9808", "TCP ports for listening healthz requests")
)

func runProbe(ctx context.Context, csiConn connection.CSIConnection) error {
	// Get CSI driver name.
	klog.Infof("Calling CSI driver to discover driver name.")
	csiDriverName, err := csiConn.GetDriverName(ctx)
	if err != nil {
		return err
	}
	klog.Infof("CSI driver name: %q", csiDriverName)
	// Sending Probe request
	klog.Infof("Sending probe request to CSI driver.")
	err = csiConn.LivenessProbe(ctx)
	return err
}

func getCSIConnection() (connection.CSIConnection, error) {
	klog.Infof("Attempting to open a gRPC connection with: %s", *csiAddress)
	return connection.NewConnection(*csiAddress)
}

func checkHealth(w http.ResponseWriter, req *http.Request) {
	klog.Infof("Request: %s from: %s\n", req.URL.Path, req.RemoteAddr)
	csiConn, err := getCSIConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		klog.Errorf("Failed to get connection to CSI  with error: %v.", err)
		return
	}
	defer csiConn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := runProbe(ctx, csiConn); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		klog.Errorf("Health check failed with: %v.", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`ok`))
		klog.Infof("Health check succeeded.")
	}
}

func main() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "true")
	flag.Parse()

	if *connectionTimeout != 0 {
		klog.Warning("--connection-timeout is deprecated and will have no effect")
	}

	addr := net.JoinHostPort("0.0.0.0", *healthzPort)
	http.HandleFunc("/healthz", checkHealth)
	klog.Infof("Serving requests to /healthz on: %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		klog.Fatalf("failed to start http server with error: %v", err)
	}
}
