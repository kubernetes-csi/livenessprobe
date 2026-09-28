# Release notes for v2.8.0

[Documentation](https://kubernetes-csi.github.io/docs/)

# Changelog since v2.7.0

## Changes by Kind

### Other (Cleanup or Flake)

- Fix: CVE-2021-38561 in image build ([#148](https://github.com/kubernetes-csi/livenessprobe/pull/148), [@andyzhangx](https://github.com/andyzhangx))
- Bump kubernetes to v1.25, CSI Spec to v1.6 ([#154](https://github.com/kubernetes-csi/livenessprobe/pull/154), [@humblec](https://github.com/humblec))

## Dependencies

### Added
- github.com/armon/go-socks5: [e753329](https://github.com/armon/go-socks5/tree/e753329)
- github.com/blang/semver/v4: [v4.0.0](https://github.com/blang/semver/v4/tree/v4.0.0)
- github.com/emicklei/go-restful/v3: [v3.8.0](https://github.com/emicklei/go-restful/v3/tree/v3.8.0)
- github.com/go-logr/zapr: [v1.2.3](https://github.com/go-logr/zapr/tree/v1.2.3)
- github.com/go-task/slim-sprig: [348f09d](https://github.com/go-task/slim-sprig/tree/348f09d)
- github.com/google/gnostic: [v0.5.7-v3refs](https://github.com/google/gnostic/tree/v0.5.7-v3refs)
- github.com/google/martian/v3: [v3.0.0](https://github.com/google/martian/v3/tree/v3.0.0)
- github.com/josharian/intern: [v1.0.0](https://github.com/josharian/intern/tree/v1.0.0)
- github.com/kubernetes-csi/csi-test/v5: [v5.0.0](https://github.com/kubernetes-csi/csi-test/v5/tree/v5.0.0)
- github.com/onsi/ginkgo/v2: [v2.1.6](https://github.com/onsi/ginkgo/v2/tree/v2.1.6)
- sigs.k8s.io/json: f223a00

### Changed
- cloud.google.com/go/bigquery: v1.4.0 → v1.8.0
- cloud.google.com/go/pubsub: v1.2.0 → v1.3.1
- cloud.google.com/go/storage: v1.6.0 → v1.10.0
- cloud.google.com/go: v0.54.0 → v0.65.0
- github.com/cespare/xxhash/v2: [v2.1.1 → v2.1.2](https://github.com/cespare/xxhash/v2/compare/v2.1.1...v2.1.2)
- github.com/cncf/udpa/go: [5459f2c → 04548b0](https://github.com/cncf/udpa/go/compare/5459f2c...04548b0)
- github.com/cncf/xds/go: [fbca930 → cb28da3](https://github.com/cncf/xds/go/compare/fbca930...cb28da3)
- github.com/container-storage-interface/spec: [v1.5.0 → v1.6.0](https://github.com/container-storage-interface/spec/compare/v1.5.0...v1.6.0)
- github.com/envoyproxy/go-control-plane: [63b5d3c → 49ff273](https://github.com/envoyproxy/go-control-plane/compare/63b5d3c...49ff273)
- github.com/evanphx/json-patch: [v4.11.0+incompatible → v4.12.0+incompatible](https://github.com/evanphx/json-patch/compare/v4.11.0...v4.12.0)
- github.com/go-kit/log: [v0.1.0 → v0.2.0](https://github.com/go-kit/log/compare/v0.1.0...v0.2.0)
- github.com/go-logfmt/logfmt: [v0.5.0 → v0.5.1](https://github.com/go-logfmt/logfmt/compare/v0.5.0...v0.5.1)
- github.com/go-logr/logr: [v1.2.0 → v1.2.3](https://github.com/go-logr/logr/compare/v1.2.0...v1.2.3)
- github.com/go-openapi/jsonpointer: [v0.19.3 → v0.19.5](https://github.com/go-openapi/jsonpointer/compare/v0.19.3...v0.19.5)
- github.com/go-openapi/jsonreference: [v0.19.3 → v0.19.5](https://github.com/go-openapi/jsonreference/compare/v0.19.3...v0.19.5)
- github.com/go-openapi/swag: [v0.19.5 → v0.19.14](https://github.com/go-openapi/swag/compare/v0.19.5...v0.19.14)
- github.com/golang/mock: [v1.5.0 → v1.6.0](https://github.com/golang/mock/compare/v1.5.0...v1.6.0)
- github.com/google/go-cmp: [v0.5.5 → v0.5.8](https://github.com/google/go-cmp/compare/v0.5.5...v0.5.8)
- github.com/google/pprof: [1ebb73c → 94a9f03](https://github.com/google/pprof/compare/1ebb73c...94a9f03)
- github.com/google/uuid: [v1.1.2 → v1.3.0](https://github.com/google/uuid/compare/v1.1.2...v1.3.0)
- github.com/ianlancetaylor/demangle: [5e5cf60 → 28f6c0f](https://github.com/ianlancetaylor/demangle/compare/5e5cf60...28f6c0f)
- github.com/json-iterator/go: [v1.1.11 → v1.1.12](https://github.com/json-iterator/go/compare/v1.1.11...v1.1.12)
- github.com/mailru/easyjson: [b2ccc51 → v0.7.6](https://github.com/mailru/easyjson/compare/b2ccc51...v0.7.6)
- github.com/moby/term: [9d4ed18 → 3f7ff69](https://github.com/moby/term/compare/9d4ed18...3f7ff69)
- github.com/modern-go/reflect2: [v1.0.1 → v1.0.2](https://github.com/modern-go/reflect2/compare/v1.0.1...v1.0.2)
- github.com/munnerz/goautoneg: [a547fc6 → a7dc8b6](https://github.com/munnerz/goautoneg/compare/a547fc6...a7dc8b6)
- github.com/nxadm/tail: [v1.4.4 → v1.4.8](https://github.com/nxadm/tail/compare/v1.4.4...v1.4.8)
- github.com/onsi/ginkgo: [v1.14.0 → v1.16.4](https://github.com/onsi/ginkgo/compare/v1.14.0...v1.16.4)
- github.com/onsi/gomega: [v1.10.1 → v1.20.1](https://github.com/onsi/gomega/compare/v1.10.1...v1.20.1)
- github.com/prometheus/client_golang: [v1.11.1 → v1.12.2](https://github.com/prometheus/client_golang/compare/v1.11.1...v1.12.2)
- github.com/prometheus/common: [v0.26.0 → v0.37.0](https://github.com/prometheus/common/compare/v0.26.0...v0.37.0)
- github.com/prometheus/procfs: [v0.6.0 → v0.8.0](https://github.com/prometheus/procfs/compare/v0.6.0...v0.8.0)
- github.com/spf13/cobra: [v1.1.3 → v1.4.0](https://github.com/spf13/cobra/compare/v1.1.3...v1.4.0)
- github.com/yuin/goldmark: [v1.3.5 → v1.4.1](https://github.com/yuin/goldmark/compare/v1.3.5...v1.4.1)
- go.opencensus.io: v0.22.3 → v0.22.4
- go.uber.org/zap: v1.17.0 → v1.19.0
- golang.org/x/crypto: 5ea612d → 089bfa5
- golang.org/x/mod: v0.4.2 → 86c51ed
- golang.org/x/net: 37e1c6a → d300de1
- golang.org/x/oauth2: bf48bf1 → ee48083
- golang.org/x/sync: 036812b → 0de741c
- golang.org/x/sys: 59db8d7 → a90be44
- golang.org/x/term: 6a3ed07 → 03fcf44
- golang.org/x/text: v0.3.6 → v0.3.8
- golang.org/x/time: 1f47c86 → 90d013b
- golang.org/x/tools: v0.1.2 → v0.1.12
- google.golang.org/api: v0.20.0 → v0.30.0
- google.golang.org/appengine: v1.6.5 → v1.6.7
- google.golang.org/genproto: f16073e → 4e6b2df
- google.golang.org/grpc: v1.40.0 → v1.48.0
- google.golang.org/protobuf: v1.26.0 → v1.28.1
- gopkg.in/yaml.v3: 496545a → v3.0.1
- honnef.co/go/tools: v0.0.1-2020.1.3 → v0.0.1-2020.1.4
- k8s.io/api: v0.22.0 → v0.25.2
- k8s.io/apimachinery: v0.22.0 → v0.25.2
- k8s.io/client-go: v0.22.0 → v0.25.2
- k8s.io/component-base: v0.22.0 → v0.25.2
- k8s.io/klog/v2: v2.60.0 → v2.70.1
- k8s.io/kube-openapi: 9528897 → 67bda5d
- k8s.io/utils: 3a6ce19 → ee6ede2
- sigs.k8s.io/structured-merge-diff/v4: v4.1.2 → v4.2.3

### Removed
- github.com/kubernetes-csi/csi-test/v4: [555d70a](https://github.com/kubernetes-csi/csi-test/v4/tree/555d70a)
- github.com/robertkrimen/otto: [c382bd3](https://github.com/robertkrimen/otto/tree/c382bd3)
- gopkg.in/sourcemap.v1: v1.0.5
