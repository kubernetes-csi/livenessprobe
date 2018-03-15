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
	"os"
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
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	// Connect to CSI.
	glog.V(1).Infof("Attempting to open a gRPC connection with: %q", csiAddress)
	csiConn, err := connection.NewConnection(*csiAddress, *connectionTimeout)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}

	// Get CSI driver name.
	glog.V(1).Infof("Calling CSI driver to discover driver name.")
	ctx, cancel := context.WithTimeout(context.Background(), csiTimeout)
	defer cancel()
	csiDriverName, err := csiConn.GetDriverName(ctx)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}
	glog.V(2).Infof("CSI driver name: %q", csiDriverName)

	// Get CSI Driver Node ID
	glog.V(1).Infof("Calling CSI driver to discover node ID.")
	ctx, cancel = context.WithTimeout(context.Background(), csiTimeout)
	defer cancel()
	csiDriverNodeID, err := csiConn.NodeGetId(ctx)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}
	glog.V(2).Infof("CSI driver node ID: %q", csiDriverNodeID)

	// Sending Probe request
	glog.V(1).Infof("Sending probe request to CSI driver.")
	ctx, cancel = context.WithTimeout(context.Background(), csiTimeout)
	defer cancel()
	err = csiConn.LivenessProbe(ctx)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}
	// Liveness Probe was sucessfuly
	os.Exit(0)
}
