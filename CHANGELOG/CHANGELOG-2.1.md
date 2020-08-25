# Release notes for v2.1.0

[Documentation](https://kubernetes-csi.github.io/docs/)

# Changelog since v2.0.0

## Changes by Kind

### Feature

- Add metrics endpoint. The metrics endpoint can be configured with `--metrics-address` and `--metrics-path` arguments. ([#70](https://github.com/kubernetes-csi/livenessprobe/pull/70), [@humblec](https://github.com/humblec))
- Make image, tag and registry configurable in docker file for windows build of CSI livenessprobe. ([#78](https://github.com/kubernetes-csi/livenessprobe/pull/78), [@mainred](https://github.com/mainred)) [SIG Windows]

### Bug or Regression

- Reports an unhealthy status on `/healthz` endpoint, if the CSI driver socket file does not exist anymore (accidentally removed on the host, for example). ([#69](https://github.com/kubernetes-csi/livenessprobe/pull/69), [@nettoclaudio](https://github.com/nettoclaudio))

### Uncategorized

- Build with Go 1.15 ([#79](https://github.com/kubernetes-csi/livenessprobe/pull/79), [@pohly](https://github.com/pohly))
- Publishing of images on k8s.gcr.io ([#71](https://github.com/kubernetes-csi/livenessprobe/pull/71), [@pohly](https://github.com/pohly))

## Dependencies

### Added
- github.com/Azure/go-autorest/autorest/adal: [v0.5.0](https://github.com/Azure/go-autorest/autorest/adal/tree/v0.5.0)
- github.com/Azure/go-autorest/autorest/date: [v0.1.0](https://github.com/Azure/go-autorest/autorest/date/tree/v0.1.0)
- github.com/Azure/go-autorest/autorest/mocks: [v0.2.0](https://github.com/Azure/go-autorest/autorest/mocks/tree/v0.2.0)
- github.com/Azure/go-autorest/autorest: [v0.9.0](https://github.com/Azure/go-autorest/autorest/tree/v0.9.0)
- github.com/Azure/go-autorest/logger: [v0.1.0](https://github.com/Azure/go-autorest/logger/tree/v0.1.0)
- github.com/Azure/go-autorest/tracing: [v0.5.0](https://github.com/Azure/go-autorest/tracing/tree/v0.5.0)
- github.com/NYTimes/gziphandler: [56545f4](https://github.com/NYTimes/gziphandler/tree/56545f4)
- github.com/PuerkitoBio/purell: [v1.0.0](https://github.com/PuerkitoBio/purell/tree/v1.0.0)
- github.com/PuerkitoBio/urlesc: [5bd2802](https://github.com/PuerkitoBio/urlesc/tree/5bd2802)
- github.com/alecthomas/template: [a0175ee](https://github.com/alecthomas/template/tree/a0175ee)
- github.com/alecthomas/units: [2efee85](https://github.com/alecthomas/units/tree/2efee85)
- github.com/beorn7/perks: [v1.0.0](https://github.com/beorn7/perks/tree/v1.0.0)
- github.com/blang/semver: [v3.5.0+incompatible](https://github.com/blang/semver/tree/v3.5.0)
- github.com/census-instrumentation/opencensus-proto: [v0.2.1](https://github.com/census-instrumentation/opencensus-proto/tree/v0.2.1)
- github.com/cncf/udpa/go: [269d4d4](https://github.com/cncf/udpa/go/tree/269d4d4)
- github.com/dgrijalva/jwt-go: [v3.2.0+incompatible](https://github.com/dgrijalva/jwt-go/tree/v3.2.0)
- github.com/docker/spdystream: [449fdfc](https://github.com/docker/spdystream/tree/449fdfc)
- github.com/elazarl/goproxy: [c4fc265](https://github.com/elazarl/goproxy/tree/c4fc265)
- github.com/emicklei/go-restful: [ff4f55a](https://github.com/emicklei/go-restful/tree/ff4f55a)
- github.com/envoyproxy/protoc-gen-validate: [v0.1.0](https://github.com/envoyproxy/protoc-gen-validate/tree/v0.1.0)
- github.com/evanphx/json-patch: [v4.5.0+incompatible](https://github.com/evanphx/json-patch/tree/v4.5.0)
- github.com/fsnotify/fsnotify: [v1.4.7](https://github.com/fsnotify/fsnotify/tree/v1.4.7)
- github.com/ghodss/yaml: [73d445a](https://github.com/ghodss/yaml/tree/73d445a)
- github.com/go-kit/kit: [v0.8.0](https://github.com/go-kit/kit/tree/v0.8.0)
- github.com/go-logfmt/logfmt: [v0.3.0](https://github.com/go-logfmt/logfmt/tree/v0.3.0)
- github.com/go-logr/logr: [v0.2.0](https://github.com/go-logr/logr/tree/v0.2.0)
- github.com/go-openapi/jsonpointer: [46af16f](https://github.com/go-openapi/jsonpointer/tree/46af16f)
- github.com/go-openapi/jsonreference: [13c6e35](https://github.com/go-openapi/jsonreference/tree/13c6e35)
- github.com/go-openapi/spec: [6aced65](https://github.com/go-openapi/spec/tree/6aced65)
- github.com/go-openapi/swag: [1d0bd11](https://github.com/go-openapi/swag/tree/1d0bd11)
- github.com/go-stack/stack: [v1.8.0](https://github.com/go-stack/stack/tree/v1.8.0)
- github.com/golang/groupcache: [5b532d6](https://github.com/golang/groupcache/tree/5b532d6)
- github.com/google/btree: [v1.0.0](https://github.com/google/btree/tree/v1.0.0)
- github.com/google/go-cmp: [v0.4.0](https://github.com/google/go-cmp/tree/v0.4.0)
- github.com/google/gofuzz: [v1.0.0](https://github.com/google/gofuzz/tree/v1.0.0)
- github.com/google/martian: [v2.1.0+incompatible](https://github.com/google/martian/tree/v2.1.0)
- github.com/google/pprof: [3ea8567](https://github.com/google/pprof/tree/3ea8567)
- github.com/google/uuid: [v1.1.1](https://github.com/google/uuid/tree/v1.1.1)
- github.com/googleapis/gax-go/v2: [v2.0.4](https://github.com/googleapis/gax-go/v2/tree/v2.0.4)
- github.com/googleapis/gnostic: [v0.2.0](https://github.com/googleapis/gnostic/tree/v0.2.0)
- github.com/gophercloud/gophercloud: [v0.1.0](https://github.com/gophercloud/gophercloud/tree/v0.1.0)
- github.com/gregjones/httpcache: [9cad4c3](https://github.com/gregjones/httpcache/tree/9cad4c3)
- github.com/hashicorp/golang-lru: [v0.5.1](https://github.com/hashicorp/golang-lru/tree/v0.5.1)
- github.com/hpcloud/tail: [v1.0.0](https://github.com/hpcloud/tail/tree/v1.0.0)
- github.com/imdario/mergo: [v0.3.5](https://github.com/imdario/mergo/tree/v0.3.5)
- github.com/json-iterator/go: [v1.1.8](https://github.com/json-iterator/go/tree/v1.1.8)
- github.com/jstemmer/go-junit-report: [af01ea7](https://github.com/jstemmer/go-junit-report/tree/af01ea7)
- github.com/julienschmidt/httprouter: [v1.2.0](https://github.com/julienschmidt/httprouter/tree/v1.2.0)
- github.com/kisielk/errcheck: [v1.2.0](https://github.com/kisielk/errcheck/tree/v1.2.0)
- github.com/kisielk/gotool: [v1.0.0](https://github.com/kisielk/gotool/tree/v1.0.0)
- github.com/konsorten/go-windows-terminal-sequences: [v1.0.2](https://github.com/konsorten/go-windows-terminal-sequences/tree/v1.0.2)
- github.com/kr/logfmt: [b84e30a](https://github.com/kr/logfmt/tree/b84e30a)
- github.com/kr/pretty: [v0.1.0](https://github.com/kr/pretty/tree/v0.1.0)
- github.com/kr/pty: [v1.1.1](https://github.com/kr/pty/tree/v1.1.1)
- github.com/kr/text: [v0.1.0](https://github.com/kr/text/tree/v0.1.0)
- github.com/kubernetes-csi/csi-test/v4: [555d70a](https://github.com/kubernetes-csi/csi-test/v4/tree/555d70a)
- github.com/mailru/easyjson: [d5b7844](https://github.com/mailru/easyjson/tree/d5b7844)
- github.com/matttproud/golang_protobuf_extensions: [v1.0.1](https://github.com/matttproud/golang_protobuf_extensions/tree/v1.0.1)
- github.com/modern-go/concurrent: [bacd9c7](https://github.com/modern-go/concurrent/tree/bacd9c7)
- github.com/modern-go/reflect2: [v1.0.1](https://github.com/modern-go/reflect2/tree/v1.0.1)
- github.com/munnerz/goautoneg: [a547fc6](https://github.com/munnerz/goautoneg/tree/a547fc6)
- github.com/mwitkow/go-conntrack: [cc309e4](https://github.com/mwitkow/go-conntrack/tree/cc309e4)
- github.com/mxk/go-flowrate: [cca7078](https://github.com/mxk/go-flowrate/tree/cca7078)
- github.com/onsi/ginkgo: [v1.10.3](https://github.com/onsi/ginkgo/tree/v1.10.3)
- github.com/onsi/gomega: [v1.7.1](https://github.com/onsi/gomega/tree/v1.7.1)
- github.com/peterbourgon/diskv: [v2.0.1+incompatible](https://github.com/peterbourgon/diskv/tree/v2.0.1)
- github.com/pkg/errors: [v0.8.1](https://github.com/pkg/errors/tree/v0.8.1)
- github.com/prometheus/client_golang: [v1.0.0](https://github.com/prometheus/client_golang/tree/v1.0.0)
- github.com/prometheus/client_model: [14fe0d1](https://github.com/prometheus/client_model/tree/14fe0d1)
- github.com/prometheus/common: [v0.4.1](https://github.com/prometheus/common/tree/v0.4.1)
- github.com/prometheus/procfs: [v0.0.2](https://github.com/prometheus/procfs/tree/v0.0.2)
- github.com/robertkrimen/otto: [c382bd3](https://github.com/robertkrimen/otto/tree/c382bd3)
- github.com/sirupsen/logrus: [v1.4.2](https://github.com/sirupsen/logrus/tree/v1.4.2)
- github.com/spf13/afero: [v1.2.2](https://github.com/spf13/afero/tree/v1.2.2)
- github.com/spf13/pflag: [v1.0.5](https://github.com/spf13/pflag/tree/v1.0.5)
- go.opencensus.io: v0.21.0
- golang.org/x/time: 9d24e82
- golang.org/x/xerrors: 9bdfabe
- google.golang.org/api: v0.4.0
- google.golang.org/protobuf: v1.23.0
- gopkg.in/alecthomas/kingpin.v2: v2.2.6
- gopkg.in/fsnotify.v1: v1.4.7
- gopkg.in/inf.v0: v0.9.1
- gopkg.in/sourcemap.v1: v1.0.5
- gopkg.in/tomb.v1: dd63297
- k8s.io/api: v0.17.0
- k8s.io/apimachinery: v0.17.1-beta.0
- k8s.io/client-go: v0.17.0
- k8s.io/component-base: v0.17.0
- k8s.io/gengo: 0689ccc
- k8s.io/klog/v2: v2.3.0
- k8s.io/kube-openapi: 30be4d1
- k8s.io/utils: e782cd3
- rsc.io/quote/v3: v3.1.0
- rsc.io/sampler: v1.3.0
- sigs.k8s.io/structured-merge-diff: 15d366b
- sigs.k8s.io/yaml: v1.1.0

### Changed
- cloud.google.com/go: v0.26.0 → v0.38.0
- github.com/container-storage-interface/spec: [v1.1.0 → v1.3.0](https://github.com/container-storage-interface/spec/compare/v1.1.0...v1.3.0)
- github.com/davecgh/go-spew: [v1.1.0 → v1.1.1](https://github.com/davecgh/go-spew/compare/v1.1.0...v1.1.1)
- github.com/envoyproxy/go-control-plane: [v0.6.9 → v0.9.4](https://github.com/envoyproxy/go-control-plane/compare/v0.6.9...v0.9.4)
- github.com/gogo/protobuf: [v1.2.0 → 65acae2](https://github.com/gogo/protobuf/compare/v1.2.0...65acae2)
- github.com/golang/mock: [v1.2.0 → v1.4.3](https://github.com/golang/mock/compare/v1.2.0...v1.4.3)
- github.com/golang/protobuf: [v1.3.1 → v1.4.2](https://github.com/golang/protobuf/compare/v1.3.1...v1.4.2)
- github.com/kubernetes-csi/csi-lib-utils: [v0.6.1 → v0.7.0](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.6.1...v0.7.0)
- github.com/stretchr/objx: [v0.1.0 → v0.1.1](https://github.com/stretchr/objx/compare/v0.1.0...v0.1.1)
- github.com/stretchr/testify: [v1.4.0 → v1.5.1](https://github.com/stretchr/testify/compare/v1.4.0...v1.5.1)
- golang.org/x/crypto: c2843e0 → 60c769a
- golang.org/x/net: eb5bcb5 → c0dbc17
- golang.org/x/oauth2: d2e6202 → 0f29369
- golang.org/x/sync: e225da7 → 1122301
- golang.org/x/sys: 9773273 → 0732a99
- golang.org/x/text: v0.3.0 → v0.3.3
- golang.org/x/tools: 1195517 → 2c0ae70
- google.golang.org/appengine: v1.4.0 → v1.5.0
- google.golang.org/genproto: 64821d5 → 5c49e3e
- google.golang.org/grpc: v1.20.0 → v1.29.0
- gopkg.in/check.v1: 20d25e2 → 788fd78
- gopkg.in/yaml.v2: v2.2.2 → v2.2.5
- honnef.co/go/tools: c2f93a9 → ea95bdf
- k8s.io/klog: v0.3.0 → v1.0.0

### Removed
- github.com/gogo/googleapis: [v1.1.0](https://github.com/gogo/googleapis/tree/v1.1.0)
- github.com/kubernetes-csi/csi-test: [v2.0.0+incompatible](https://github.com/kubernetes-csi/csi-test/tree/v2.0.0)
- github.com/lyft/protoc-gen-validate: [v0.0.13](https://github.com/lyft/protoc-gen-validate/tree/v0.0.13)
