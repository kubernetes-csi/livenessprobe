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
	"net/http"
	"net/http/httptest"
	"testing"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/mock/gomock"
	connlib "github.com/kubernetes-csi/csi-lib-utils/connection"
	"github.com/kubernetes-csi/csi-test/driver"
	"google.golang.org/grpc"
)

const (
	driverName = "foo/bar"
)

func createMockServer(t *testing.T) (
	*gomock.Controller,
	*driver.MockCSIDriver,
	*driver.MockIdentityServer,
	*driver.MockControllerServer,
	*driver.MockNodeServer,
	*grpc.ClientConn,
	error) {
	// Start the mock server
	mockController := gomock.NewController(t)
	identityServer := driver.NewMockIdentityServer(mockController)
	controllerServer := driver.NewMockControllerServer(mockController)
	nodeServer := driver.NewMockNodeServer(mockController)
	drv := driver.NewMockCSIDriver(&driver.MockCSIDriverServers{
		Identity:   identityServer,
		Controller: controllerServer,
		Node:       nodeServer,
	})
	drv.Start()

	// Create a client connection to it
	addr := drv.Address()
	csiConn, err := connlib.Connect(addr)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return mockController, drv, identityServer, controllerServer, nodeServer, csiConn, nil
}

func TestProbe(t *testing.T) {
	mockController, driver, idServer, _, _, csiConn, err := createMockServer(t)
	if err != nil {
		t.Fatal(err)
	}
	defer mockController.Finish()
	defer driver.Stop()
	defer csiConn.Close()

	var injectedErr error

	inProbe := &csi.ProbeRequest{}
	outProbe := &csi.ProbeResponse{}
	idServer.EXPECT().Probe(gomock.Any(), inProbe).Return(outProbe, injectedErr).Times(1)

	hp := &healthProbe{
		conn:       csiConn,
		driverName: driverName,
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() == "/healthz" {
			hp.checkProbe(rw, req)
		}
	}))
	defer server.Close()

	httpreq, err := http.NewRequest("GET", server.URL+"/healthz", nil)
	if err != nil {
		t.Fatalf("failed to build test request for health check: %v", err)
	}
	_, err = http.DefaultClient.Do(httpreq)
	if err != nil {
		t.Errorf("failed to check probe: %v", err)
	}
}
