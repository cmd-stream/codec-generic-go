package codec

import "reflect"

type DecodeValueFn[T, V any] func(tp reflect.Type, ser Serializer[T, V],
	bs []byte) (v V, err error)

func decodeValue[T, V any](tp reflect.Type, ser Serializer[T, V], bs []byte) (
	v V, err error,
) {
	ptr := reflect.New(tp)
	err = ser.Unmarshal(bs, ptr.Interface().(V))
	if err != nil {
		err = NewFailedToUnmarshalValue(err)
		return
	}
	v = ptr.Elem().Interface().(V)
	return
}
