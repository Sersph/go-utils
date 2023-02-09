package verify_method

import (
	"reflect"
	"strings"
)

type RequiredMetadata struct {
	Zero bool
	Trim bool
}

type Required struct {
	Metadata *RequiredMetadata
}

func (r Required) Verify(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return !(!r.Metadata.Zero && value.Int() == 0)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return !(!r.Metadata.Zero && value.Uint() == 0)
	case reflect.Float32, reflect.Float64:
		return !(!r.Metadata.Zero && value.Float() == 0)
	case reflect.Array, reflect.Slice:
		return value.Len() != 0
	case reflect.String:
		if r.Metadata.Trim {
			return strings.Trim(value.String(), " ") != ""
		}
		return value.String() != ""
	case reflect.Ptr:
		return value.IsNil()
	}

	return false
}
