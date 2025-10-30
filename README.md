# codec-generic-go

**codec-generic-go** provides the generic codec abstraction that can be
used by concrete [cmd-stream-go](https://github.com/cmd-stream/cmd-stream-go) codecs.

It defines the common [Codec](./codec.go) structure independent of any specific
serialization format.

## Used By

- [codec-json-go](https://github.com/cmd-stream/codec-json-go)
- [codec-protobuf-go](https://github.com/cmd-stream/codec-protobuf-go)
