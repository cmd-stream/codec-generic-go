# codec-generic-go

[![Go Reference](https://pkg.go.dev/badge/github.com/cmd-stream/codec-generic-go.svg)](https://pkg.go.dev/github.com/cmd-stream/codec-generic-go)
[![GoReportCard](https://goreportcard.com/badge/cmd-stream/codec-generic-go)](https://goreportcard.com/report/github.com/cmd-stream/codecs-generic-go)
[![codecov](https://codecov.io/gh/cmd-stream/codec-generic-go/graph/badge.svg?token=GVcvey0TbG)](https://codecov.io/gh/cmd-stream/codec-generic-go)

**codec-generic-go** provides the generic codec abstraction that can be
used by concrete [cmd-stream-go](https://github.com/cmd-stream/cmd-stream-go) codecs.

It defines the common [Codec](./codec.go) structure independent of any specific
serialization format.

## Used By

- [codec-json-go](https://github.com/cmd-stream/codec-json-go)
- [codec-protobuf-go](https://github.com/cmd-stream/codec-protobuf-go)
