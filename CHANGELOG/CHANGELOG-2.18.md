# Release notes for v2.18.0

[Documentation](https://kubernetes-csi.github.io/docs/)

# Changelog since v2.17.0

## Changes by Kind

### Bug or Regression

- Updated go version to fix CVE-2025-68121. ([#399](https://github.com/kubernetes-csi/livenessprobe/pull/399), [@jsafrane](https://github.com/jsafrane))

### Uncategorized

- Bump dependencies to kubernetes v1.35 ([#397](https://github.com/kubernetes-csi/livenessprobe/pull/397), [@rhrmo](https://github.com/rhrmo))

## Dependencies

### Added
- github.com/NYTimes/gziphandler: [v1.1.1](https://github.com/NYTimes/gziphandler/tree/v1.1.1)
- github.com/go-task/slim-sprig/v3: [v3.0.0](https://github.com/go-task/slim-sprig/tree/v3.0.0)
- github.com/golang-jwt/jwt/v5: [v5.3.0](https://github.com/golang-jwt/jwt/tree/v5.3.0)
- golang.org/x/tools/go/expect: v0.1.0-deprecated
- golang.org/x/tools/go/packages/packagestest: v0.1.1-deprecated
- gonum.org/v1/gonum: v0.16.0
- k8s.io/gengo/v2: 85fd79d

### Changed
- cel.dev/expr: v0.20.0 → v0.24.0
- cloud.google.com/go/compute/metadata: v0.6.0 → v0.9.0
- github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp: [v1.26.0 → v1.30.0](https://github.com/GoogleCloudPlatform/opentelemetry-operations-go/compare/detectors/gcp/v1.26.0...detectors/gcp/v1.30.0)
- github.com/alecthomas/units: [b94a6e3 → 0f3dac3](https://github.com/alecthomas/units/compare/b94a6e3...0f3dac3)
- github.com/cncf/xds/go: [2f00578 → 0feb691](https://github.com/cncf/xds/compare/2f00578...0feb691)
- github.com/container-storage-interface/spec: [v1.11.0 → v1.12.0](https://github.com/container-storage-interface/spec/compare/v1.11.0...v1.12.0)
- github.com/envoyproxy/go-control-plane/envoy: [v1.32.4 → v1.35.0](https://github.com/envoyproxy/go-control-plane/compare/envoy/v1.32.4...envoy/v1.35.0)
- github.com/envoyproxy/go-control-plane: [v0.13.4 → 75eaa19](https://github.com/envoyproxy/go-control-plane/compare/v0.13.4...75eaa19)
- github.com/go-jose/go-jose/v4: [v4.0.4 → v4.1.3](https://github.com/go-jose/go-jose/compare/v4.0.4...v4.1.3)
- github.com/go-logr/logr: [v1.4.2 → v1.4.3](https://github.com/go-logr/logr/compare/v1.4.2...v1.4.3)
- github.com/golang/glog: [v1.2.4 → v1.2.5](https://github.com/golang/glog/compare/v1.2.4...v1.2.5)
- github.com/google/pprof: [94a9f03 → 40e02aa](https://github.com/google/pprof/compare/94a9f03...40e02aa)
- github.com/kubernetes-csi/csi-lib-utils: [v0.22.0 → v0.23.1](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.22.0...v0.23.1)
- github.com/kubernetes-csi/csi-test/v5: [v5.3.1 → v5.4.0](https://github.com/kubernetes-csi/csi-test/compare/v5.3.1...v5.4.0)
- github.com/onsi/ginkgo/v2: [v2.21.0 → v2.27.2](https://github.com/onsi/ginkgo/compare/v2.21.0...v2.27.2)
- github.com/onsi/gomega: [v1.35.1 → v1.38.2](https://github.com/onsi/gomega/compare/v1.35.1...v1.38.2)
- github.com/prometheus/client_golang: [v1.22.0 → v1.23.2](https://github.com/prometheus/client_golang/compare/v1.22.0...v1.23.2)
- github.com/prometheus/common: [v0.63.0 → v0.67.5](https://github.com/prometheus/common/compare/v0.63.0...v0.67.5)
- github.com/prometheus/procfs: [v0.16.1 → v0.19.2](https://github.com/prometheus/procfs/compare/v0.16.1...v0.19.2)
- github.com/rogpeppe/go-internal: [v1.13.1 → v1.14.1](https://github.com/rogpeppe/go-internal/compare/v1.13.1...v1.14.1)
- github.com/spf13/cobra: [v1.9.1 → v1.10.2](https://github.com/spf13/cobra/compare/v1.9.1...v1.10.2)
- github.com/spf13/pflag: [v1.0.6 → v1.0.10](https://github.com/spf13/pflag/compare/v1.0.6...v1.0.10)
- github.com/spiffe/go-spiffe/v2: [v2.5.0 → v2.6.0](https://github.com/spiffe/go-spiffe/compare/v2.5.0...v2.6.0)
- github.com/stretchr/testify: [v1.10.0 → v1.11.1](https://github.com/stretchr/testify/compare/v1.10.0...v1.11.1)
- go.opentelemetry.io/auto/sdk: v1.1.0 → v1.2.1
- go.opentelemetry.io/contrib/detectors/gcp: v1.34.0 → v1.38.0
- go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc: v0.60.0 → v0.64.0
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.58.0 → v0.61.0
- go.opentelemetry.io/otel/metric: v1.35.0 → v1.39.0
- go.opentelemetry.io/otel/sdk/metric: v1.34.0 → v1.39.0
- go.opentelemetry.io/otel/sdk: v1.34.0 → v1.39.0
- go.opentelemetry.io/otel/trace: v1.35.0 → v1.39.0
- go.opentelemetry.io/otel: v1.35.0 → v1.39.0
- go.uber.org/zap: v1.27.0 → v1.27.1
- go.yaml.in/yaml/v2: v2.4.2 → v2.4.3
- golang.org/x/crypto: v0.38.0 → v0.47.0
- golang.org/x/mod: v0.17.0 → v0.31.0
- golang.org/x/net: v0.40.0 → v0.49.0
- golang.org/x/oauth2: v0.27.0 → v0.34.0
- golang.org/x/sync: v0.14.0 → v0.19.0
- golang.org/x/sys: v0.33.0 → v0.40.0
- golang.org/x/term: v0.32.0 → v0.39.0
- golang.org/x/text: v0.25.0 → v0.33.0
- golang.org/x/tools: e35e4cc → v0.40.0
- google.golang.org/genproto/googleapis/api: a0af3ef → ab9386a
- google.golang.org/genproto/googleapis/rpc: a0af3ef → 99fd39f
- google.golang.org/grpc: v1.72.1 → v1.78.0
- google.golang.org/protobuf: v1.36.6 → v1.36.11
- gopkg.in/evanphx/json-patch.v4: v4.12.0 → v4.13.0
- k8s.io/api: v0.34.0 → v0.35.0
- k8s.io/apimachinery: v0.34.0 → v0.35.0
- k8s.io/client-go: v0.34.0 → v0.35.0
- k8s.io/component-base: v0.34.0 → v0.35.0
- k8s.io/kube-openapi: f3f2b99 → 4e65d59
- k8s.io/utils: 4c0f3b2 → 914a6e7
- sigs.k8s.io/json: cfa47c3 → 2d32026
- sigs.k8s.io/structured-merge-diff/v6: v6.3.0 → v6.3.1

### Removed
- github.com/go-task/slim-sprig: [52ccab3](https://github.com/go-task/slim-sprig/tree/52ccab3)
- github.com/kisielk/errcheck: [v1.5.0](https://github.com/kisielk/errcheck/tree/v1.5.0)
- github.com/kisielk/gotool: [v1.0.0](https://github.com/kisielk/gotool/tree/v1.0.0)
- github.com/zeebo/errs: [v1.4.0](https://github.com/zeebo/errs/tree/v1.4.0)
- sigs.k8s.io/structured-merge-diff/v4: v4.6.0
