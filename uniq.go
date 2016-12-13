package godash

import (
	"errors"
	"reflect"
)

// Uniq removes duplicate values from a slice and returns the new slice.
// The new slice is returned as an interface{} and may need to have a type assertion applied to it afterwards.
func Uniq(slice interface{}) (interface{}, error) {

	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Type().Kind() != reflect.Slice {
		return nil, errors.New("godash: invalid parameter type. Uniq func expects parameter 1 to be a slice")
	}

	dest := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(slice).Elem()), 0, sliceVal.Len())
	m := make(map[interface{}]bool)

	for i := 0; i < sliceVal.Len(); i++ {
		val := sliceVal.Index(i).Interface()
		_, appended := m[val]
		if !appended {
			dest = reflect.Append(dest, sliceVal.Index(i))
			m[val] = true
		}
	}
	return dest.Interface(), nil

}
