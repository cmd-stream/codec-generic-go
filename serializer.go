package codec

// Serializer defines the interface for encoding and decoding values.
//
// T is the type used for marshaling (encoding), and V is the type used for
// unmarshaling (decoding).
type Serializer[T, V any] interface {
	Marshal(t T) (bs []byte, err error)
	Unmarshal(bs []byte, v V) (err error)
}
