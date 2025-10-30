package codec

import (
	"fmt"
	"reflect"

	com "github.com/mus-format/common-go"
)

// NewUnrecognizedType returns an error indicating that an unsupported type
// was encountered during encoding.
func NewUnrecognizedType(t reflect.Type) error {
	return fmt.Errorf("unrecognized type: %T", t)
}

// NewFailedToMarshalDTM returns an error indicating that the data type marker
// (DTM) could not be marshaled.
func NewFailedToMarshalDTM(err error) error {
	return fmt.Errorf("failed to marshal DTM: %w", err)
}

// NewFailedToMarshalValue returns an error indicating that value marshaling
// failed.
func NewFailedToMarshalValue(value any, cause error) error {
	return fmt.Errorf("failed to marshal %T value: %w", value, cause)
}

// NewFailedToMarshalByteSlice returns an error indicating that a byte slice
// could not be marshaled.
func NewFailedToMarshalByteSlice(err error) error {
	return fmt.Errorf("failed to marshal byte slice: %w", err)
}

// NewFailedToUnmarshalDTM returns an error indicating that the data type
// marker (DTM) could not be unmarshaled.
func NewFailedToUnmarshalDTM(err error) error {
	return fmt.Errorf("failed to unmarshal DTM: %w", err)
}

// NewUnrecognizedDTM returns an error indicating that an unknown data type
// marker (DTM) was received.
func NewUnrecognizedDTM(dtm com.DTM) error {
	return fmt.Errorf("unrecognized DTM: %v", dtm)
}

// NewFailedToUnmarshalByteSlice returns an error indicating that a byte slice
// could not be unmarshaled.
func NewFailedToUnmarshalByteSlice(err error) error {
	return fmt.Errorf("failed to unmarshal byte slice: %w", err)
}

// NewFailedToUnmarshalValue returns an error indicating that value
// unmarshaling failed.
func NewFailedToUnmarshalValue(err error) error {
	return fmt.Errorf("failed to unmarshal value: %w", err)
}
