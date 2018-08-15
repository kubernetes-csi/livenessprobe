#!/bin/bash
set -x

## This file is for livenessprove which runs in a pair with csi
## hostpath

## Must be run from the root of the repo
UDS="/tmp/e2e-csi-sanity.sock"
CSI_ENDPOINT="unix://${UDS}"
CSI_MOUNTPOINT="/mnt"
APP=hostpathplugin

SKIP="WithCapacity"
if [ x${TRAVIS} = x"true" ] ; then
	SKIP="WithCapacity|NodeUnpublishVolume|NodePublishVolume"
fi

git clone https://github.com/kubernetes-csi/drivers $GOPATH/src/github.com/kubernetes-csi/drivers
pushd $GOPATH/src/github.com/kubernetes-csi/drivers
# Build
make hostpath; ret=$?
if [ $ret -ne 0 ]; then
 echo "Failed to build hostpath plugin, file a bug against drivers repo"
 exit 1
fi
popd

sudo rm -rf "$UDS" || true

# Start hostpathplugin in the background
sudo $GOPATH/src/github.com/kubernetes-csi/drivers/_output/$APP --endpoint=$CSI_ENDPOINT --nodeid=1 --v=5 &

# Start liveness probe in the background
sudo ./bin/livenessprobe --csi-address=$CSI_ENDPOINT &

# Give time to CSI hostpathplugin and livenessprobe to initialize
sleep 3

# Requesting health
health=$(curl -I http://localhost:9808/healthz | grep HTTP | awk '{print $2}')
if [[ "x$health" != "x200" ]]; then
  echo "Health check failed, but it was not supposed to, exiting..."
  exit 1
fi

# Killing hostpathplugin
sudo kill -9 $(pidof hostpathplugin)
sleep 3

# Requesting health, should fail since hostpathplugin is gone
health=$(curl -I http://localhost:9808/healthz| grep HTTP | awk '{print $2}')
if [[ "x$health" != "x500" ]]; then
  echo "Health check did not detect driver failure, returned code: $health, exiting..."
  exit 1
fi

sudo rm -f $UDS
exit 0
