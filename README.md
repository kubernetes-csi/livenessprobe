# Liveness Probe

The liveness probe is a sidecar container that exposes an HTTP `/healthz`
endpoint, which serves as kubelet's [livenessProbe hook](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes)
to monitor health of a CSI driver.

The liveness probe uses `Probe()` call to check the CSI driver is healthy.
See CSI spec for more information about Probe API call.
[Container Storage Interface (CSI)](https://github.com/container-storage-interface/spec/blob/master/spec.md#probe)

## Compatibility
This information reflects the head of this branch.

| Compatible with CSI Version                                                                | Container Image                      | [Min K8s Version](https://kubernetes-csi.github.io/docs/kubernetes-compatibility.html#minimum-version) |
| ------------------------------------------------------------------------------------------ | -------------------------------------| --------------- |
| [CSI Spec v1.0.0](https://github.com/container-storage-interface/spec/releases/tag/v1.0.0) | k8s.gcr.io/sig-storage/livenessprobe | 1.13            |


## Usage

See [hostpath-with-livenessprobe.yaml](https://github.com/kubernetes-csi/livenessprobe/blob/master/deployment/kubernetes/livenessprobe-sidecar.yaml)
for example how to use the liveness probe with a CSI driver. Notice that actual
`livenessProbe` is set on the container with the CSI driver. This way, Kubernetes
restarts the CSI driver container when the probe fails. The liveness probe
sidecar container only provides the HTTP endpoint for the probe and does not
contain `livenessProbe` section by itself.

```yaml
kind: Pod
spec:
  containers:
  # Container with CSI driver
  - name: hostpath-driver
    image: quay.io/k8scsi/hostpathplugin:v0.2.0
    # Defining port which will be used to GET plugin health status
    # 9808 is default, but can be changed.
    ports:
    - containerPort: 9808
      name: healthz
      protocol: TCP
    # The probe
    livenessProbe:
      failureThreshold: 5
      httpGet:
        path: /healthz
        port: healthz
      initialDelaySeconds: 10
      timeoutSeconds: 3
      periodSeconds: 2
    volumeMounts:
    - mountPath: /csi
      name: socket-dir
    # ...
 # The liveness probe sidecar container
 - name: liveness-probe
    imagePullPolicy: Always
    image: quay.io/k8scsi/livenessprobe:v1.1.0
    args:
    - --csi-address=/csi/csi.sock
    volumeMounts:
    - mountPath: /csi
      name: socket-dir
    # ...
```

### Command line options

#### Recommended optional arguments

* `--csi-address <path to CSI socket>`: This is the path to the CSI driver socket inside the pod that the external-provisioner container will use to issue CSI operations (`/run/csi/socket` is used by default).

#### Other recognized arguments

* `--health-port <number>`: TCP ports for listening for HTTP requests (default "9808")

* `--probe-timeout <duration>`: Maximum duration of single `Probe()` call (default "1s").

* `--metrics-address <port>`: The TCP network address where the prometheus metrics endpoint will listen (example: `:8080`). The default is empty string, which means metrics endpoint is disabled.

* `--metrics-path <path>`: The HTTP path where prometheus metrics will be exposed. Default is `/metrics`."

* All glog / klog arguments are supported, such as `-v <log level>` or `-alsologtostderr`.

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

* [Slack channel](https://kubernetes.slack.com/messages/sig-storage)
* [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-storage)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).
