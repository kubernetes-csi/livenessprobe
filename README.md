# Liveness Probe

## Overview

The purpose of liveness probe is to be able to detect liveness of CSI driver
and in case of as a driver's failure report it by returning non zero return code.
Livenessprobe leverages CSI Probe API call, which must be answered by CSI
compatible driver.
Liveness probe is meant to be used in pair with kubelet's LinessProbe hook, which
executes it periodically and check the rerutn code. Non zero return code indicates
to kubelet that pod is not healthy and kubelet schedules pod restart to recover it.

See CSI spec for more information about Probe API call.
[Container Storage Interface (CSI)](https://github.com/container-storage-interface/spec/blob/master/spec.md#probe)

## Livenessprobe
### Configuration Requirements

* --v defines level of logging
* --csi-address path to the location of csi socket opened by the ddriver
* --connection-timeout  time to wait for the driver to return reply on Probe request

### Compiling
Livenessprobe can be compiled in a form of a binary file or in a form of a container. When compiled
as a binary file, it gets stored in bin folder with the name livenessprobe. When compiled as a container,
the resulting image is stored in a local docker's image store and tagged as
quay.io/k8scsi/livenessprobe:canary

To compile just a binary file:
```
$ make livenessprobe
```

To build a container:
```
$ make livenessprobe-container
```
By running:
```
$ docker images | grep livenessprobe
```
You should see the following line in the output:
```
quay.io/k8scsi/livenessprobe                    canary     8f65dd5f789a        16 hours ago        16MB
```

### Using livenessprobe

Below is an example of sidecar container which needs to be added to the CSI driver yaml.

```yaml
  - name: liveness-probe
    image: quay.io/k8scsi/livenessprobe:v0.2.0
    imagePullPolicy: Always
    command: ["/bin/sh"]
    args: ["-c", "while true; do sleep 10;done"]    
    livenessProbe:
      exec:
        command:
        - ./livenessprobe
        - --v=6
        - --csi-address=/csi/csi.sock
        - --connection-timeout=3s
      initialDelaySeconds: 10
      timeoutSeconds: 3
      periodSeconds: 2
      failureThreshold: 1 
    volumeMounts:
    - mountPath: /csi
      name: socket-dir

```

Please submit an issue at:[Issues](https://github.com/kubernetes-csi/livenessprobe/issues)

