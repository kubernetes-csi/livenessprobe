# Release notes for v2.11.0

[Documentation](https://kubernetes-csi.github.io/docs/)

# Changelog since v2.10.0

## Changes by Kind

### Bug or Regression

- Fixed panic on connection failure. (#192, @jsafrane)

### Uncategorized

- Update kubernetes dependencies to v1.28.0 (#203, @Sneha-at)

## Dependencies

### Added
- cloud.google.com/go/compute/metadata: v0.2.3
- cloud.google.com/go/compute: v1.21.0
- github.com/alecthomas/kingpin/v2: [v2.3.2](https://github.com/alecthomas/kingpin/v2/tree/v2.3.2)
- github.com/google/gnostic-models: [v0.6.8](https://github.com/google/gnostic-models/tree/v0.6.8)
- github.com/xhit/go-str2duration/v2: [v2.1.0](https://github.com/xhit/go-str2duration/v2/tree/v2.1.0)
- google.golang.org/genproto/googleapis/api: 782d3b1
- google.golang.org/genproto/googleapis/rpc: 782d3b1

### Changed
- cloud.google.com/go: v0.65.0 → v0.34.0
- github.com/alecthomas/units: [f65c72e → b94a6e3](https://github.com/alecthomas/units/compare/f65c72e...b94a6e3)
- github.com/cenkalti/backoff/v4: [v4.1.3 → v4.2.1](https://github.com/cenkalti/backoff/v4/compare/v4.1.3...v4.2.1)
- github.com/census-instrumentation/opencensus-proto: [v0.2.1 → v0.4.1](https://github.com/census-instrumentation/opencensus-proto/compare/v0.2.1...v0.4.1)
- github.com/cespare/xxhash/v2: [v2.1.2 → v2.2.0](https://github.com/cespare/xxhash/v2/compare/v2.1.2...v2.2.0)
- github.com/cncf/udpa/go: [04548b0 → c52dc94](https://github.com/cncf/udpa/go/compare/04548b0...c52dc94)
- github.com/cncf/xds/go: [cb28da3 → e9ce688](https://github.com/cncf/xds/go/compare/cb28da3...e9ce688)
- github.com/container-storage-interface/spec: [v1.7.0 → v1.8.0](https://github.com/container-storage-interface/spec/compare/v1.7.0...v1.8.0)
- github.com/envoyproxy/go-control-plane: [49ff273 → v0.11.1](https://github.com/envoyproxy/go-control-plane/compare/49ff273...v0.11.1)
- github.com/envoyproxy/protoc-gen-validate: [v0.1.0 → v1.0.2](https://github.com/envoyproxy/protoc-gen-validate/compare/v0.1.0...v1.0.2)
- github.com/evanphx/json-patch: [v4.12.0+incompatible → v5.6.0+incompatible](https://github.com/evanphx/json-patch/compare/v4.12.0...v5.6.0)
- github.com/go-kit/log: [v0.2.0 → v0.2.1](https://github.com/go-kit/log/compare/v0.2.0...v0.2.1)
- github.com/go-logr/logr: [v1.2.3 → v1.2.4](https://github.com/go-logr/logr/compare/v1.2.3...v1.2.4)
- github.com/go-openapi/jsonreference: [v0.20.1 → v0.20.2](https://github.com/go-openapi/jsonreference/compare/v0.20.1...v0.20.2)
- github.com/go-task/slim-sprig: [348f09d → 52ccab3](https://github.com/go-task/slim-sprig/compare/348f09d...52ccab3)
- github.com/golang/glog: [23def4e → v1.1.0](https://github.com/golang/glog/compare/23def4e...v1.1.0)
- github.com/google/gofuzz: [v1.1.0 → v1.2.0](https://github.com/google/gofuzz/compare/v1.1.0...v1.2.0)
- github.com/inconshreveable/mousetrap: [v1.0.1 → v1.1.0](https://github.com/inconshreveable/mousetrap/compare/v1.0.1...v1.1.0)
- github.com/kr/pretty: [v0.3.0 → v0.3.1](https://github.com/kr/pretty/compare/v0.3.0...v0.3.1)
- github.com/kubernetes-csi/csi-lib-utils: [v0.12.0 → v0.14.0](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.12.0...v0.14.0)
- github.com/matttproud/golang_protobuf_extensions: [v1.0.2 → v1.0.4](https://github.com/matttproud/golang_protobuf_extensions/compare/v1.0.2...v1.0.4)
- github.com/onsi/ginkgo/v2: [v2.9.1 → v2.9.4](https://github.com/onsi/ginkgo/v2/compare/v2.9.1...v2.9.4)
- github.com/onsi/gomega: [v1.27.4 → v1.20.0](https://github.com/onsi/gomega/compare/v1.27.4...v1.20.0)
- github.com/prometheus/client_golang: [v1.14.0 → v1.16.0](https://github.com/prometheus/client_golang/compare/v1.14.0...v1.16.0)
- github.com/prometheus/client_model: [v0.3.0 → v0.4.0](https://github.com/prometheus/client_model/compare/v0.3.0...v0.4.0)
- github.com/prometheus/common: [v0.37.0 → v0.44.0](https://github.com/prometheus/common/compare/v0.37.0...v0.44.0)
- github.com/prometheus/procfs: [v0.8.0 → v0.10.1](https://github.com/prometheus/procfs/compare/v0.8.0...v0.10.1)
- github.com/rogpeppe/go-internal: [v1.3.0 → v1.10.0](https://github.com/rogpeppe/go-internal/compare/v1.3.0...v1.10.0)
- github.com/spf13/cobra: [v1.6.0 → v1.7.0](https://github.com/spf13/cobra/compare/v1.6.0...v1.7.0)
- github.com/stretchr/objx: [v0.1.1 → v0.1.0](https://github.com/stretchr/objx/compare/v0.1.1...v0.1.0)
- github.com/stretchr/testify: [v1.8.1 → v1.8.2](https://github.com/stretchr/testify/compare/v1.8.1...v1.8.2)
- go.uber.org/atomic: v1.7.0 → v1.10.0
- go.uber.org/multierr: v1.6.0 → v1.11.0
- golang.org/x/crypto: 089bfa5 → v0.11.0
- golang.org/x/exp: 6cc2880 → 509febe
- golang.org/x/lint: 738671d → d0100b6
- golang.org/x/net: v0.8.0 → v0.13.0
- golang.org/x/oauth2: ee48083 → v0.10.0
- golang.org/x/sync: 0de741c → v0.3.0
- golang.org/x/sys: v0.6.0 → v0.10.0
- golang.org/x/term: v0.6.0 → v0.10.0
- golang.org/x/text: v0.8.0 → v0.11.0
- golang.org/x/time: 90d013b → v0.3.0
- golang.org/x/tools: v0.7.0 → v0.8.0
- google.golang.org/genproto: 4e6b2df → 782d3b1
- google.golang.org/grpc: v1.51.0 → v1.58.0
- google.golang.org/protobuf: v1.28.1 → v1.31.0
- gopkg.in/check.v1: 41f04d3 → 10cb982
- honnef.co/go/tools: v0.0.1-2020.1.4 → ea95bdf
- k8s.io/api: v0.27.0 → v0.28.0
- k8s.io/apimachinery: v0.27.0 → v0.28.0
- k8s.io/client-go: v0.27.0 → v0.28.0
- k8s.io/component-base: v0.27.0 → v0.28.0
- k8s.io/klog/v2: v2.90.1 → v2.100.1
- k8s.io/kube-openapi: 15aac26 → 2695361
- k8s.io/utils: a36077c → d93618c

### Removed
- cloud.google.com/go/bigquery: v1.8.0
- cloud.google.com/go/datastore: v1.1.0
- cloud.google.com/go/pubsub: v1.3.1
- cloud.google.com/go/storage: v1.10.0
- dmitri.shuralyov.com/gpu/mtl: 666a987
- github.com/BurntSushi/xgb: [27f1227](https://github.com/BurntSushi/xgb/tree/27f1227)
- github.com/alecthomas/template: [fb15b89](https://github.com/alecthomas/template/tree/fb15b89)
- github.com/go-gl/glfw/v3.3/glfw: [6f7a984](https://github.com/go-gl/glfw/v3.3/glfw/tree/6f7a984)
- github.com/go-gl/glfw: [e6da0ac](https://github.com/go-gl/glfw/tree/e6da0ac)
- github.com/go-kit/kit: [v0.9.0](https://github.com/go-kit/kit/tree/v0.9.0)
- github.com/go-stack/stack: [v1.8.0](https://github.com/go-stack/stack/tree/v1.8.0)
- github.com/google/btree: [v1.0.0](https://github.com/google/btree/tree/v1.0.0)
- github.com/google/martian/v3: [v3.0.0](https://github.com/google/martian/v3/tree/v3.0.0)
- github.com/google/martian: [v2.1.0+incompatible](https://github.com/google/martian/tree/v2.1.0)
- github.com/google/renameio: [v0.1.0](https://github.com/google/renameio/tree/v0.1.0)
- github.com/googleapis/gax-go/v2: [v2.0.5](https://github.com/googleapis/gax-go/v2/tree/v2.0.5)
- github.com/hashicorp/golang-lru: [v0.5.1](https://github.com/hashicorp/golang-lru/tree/v0.5.1)
- github.com/jstemmer/go-junit-report: [v0.9.1](https://github.com/jstemmer/go-junit-report/tree/v0.9.1)
- github.com/kisielk/gotool: [v1.0.0](https://github.com/kisielk/gotool/tree/v1.0.0)
- github.com/konsorten/go-windows-terminal-sequences: [v1.0.3](https://github.com/konsorten/go-windows-terminal-sequences/tree/v1.0.3)
- github.com/kr/logfmt: [b84e30a](https://github.com/kr/logfmt/tree/b84e30a)
- github.com/kr/pty: [v1.1.1](https://github.com/kr/pty/tree/v1.1.1)
- github.com/kr/text: [v0.1.0](https://github.com/kr/text/tree/v0.1.0)
- github.com/sirupsen/logrus: [v1.6.0](https://github.com/sirupsen/logrus/tree/v1.6.0)
- go.opencensus.io: v0.22.4
- golang.org/x/image: cff245a
- golang.org/x/mobile: d2bd2a2
- google.golang.org/api: v0.30.0
- gopkg.in/alecthomas/kingpin.v2: v2.2.6
- gopkg.in/errgo.v2: v2.1.0
- rsc.io/binaryregexp: v0.2.0
- rsc.io/quote/v3: v3.1.0
- rsc.io/sampler: v1.3.0
