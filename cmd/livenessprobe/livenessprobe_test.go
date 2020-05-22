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
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-csi/csi-test/driver"
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
	func()) {
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

	tmpDir, err := ioutil.TempDir("", "livenessprobe_test.*")
	if err != nil {
		t.Errorf("failed to create a temporary socket file name: %v", err)
	}

	csiEndpoint := fmt.Sprintf("%s/csi.sock", tmpDir)
	err = drv.StartOnAddress("unix", csiEndpoint)
	if err != nil {
		t.Errorf("failed to start the csi driver at %s: %v", csiEndpoint, err)
	}

	return mockController, drv, identityServer, controllerServer, nodeServer, func() {
		mockController.Finish()
		drv.Stop()
		os.RemoveAll(csiEndpoint)
	}
}

func TestProbe(t *testing.T) {
	_, driver, idServer, _, _, cleanUpFunc := createMockServer(t)
	defer cleanUpFunc()

	flag.Set("csi-address", driver.Address())
	flag.Parse()

	var injectedErr error

	inProbe := &csi.ProbeRequest{}
	outProbe := &csi.ProbeResponse{}
	idServer.EXPECT().Probe(gomock.Any(), inProbe).Return(outProbe, injectedErr).Times(1)

	hp := &healthProbe{driverName: driverName}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() == "/healthz" {
			hp.checkProbe(rw, req)
		}
	}))
	defer server.Close()

	httpreq, err := http.NewRequest("GET", fmt.Sprintf("%s/healthz", server.URL), nil)
	if err != nil {
		t.Fatalf("failed to build test request for health check: %v", err)
	}

	httpresp, err := http.DefaultClient.Do(httpreq)
	if err != nil {
		t.Errorf("failed to check probe: %v", err)
	}

	expectedStatusCode := http.StatusOK
	if httpresp.StatusCode != expectedStatusCode {
		t.Errorf("expected status code %d but got %d", expectedStatusCode, httpresp.StatusCode)
	}
}

func TestProbe_issue68(t *testing.T) {
	_, driver, idServer, _, _, cleanUpFunc := createMockServer(t)
	defer cleanUpFunc()

	flag.Set("csi-address", driver.Address())
	flag.Parse()

	var injectedErr error

	inProbe := &csi.ProbeRequest{}
	outProbe := &csi.ProbeResponse{}
	idServer.EXPECT().Probe(gomock.Any(), inProbe).Return(outProbe, injectedErr).Times(1)

	hp := &healthProbe{driverName: driverName}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() == "/healthz" {
			hp.checkProbe(rw, req)
		}
	}))
	defer server.Close()

	httpreq, err := http.NewRequest("GET", fmt.Sprintf("%s/healthz", server.URL), nil)
	if err != nil {
		t.Fatalf("failed to build test request for health check: %v", err)
	}

	httpresp, err := http.DefaultClient.Do(httpreq)
	if err != nil {
		t.Errorf("failed to check probe: %v", err)
	}

	expectedStatusCode := http.StatusOK
	if httpresp.StatusCode != expectedStatusCode {
		t.Errorf("expected status code %d but got %d", expectedStatusCode, httpresp.StatusCode)
	}

	err = os.Remove(driver.Address())
	if err != nil {
		t.Errorf("failed to remove the csi driver socket file: %v", err)
	}

	httpreq, err = http.NewRequest("GET", fmt.Sprintf("%s/healthz", server.URL), nil)
	if err != nil {
		t.Fatalf("failed to build test request for health check: %v", err)
	}

	httpresp, err = http.DefaultClient.Do(httpreq)
	if err != nil {
		t.Errorf("failed to check probe: %v", err)
	}

	expectedStatusCode = http.StatusInternalServerError
	if httpresp.StatusCode != expectedStatusCode {
		t.Errorf("expected status code %d but got %d", expectedStatusCode, httpresp.StatusCode)
	}
}
