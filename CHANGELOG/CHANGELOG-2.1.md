# Release notes for v2.1.0

[Documentation](https://kubernetes-csi.github.io/docs/)
# Changelog since v2.0.0

## Changes by Kind

### Bug or Regression

- Reports an unhealthy status on `/healthz` endpoint, if the CSI driver socket file does not exist anymore (accidentally removed on the host, for example). ([#69](https://github.com/kubernetes-csi/livenessprobe/pull/69), [@nettoclaudio](https://github.com/nettoclaudio))

### Uncategorized

- Publishing of images on k8s.gcr.io ([#71](https://github.com/kubernetes-csi/livenessprobe/pull/71), [@pohly](https://github.com/pohly))

## Dependencies

### Added
_Nothing has changed._

### Changed
_Nothing has changed._

### Removed
_Nothing has changed._
