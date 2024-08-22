# Release notes for v2.14.0

[Documentation](https://kubernetes-csi.github.io/docs/)

# Changelog since v2.13.0

## Changes by Kind

### Other (Cleanup or Flake)

- Updates Kubernetes dependencies to 1.31.0-rc.0 ([#284](https://github.com/kubernetes-csi/livenessprobe/pull/284), [@dfajmon](https://github.com/dfajmon))
- Updates Kubernetes dependencies to v1.31.0 ([#288](https://github.com/kubernetes-csi/livenessprobe/pull/288), [@dfajmon](https://github.com/dfajmon))

### Uncategorized

- Updated go to 1.22.5 to resolve CVE-2024-24790 ([#283](https://github.com/kubernetes-csi/livenessprobe/pull/283), [@sdx-jkataja](https://github.com/sdx-jkataja))

## Dependencies

### Added
- cel.dev/expr: v0.15.0
- github.com/go-task/slim-sprig/v3: [v3.0.0](https://github.com/go-task/slim-sprig/tree/v3.0.0)
- github.com/klauspost/compress: [v1.17.9](https://github.com/klauspost/compress/tree/v1.17.9)
- github.com/kylelemons/godebug: [v1.1.0](https://github.com/kylelemons/godebug/tree/v1.1.0)
- gopkg.in/evanphx/json-patch.v4: v4.12.0

### Changed
- cloud.google.com/go/compute/metadata: v0.2.3 → v0.3.0
- github.com/cenkalti/backoff/v4: [v4.2.1 → v4.3.0](https://github.com/cenkalti/backoff/compare/v4.2.1...v4.3.0)
- github.com/cncf/xds/go: [0fa0005 → 555b57e](https://github.com/cncf/xds/compare/0fa0005...555b57e)
- github.com/cpuguy83/go-md2man/v2: [v2.0.3 → v2.0.4](https://github.com/cpuguy83/go-md2man/compare/v2.0.3...v2.0.4)
- github.com/davecgh/go-spew: [v1.1.1 → d8f796a](https://github.com/davecgh/go-spew/compare/v1.1.1...d8f796a)
- github.com/felixge/httpsnoop: [v1.0.3 → v1.0.4](https://github.com/felixge/httpsnoop/compare/v1.0.3...v1.0.4)
- github.com/fxamacker/cbor/v2: [v2.6.0 → v2.7.0](https://github.com/fxamacker/cbor/compare/v2.6.0...v2.7.0)
- github.com/go-logr/logr: [v1.4.1 → v1.4.2](https://github.com/go-logr/logr/compare/v1.4.1...v1.4.2)
- github.com/go-openapi/swag: [v0.22.3 → v0.22.4](https://github.com/go-openapi/swag/compare/v0.22.3...v0.22.4)
- github.com/golang/glog: [v1.2.0 → v1.2.1](https://github.com/golang/glog/compare/v1.2.0...v1.2.1)
- github.com/google/pprof: [4bb14d4 → 4bfdf5a](https://github.com/google/pprof/compare/4bb14d4...4bfdf5a)
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.16.0 → v2.20.0](https://github.com/grpc-ecosystem/grpc-gateway/compare/v2.16.0...v2.20.0)
- github.com/kubernetes-csi/csi-lib-utils: [v0.18.0 → v0.19.0](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.18.0...v0.19.0)
- github.com/moby/spdystream: [v0.2.0 → v0.4.0](https://github.com/moby/spdystream/compare/v0.2.0...v0.4.0)
- github.com/moby/term: [1aeaba8 → v0.5.0](https://github.com/moby/term/compare/1aeaba8...v0.5.0)
- github.com/onsi/ginkgo/v2: [v2.15.0 → v2.19.0](https://github.com/onsi/ginkgo/compare/v2.15.0...v2.19.0)
- github.com/onsi/gomega: [v1.31.0 → v1.33.1](https://github.com/onsi/gomega/compare/v1.31.0...v1.33.1)
- github.com/pmezard/go-difflib: [v1.0.0 → 5d4384e](https://github.com/pmezard/go-difflib/compare/v1.0.0...5d4384e)
- github.com/prometheus/client_golang: [v1.19.1 → v1.20.0](https://github.com/prometheus/client_golang/compare/v1.19.1...v1.20.0)
- github.com/prometheus/common: [v0.53.0 → v0.55.0](https://github.com/prometheus/common/compare/v0.53.0...v0.55.0)
- github.com/prometheus/procfs: [v0.14.0 → v0.15.1](https://github.com/prometheus/procfs/compare/v0.14.0...v0.15.1)
- github.com/rogpeppe/go-internal: [v1.10.0 → v1.12.0](https://github.com/rogpeppe/go-internal/compare/v1.10.0...v1.12.0)
- github.com/spf13/cobra: [v1.8.0 → v1.8.1](https://github.com/spf13/cobra/compare/v1.8.0...v1.8.1)
- go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc: v0.51.0 → v0.53.0
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.44.0 → v0.53.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.19.0 → v1.27.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.19.0 → v1.28.0
- go.opentelemetry.io/otel/metric: v1.26.0 → v1.28.0
- go.opentelemetry.io/otel/sdk: v1.19.0 → v1.28.0
- go.opentelemetry.io/otel/trace: v1.26.0 → v1.28.0
- go.opentelemetry.io/otel: v1.26.0 → v1.28.0
- go.opentelemetry.io/proto/otlp: v1.0.0 → v1.3.1
- golang.org/x/crypto: v0.23.0 → v0.26.0
- golang.org/x/mod: v0.8.0 → v0.17.0
- golang.org/x/net: v0.25.0 → v0.28.0
- golang.org/x/oauth2: v0.18.0 → v0.21.0
- golang.org/x/sync: v0.7.0 → v0.8.0
- golang.org/x/sys: v0.20.0 → v0.23.0
- golang.org/x/term: v0.20.0 → v0.23.0
- golang.org/x/text: v0.15.0 → v0.17.0
- golang.org/x/tools: v0.18.0 → e35e4cc
- google.golang.org/genproto/googleapis/api: 6ceb2ff → 5315273
- google.golang.org/genproto/googleapis/rpc: 6ceb2ff → f6361c8
- google.golang.org/grpc: v1.63.2 → v1.65.0
- google.golang.org/protobuf: v1.34.1 → v1.34.2
- k8s.io/api: v0.30.0 → v0.31.0
- k8s.io/apimachinery: v0.30.0 → v0.31.0
- k8s.io/client-go: v0.30.0 → v0.31.0
- k8s.io/component-base: v0.30.0 → v0.31.0
- k8s.io/klog/v2: v2.120.1 → v2.130.1
- k8s.io/utils: 3b25d92 → 18e509b
- sigs.k8s.io/yaml: v1.3.0 → v1.4.0

### Removed
- cloud.google.com/go/compute: v1.24.0
- github.com/evanphx/json-patch: [v4.12.0+incompatible](https://github.com/evanphx/json-patch/tree/v4.12.0)
- github.com/matttproud/golang_protobuf_extensions: [v1.0.4](https://github.com/matttproud/golang_protobuf_extensions/tree/v1.0.4)
- google.golang.org/appengine: v1.6.8
- google.golang.org/genproto: 6ceb2ff
