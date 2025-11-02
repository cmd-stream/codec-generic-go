package codec_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/cmd-stream/codec-generic-go"
	"github.com/cmd-stream/codec-generic-go/testdata/mock"
	tmocks "github.com/cmd-stream/testkit-go/mocks/transport"
	com "github.com/mus-format/common-go"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestCodec(t *testing.T) {
	t.Run("Encoding should work", func(t *testing.T) {
		var (
			wantDTM = 0
			wantBs  = []byte{1, 2, 3}
			wantLen = len(wantBs)
			wantN   = 1 + 1 + wantLen
		)

		ser := mock.NewSerializer[MyInterface, MyInterface]().RegisterMarshal(
			func(t MyInterface) ([]byte, error) {
				return wantBs, nil
			})
		c := codec.NewCodec(
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			ser,
		)

		w := tmocks.NewWriter().RegisterWriteByte(func(b byte) error {
			assertfatal.Equal(b, byte(wantDTM), t)
			return nil
		}).RegisterWriteByte(func(b byte) error {
			assertfatal.Equal(b, byte(wantLen), t)
			return nil
		}).RegisterWrite(func(p []byte) (n int, err error) {
			assertfatal.EqualDeep(p, wantBs, t)
			return len(p), nil
		})

		n, err := c.Encode(MyStruct1{}, w)
		assertfatal.EqualError(err, nil, t)
		assertfatal.Equal(n, wantN, t)
	})

	t.Run("Failed to marshal DTM", func(t *testing.T) {
		c := codec.NewCodec[MyInterface, MyInterface](
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			nil,
		)

		writeErr := errors.New("failed to write DTM")
		wantErr := codec.NewFailedToMarshalDTM(writeErr)

		w := tmocks.NewWriter().RegisterWriteByte(func(b byte) error {
			return writeErr
		})
		n, err := c.Encode(MyStruct1{}, w)
		assertfatal.EqualError(err, wantErr, t)
		assertfatal.Equal(n, 0, t)
	})

	t.Run("Decoding should work", func(t *testing.T) {
		wantDTM := 1
		wantV := MyStruct2{Y: "hello"}
		wantBs := []byte{1, 2, 3}
		wantLen := len(wantBs)
		wantN := 1 + 1 + wantLen

		ser := mock.NewSerializer[MyInterface, MyInterface]().RegisterUnmarshal(
			func(bs []byte, v MyInterface) error {
				assertfatal.EqualDeep(bs, wantBs, t)

				rv := reflect.ValueOf(v)
				if rv.Kind() == reflect.Ptr && !rv.IsNil() {
					rv.Elem().Set(reflect.ValueOf(wantV))
				}

				return nil
			},
		)
		c := codec.NewCodec(
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			ser,
		)

		r := tmocks.NewReader().RegisterReadByte(func() (b byte, err error) {
			return byte(wantDTM), nil
		}).RegisterReadByte(func() (b byte, err error) {
			return byte(wantLen), nil
		}).RegisterRead(func(p []byte) (n int, err error) {
			copy(p, wantBs)
			return wantLen, nil
		})

		v, n, err := c.Decode(r)
		assertfatal.EqualError(err, nil, t)
		assertfatal.Equal(n, wantN, t)
		assertfatal.EqualDeep(v, wantV, t)
	})

	t.Run("Unrecognized type", func(t *testing.T) {
		c := codec.NewCodec[MyInterface, MyInterface](
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			nil,
		)

		v := MyStruct3{Z: 3.14}
		wantType := reflect.TypeOf(v)
		wantErr := codec.NewUnrecognizedType(wantType)

		w := tmocks.NewWriter()

		n, err := c.Encode(v, w)
		assertfatal.EqualError(err, wantErr, t)
		assertfatal.Equal(n, 0, t)
	})

	t.Run("Failed to marshal byte slice", func(t *testing.T) {
		ser := mock.NewSerializer[MyInterface, MyInterface]().RegisterMarshal(
			func(t MyInterface) ([]byte, error) {
				return []byte{}, nil
			},
		)
		c := codec.NewCodec(
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			ser,
		)

		writeErr := errors.New("failed to write byte slice length")
		wantErr := codec.NewFailedToMarshalByteSlice(writeErr)

		w := tmocks.NewWriter().RegisterWriteByte(func(b byte) error {
			return nil
		}).RegisterWriteByte(func(b byte) error {
			return writeErr
		})

		n, err := c.Encode(MyStruct1{}, w)
		assertfatal.EqualError(err, wantErr, t)
		assertfatal.Equal(n, 1, t)
	})

	t.Run("Failed to unmarshal DTM", func(t *testing.T) {
		c := codec.NewCodec[MyInterface, MyInterface](
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			nil,
		)

		readErr := errors.New("failed to read DTM")
		wantErr := codec.NewFailedToUnmarshalDTM(readErr)

		r := tmocks.NewReader().RegisterReadByte(func() (b byte, err error) {
			return 0, readErr
		})

		_, n, err := c.Decode(r)
		assertfatal.EqualError(err, wantErr, t)
		assertfatal.Equal(n, 0, t)
	})

	t.Run("Unrecognized DTM", func(t *testing.T) {
		c := codec.NewCodec[MyInterface, MyInterface](
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			nil,
		)

		const unrecognizedDTM com.DTM = 99
		wantErr := codec.NewUnrecognizedDTM(unrecognizedDTM)

		r := tmocks.NewReader().RegisterReadByte(func() (b byte, err error) {
			return byte(unrecognizedDTM), nil
		})

		_, n, err := c.Decode(r)
		assertfatal.EqualError(err, wantErr, t)
		assertfatal.Equal(n, 1, t)
	})

	t.Run("Failed to unmarshal byte slice", func(t *testing.T) {
		c := codec.NewCodec[MyInterface, MyInterface](
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			[]reflect.Type{
				reflect.TypeFor[MyStruct1](),
				reflect.TypeFor[MyStruct2](),
			},
			nil,
		)

		readErr := errors.New("failed to read byte slice")
		wantErr := codec.NewFailedToUnmarshalByteSlice(readErr)

		r := tmocks.NewReader().RegisterReadByte(func() (b byte, err error) {
			return 0, nil
		}).RegisterReadByte(func() (b byte, err error) {
			return 0, readErr
		})

		_, n, err := c.Decode(r)
		assertfatal.EqualError(err, wantErr, t)
		assertfatal.Equal(n, 1, t)
	})
}
