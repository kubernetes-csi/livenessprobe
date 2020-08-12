# Changes since v1.1.0

## Action Required

- Introduce V(5) on the health check begin/success log lines to allow filtering of these entries from logs. If you would like to retain these log entries the action required would be to set `-v==5` or higher for the livenessprobe container. ([#57](https://github.com/kubernetes-csi/livenessprobe/pull/57), [@stefansedich](https://github.com/stefansedich))
- Deprecated "--connection-timeout" argument has been removed. ([#59](https://github.com/kubernetes-csi/livenessprobe/pull/59), [@msau42](https://github.com/msau42))

## Other Notable Changes

- Fix nil pointer bug when driver responds with not ready ([#58](https://github.com/kubernetes-csi/livenessprobe/pull/58), [@scuzhanglei](https://github.com/scuzhanglei))
- Migrated to Go modules, so the source builds also outside of GOPATH. ([#53](https://github.com/kubernetes-csi/livenessprobe/pull/53), [@pohly](https://github.com/pohly))


