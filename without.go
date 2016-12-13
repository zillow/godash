package godash

import (
	"errors"
	"reflect"
)

// Without removes values from a slice and returns the new slice.
// It accepts a slice of any type as the first parameter, followed by a list of parameter values to remove from the slice.
// The additional values must be of the same type as the provided slice.
// If using a basic type, such as string or int, it is recommended to use the more specific functions, such as WithoutString or WithoutInt.
// Otherwise, if using this function directly, the returned result will need to have a type assertion applied.
func Without(slice interface{}, values ...interface{}) (interface{}, error) {

	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Type().Kind() != reflect.Slice {
		return nil, errors.New("godash: invalid parameter type. Without func expects parameter 1 to be a slice")
	}
	for _, v := range values {
		if sliceVal.Type().Elem() != reflect.TypeOf(v) {
			return nil, errors.New("godash: invalid parameter type. Without func expects additional parameters to match the type of the provided slice")
		}
	}

	dest := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(slice).Elem()), 0, sliceVal.Len())

	for i := 0; i < sliceVal.Len(); i++ {
		remove := false
		for _, v := range values {
			if sliceVal.Index(i).Interface() == v {
				remove = true
				break
			}
		}
		if !remove {
			dest = reflect.Append(dest, sliceVal.Index(i))
		}
	}
	return dest.Interface(), nil

}

// WithoutString removes string values from a string slice
func WithoutString(slice []string, values ...interface{}) ([]string, error) {

	result, err := Without(slice, values...)
	if err != nil {
		return nil, err
	}
	return result.([]string), nil

}

// WithoutInt removes int values from an int slice
func WithoutInt(slice []int, values ...interface{}) ([]int, error) {

	result, err := Without(slice, values...)
	if err != nil {
		return nil, err
	}
	return result.([]int), nil

}

// WithoutInt8 removes int8 values from an int8 slice
func WithoutInt8(slice []int8, values ...interface{}) ([]int8, error) {

	result, err := Without(slice, values...)
	if err != nil {
		return nil, err
	}
	return result.([]int8), nil

}

// WithoutFloat32 removes float32 values from a float32 slice
func WithoutFloat32(slice []float32, values ...interface{}) ([]float32, error) {

	result, err := Without(slice, values...)
	if err != nil {
		return nil, err
	}
	return result.([]float32), nil

}

// WithoutBy removes values from a slice based on output from a provided validator function and returns the new slice.
// The supplied function must accept an interface{} parameter and return bool.
// Values for which the validator function returns true will be removed from the slice.
func WithoutBy(slice interface{}, fn validator) (interface{}, error) {

	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Type().Kind() != reflect.Slice {
		return nil, errors.New("godash: invalid parameter type. WithoutBy func expects parameter 1 to be a slice")
	}

	dest := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(slice).Elem()), 0, sliceVal.Len())

	for i := 0; i < sliceVal.Len(); i++ {
		v := sliceVal.Index(i).Interface()
		if remove := fn(v); !remove {
			dest = reflect.Append(dest, sliceVal.Index(i))
		}
	}
	return dest.Interface(), nil

}
