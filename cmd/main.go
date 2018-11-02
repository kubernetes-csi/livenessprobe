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

	"github.com/golang/glog"

	"github.com/kubernetes-csi/livenessprobe/pkg/connection"
)

const (
	// Default timeout of short CSI calls like GetPluginInfo
	csiTimeout = time.Second
)

// Command line flags
var (
	// kubeconfig        = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file. Required only when running out of cluster.")
	connectionTimeout = flag.Duration("connection-timeout", 30*time.Second, "Timeout for waiting for CSI driver socket in seconds.")
	csiAddress        = flag.String("csi-address", "/run/csi/socket", "Address of the CSI driver socket.")
	healthzPort       = flag.String("health-port", "9808", "TCP ports for listening healthz requests")
)

func runProbe(ctx context.Context, csiConn connection.CSIConnection) error {
	// Get CSI driver name.
	glog.Infof("Calling CSI driver to discover driver name.")
	csiDriverName, err := csiConn.GetDriverName(ctx)
	if err != nil {
		return err
	}
	glog.Infof("CSI driver name: %q", csiDriverName)
	// Sending Probe request
	glog.Infof("Sending probe request to CSI driver.")
	if err := csiConn.LivenessProbe(ctx); err != nil {
		return err
	}
	return nil
}

func getCSIConnection() (connection.CSIConnection, error) {
	// Connect to CSI.
	glog.Infof("Attempting to open a gRPC connection with: %s", *csiAddress)
	csiConn, err := connection.NewConnection(*csiAddress, *connectionTimeout)
	if err != nil {
		return nil, err
	}
	return csiConn, nil
}

func chekcHealth(w http.ResponseWriter, req *http.Request) {

	glog.Infof("Request: %s from: %s\n", req.URL.Path, req.RemoteAddr)
	csiConn, err := getCSIConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		glog.Infof("Failed to get connection to CSI  with error: %v.", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), *connectionTimeout)
	defer cancel()
	if err := runProbe(ctx, csiConn); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		glog.Infof("Health check failed with: %v.", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`ok`))
		glog.Infof("Health check succeeded.")
	}
}

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	addr := net.JoinHostPort("0.0.0.0", *healthzPort)
	http.HandleFunc("/healthz", chekcHealth)
	glog.Infof("Serving requests to /healthz on: %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		glog.Fatalf("failed to start http server with error: %v", err)
	}
}
