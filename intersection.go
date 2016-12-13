package godash

import (
	"errors"
	"reflect"
)

// Intersection creates a slice of unique values that were present in both of the provided slices.
// The order of the items in the resulting slice is determined by the first given slice.
// The new slice is returned as an interface{} and may need to have a type assertion applied to it afterwards.
func Intersection(slice1 interface{}, slice2 interface{}) (interface{}, error) {

	sliceVal1 := reflect.ValueOf(slice1)
	sliceVal2 := reflect.ValueOf(slice2)

	if sliceVal1.Type().Kind() != reflect.Slice {
		return nil, errors.New("godash: invalid parameter type. Intersection func expects parameter 1 to be a slice")
	}
	if sliceVal2.Type().Kind() != reflect.Slice {
		return nil, errors.New("godash: invalid parameter type. Intersection func expects parameter 2 to be a slice")
	}
	if sliceVal1.Type().Elem() != sliceVal2.Type().Elem() {
		return nil, errors.New("godash: invalid parameter type. Intersection func expects two slice parameters of the same type")
	}

	dest := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(slice1).Elem()), 0, sliceVal1.Len())
	m := make(map[interface{}]bool)

	for i := 0; i < sliceVal2.Len(); i++ {
		val := sliceVal2.Index(i).Interface()
		m[val] = false
	}
	for i := 0; i < sliceVal1.Len(); i++ {
		val := sliceVal1.Index(i).Interface()
		appended, exists := m[val]
		if exists {
			if !appended {
				dest = reflect.Append(dest, sliceVal1.Index(i))
			}
			m[val] = true
		}
	}
	return dest.Interface(), nil

}

// IntersectionBy passes items from two provided slices through a provided mutator function and creates a new slice with items that resulted in common mutated values.
// The supplied mutator function must accept an interface{} parameter and return interface{} with the value to be compared.
// The order and values of the items in the resulting slice are determined by the first given slice.
// The new slice is returned as an interface{} and may need to have a type assertion applied to it afterwards.
func IntersectionBy(slice1 interface{}, slice2 interface{}, fn mutator) (interface{}, error) {

	sliceVal1 := reflect.ValueOf(slice1)
	sliceVal2 := reflect.ValueOf(slice2)

	if sliceVal1.Type().Kind() != reflect.Slice {
		return nil, errors.New("godash: invalid parameter type. IntersectionBy func expects parameter 1 to be a slice")
	}
	if sliceVal2.Type().Kind() != reflect.Slice {
		return nil, errors.New("godash: invalid parameter type. IntersectionBy func expects parameter 2 to be a slice")
	}
	if sliceVal1.Type().Elem() != sliceVal2.Type().Elem() {
		return nil, errors.New("godash: invalid parameter type. IntersectionBy func expects two slice parameters of the same type")
	}

	dest := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(slice1).Elem()), 0, sliceVal1.Len())
	m := make(map[interface{}]bool)

	for i := 0; i < sliceVal2.Len(); i++ {
		item := sliceVal2.Index(i).Interface()
		val := fn(item)
		m[val] = false
	}

	for i := 0; i < sliceVal1.Len(); i++ {
		item := sliceVal1.Index(i).Interface()
		val := fn(item)
		appended, exists := m[val]
		if exists {
			if !appended {
				dest = reflect.Append(dest, sliceVal1.Index(i))
			}
			m[val] = true
		}
	}
	return dest.Interface(), nil

}
