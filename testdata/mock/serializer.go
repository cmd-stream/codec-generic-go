package mock

import "github.com/ymz-ncnk/mok"

type (
	MarshalFn[T any]   func(t T) ([]byte, error)
	UnmarshalFn[V any] func(bs []byte, v V) error
	FormatFn           func() string
)

func NewSerializer[T, V any]() Serializer[T, V] {
	return Serializer[T, V]{mok.New("Serializer")}
}

type Serializer[T, V any] struct {
	*mok.Mock
}

func (m Serializer[T, V]) RegisterMarshal(fn MarshalFn[T]) Serializer[T, V] {
	m.Register("Marshal", fn)
	return m
}

func (m Serializer[T, V]) RegisterUnmarshal(fn UnmarshalFn[V]) Serializer[T, V] {
	m.Register("Unmarshal", fn)
	return m
}

func (m Serializer[T, V]) Marshal(t T) (bs []byte, err error) {
	result, err := m.Call("Marshal", t)
	if err != nil {
		panic(err)
	}
	bs = result[0].([]byte)
	err, _ = result[1].(error)
	return
}

func (m Serializer[T, V]) Unmarshal(bs []byte, v V) (err error) {
	result, err := m.Call("Unmarshal", bs, mok.SafeVal[V](v))
	if err != nil {
		panic(err)
	}
	err, _ = result[0].(error)
	return
}
