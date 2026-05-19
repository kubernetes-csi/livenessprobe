# Release notes for v2.19.0

[Documentation](https://kubernetes-csi.github.io/docs/)

# Changelog since v2.18.0

## Changes by Kind

### Bug or Regression

- Fix: CVE-2026-33186 ([#407](https://github.com/kubernetes-csi/livenessprobe/pull/407), [@andyzhangx](https://github.com/andyzhangx))

### Other (Cleanup or Flake)

- Bump k8s dependencies to v1.36.1 ([#414](https://github.com/kubernetes-csi/livenessprobe/pull/414), [@dfajmon](https://github.com/dfajmon))

### Uncategorized

- Fix - CVE-2026-39882, CVE-2026-24051, CVE-2026-39883 ([#413](https://github.com/kubernetes-csi/livenessprobe/pull/413), [@sammedsingalkar09](https://github.com/sammedsingalkar09))

## Dependencies

### Added
- github.com/cenkalti/backoff/v5: [v5.0.3](https://github.com/cenkalti/backoff/tree/v5.0.3)
- k8s.io/streaming: v0.36.1

### Changed
- cel.dev/expr: v0.24.0 → v0.25.1
- github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp: [v1.30.0 → v1.31.0](https://github.com/GoogleCloudPlatform/opentelemetry-operations-go/compare/detectors/gcp/v1.30.0...detectors/gcp/v1.31.0)
- github.com/cncf/xds/go: [0feb691 → ee656c7](https://github.com/cncf/xds/compare/0feb691...ee656c7)
- github.com/emicklei/go-restful/v3: [v3.12.2 → v3.13.0](https://github.com/emicklei/go-restful/compare/v3.12.2...v3.13.0)
- github.com/envoyproxy/go-control-plane/envoy: [v1.35.0 → v1.36.0](https://github.com/envoyproxy/go-control-plane/compare/envoy/v1.35.0...envoy/v1.36.0)
- github.com/envoyproxy/go-control-plane: [75eaa19 → v0.14.0](https://github.com/envoyproxy/go-control-plane/compare/75eaa19...v0.14.0)
- github.com/envoyproxy/protoc-gen-validate: [v1.2.1 → v1.3.0](https://github.com/envoyproxy/protoc-gen-validate/compare/v1.2.1...v1.3.0)
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.26.3 → v2.27.7](https://github.com/grpc-ecosystem/grpc-gateway/compare/v2.26.3...v2.27.7)
- github.com/kubernetes-csi/csi-lib-utils: [v0.23.1 → v0.24.0](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.23.1...v0.24.0)
- github.com/moby/spdystream: [v0.5.0 → v0.5.1](https://github.com/moby/spdystream/compare/v0.5.0...v0.5.1)
- github.com/onsi/ginkgo/v2: [v2.27.2 → v2.22.0](https://github.com/onsi/ginkgo/compare/v2.27.2...v2.22.0)
- github.com/onsi/gomega: [v1.38.2 → v1.36.1](https://github.com/onsi/gomega/compare/v1.38.2...v1.36.1)
- go.opentelemetry.io/contrib/detectors/gcp: v1.38.0 → v1.39.0
- go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc: v0.64.0 → v0.68.0
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.61.0 → v0.65.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.34.0 → v1.40.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.34.0 → v1.40.0
- go.opentelemetry.io/otel/metric: v1.39.0 → v1.43.0
- go.opentelemetry.io/otel/sdk/metric: v1.39.0 → v1.43.0
- go.opentelemetry.io/otel/sdk: v1.39.0 → v1.43.0
- go.opentelemetry.io/otel/trace: v1.39.0 → v1.43.0
- go.opentelemetry.io/otel: v1.39.0 → v1.43.0
- go.opentelemetry.io/proto/otlp: v1.5.0 → v1.9.0
- golang.org/x/crypto: v0.47.0 → v0.49.0
- golang.org/x/mod: v0.31.0 → v0.33.0
- golang.org/x/net: v0.49.0 → v0.52.0
- golang.org/x/sync: v0.19.0 → v0.20.0
- golang.org/x/sys: v0.40.0 → v0.42.0
- golang.org/x/term: v0.39.0 → v0.41.0
- golang.org/x/text: v0.33.0 → v0.35.0
- golang.org/x/time: v0.9.0 → v0.14.0
- golang.org/x/tools: v0.40.0 → v0.42.0
- gonum.org/v1/gonum: v0.16.0 → v0.17.0
- google.golang.org/genproto/googleapis/api: ab9386a → 8636f87
- google.golang.org/genproto/googleapis/rpc: 99fd39f → 6f92a3b
- google.golang.org/grpc: v1.78.0 → v1.80.0
- google.golang.org/protobuf: v1.36.11 → f2248ac
- k8s.io/api: v0.35.0 → v0.36.1
- k8s.io/apimachinery: v0.35.0 → v0.36.1
- k8s.io/client-go: v0.35.0 → v0.36.1
- k8s.io/component-base: v0.35.0 → v0.36.1
- k8s.io/klog/v2: v2.130.1 → v2.140.0
- k8s.io/kube-openapi: 4e65d59 → 43fb72c
- k8s.io/utils: 914a6e7 → b8788ab
- sigs.k8s.io/structured-merge-diff/v6: v6.3.1 → v6.3.2

### Removed
- github.com/armon/go-socks5: [e753329](https://github.com/armon/go-socks5/tree/e753329)
- github.com/cenkalti/backoff/v4: [v4.3.0](https://github.com/cenkalti/backoff/tree/v4.3.0)
- github.com/gogo/protobuf: [v1.3.2](https://github.com/gogo/protobuf/tree/v1.3.2)
- github.com/pkg/errors: [v0.9.1](https://github.com/pkg/errors/tree/v0.9.1)
