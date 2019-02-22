# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

REGISTRY_NAME = quay.io/k8scsi
IMAGE_VERSION = canary
IMAGE_NAME=livenessprobe
IMAGE_TAG=$(REGISTRY_NAME)/$(IMAGE_NAME):$(IMAGE_VERSION)

.PHONY: all liveness clean test

ifdef V
TESTARGS = -v -args -alsologtostderr -v 5
else
TESTARGS =
endif

all: livenessprobe

livenessprobe:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ./bin/livenessprobe ./cmd/livenessprobe

macos-livenessprobe:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin go build -a -ldflags '-extldflags "-static"' -o ./bin/livenessprobe.osx ./cmd/livenessprobe

livenessprobe-container: livenessprobe
	docker build -t $(IMAGE_TAG) -f ./Dockerfile .

push: livenessprobe-container
	docker push $(IMAGE_TAG)

clean:
	rm -rf bin

test:
	go test `go list ./... | grep -v 'vendor'` $(TESTARGS)
	go vet `go list ./... | grep -v vendor`
