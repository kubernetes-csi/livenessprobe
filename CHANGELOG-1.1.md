# Changelog since v1.0.2

## Deprecations
* Command line flag `--connection-timeout` is deprecated and has no effect.

## Notable Features
* livenessprobe tries to connect to the CSI driver indefinitely only at startup ([##37](https://github.com/kubernetes-csi/livenessprobe/pull/37))

## Other Notable Changes
* Use distroless as base image ([##40](https://github.com/kubernetes-csi/livenessprobe/pull/40))
* Migrate to k8s.io/klog from glog ([##36](https://github.com/kubernetes-csi/livenessprobe/pull/36))
* add prune to gopkg.toml ([##30](https://github.com/kubernetes-csi/livenessprobe/pull/30))
