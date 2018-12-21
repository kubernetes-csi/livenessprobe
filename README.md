# Liveness Probe

## Overview

The purpose of liveness probe is to be able to detect liveness of CSI driver
and in case of a driver's failure, report it by returning non zero return code.
Livenessprobe leverages CSI Probe API call, which must be answered by CSI
compatible driver.
Liveness probe is meant to be used in pair with kubelet's LivenessProbe hook, which
executes it periodically and check the return code. Non zero return code indicates
to kubelet that pod is not healthy and kubelet schedules pod restart to recover it.

See CSI spec for more information about Probe API call.
[Container Storage Interface (CSI)](https://github.com/container-storage-interface/spec/blob/master/spec.md#probe)

## Livenessprobe
### Configuration Requirements

*  -connection-timeout duration
       Timeout for waiting for CSI driver socket in seconds. (default 30s)
*  -csi-address string
       Address of the CSI driver socket. (default "/run/csi/socket")
*  -health-port string
       TCP ports for listening healthz requests (default "9808")

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
  - name: hostpath-driver
    image: quay.io/k8scsi/hostpathplugin:v0.2.0
    imagePullPolicy: Always
    securityContext:
      privileged: true
#
# Defining port which will be used to GET plugin health status
# 9808 is default, but can be changed.
#
    ports:
    - containerPort: 9808
      name: healthz
      protocol: TCP
    livenessProbe:
      failureThreshold: 5
      httpGet:
        path: /healthz
        port: healthz
      initialDelaySeconds: 10
      timeoutSeconds: 3
      periodSeconds: 2
# 
   volumeMounts:
    - mountPath: /csi
      name: socket-dir
    - mountPath: /var/lib/kubelet/pods
      mountPropagation: Bidirectional
      name: mountpoint-dir
    args:
    - --v=5
    - --endpoint=$(CSI_ENDPOINT)
    - --nodeid=$(KUBE_NODE_NAME)
    env:
    - name: CSI_ENDPOINT
      value: unix:///csi/csi.sock
    - name: KUBE_NODE_NAME
      valueFrom:
        fieldRef:
          apiVersion: v1
          fieldPath: spec.nodeName
#
# Spec for liveness probe sidecar container
# 
 - name: liveness-probe
    imagePullPolicy: Always
    volumeMounts:
    - mountPath: /csi
      name: socket-dir
    image: quay.io/k8scsi/livenessprobe:v0.2.0
    args:
    - --csi-address=/csi/csi.sock
    - --connection-timeout=3s
#

```

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- [Slack channel](https://kubernetes.slack.com/messages/sig-storage)
- [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-storage)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).
