package godash

import (
	"errors"
	"reflect"
)

// FindBy returns the first element of the slice that the provided validator function returns true for.
// The supplied function must accept an interface{} parameter and return bool.
// If the validator function does not return true for any values in the slice, nil is returned.
func FindBy(slice interface{}, fn validator) (interface{}, error) {

	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Type().Kind() != reflect.Slice {
		return nil, errors.New("godash: invalid parameter type. FindBy func expects parameter 1 to be a slice")
	}

	for i := 0; i < sliceVal.Len(); i++ {
		val := sliceVal.Index(i).Interface()
		if match := fn(val); match == true {
			return val, nil
		}
	}
	return nil, nil

}

// FindLastBy returns the last element of the slice that the provided validator function returns true for.
// The supplied function must accept an interface{} parameter and return bool.
// If the validator function does not return true for any values in the slice, nil is returned.
func FindLastBy(slice interface{}, fn validator) (interface{}, error) {

	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Type().Kind() != reflect.Slice {
		return nil, errors.New("godash: invalid parameter type. FindBy func expects parameter 1 to be a slice")
	}

	for i := sliceVal.Len() - 1; i != -1; i-- {
		val := sliceVal.Index(i).Interface()
		if match := fn(val); match == true {
			return val, nil
		}
	}
	return nil, nil

}

// FindIndex returns the index of the first element in a slice that equals the provided value.
// If the value is not found in the slice, -1 is returned.
func FindIndex(slice interface{}, value interface{}) (int, error) {

	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Type().Kind() != reflect.Slice {
		return -1, errors.New("godash: invalid parameter type. FindIndex func expects parameter 1 to be a slice")
	}

	for i := 0; i < sliceVal.Len(); i++ {
		if sliceVal.Index(i).Interface() == value {
			return i, nil
		}
	}
	return -1, nil

}

// FindIndexBy returns the index of the first element of a slice that the provided validator function returns true for.
// The supplied function must accept an interface{} parameter and return bool.
// If the validator function does not return true for any values in the slice, -1 is returned.
func FindIndexBy(slice interface{}, fn validator) (int, error) {

	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Type().Kind() != reflect.Slice {
		return -1, errors.New("godash: invalid parameter type. FindIndexBy func expects parameter 1 to be a slice")
	}

	for i := 0; i < sliceVal.Len(); i++ {
		if match := fn(sliceVal.Index(i).Interface()); match == true {
			return i, nil
		}
	}
	return -1, nil

}

// FindLastIndex returns the index of the last element in a slice that equals the provided value.
// If the value is not found in the slice, -1 is returned.
func FindLastIndex(slice interface{}, value interface{}) (int, error) {

	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Type().Kind() != reflect.Slice {
		return -1, errors.New("godash: invalid parameter type. FindLastIndex func expects parameter 1 to be a slice")
	}

	for i := sliceVal.Len() - 1; i != -1; i-- {
		if sliceVal.Index(i).Interface() == value {
			return i, nil
		}
	}
	return -1, nil

}
