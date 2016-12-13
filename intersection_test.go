package godash_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/zillow/godash"
)

func TestIntersection(t *testing.T) {

	// test int success
	intSlice, err := godash.Intersection([]int{3, 17, 8, 11, 4, 8}, []int{11, 8, 5})
	intExpected := []int{8, 11}
	if err != nil {
		t.Errorf("Expected Intersection to return no error, but got %v", err)
	}
	if reflect.TypeOf(intSlice).Kind() != reflect.Slice {
		t.Error("Expected Intersection to return slice")
	}
	if !reflect.DeepEqual(intSlice, intExpected) {
		t.Errorf("Expected Intersection to return %v, but it returned %v", intExpected, intSlice)
	}

	// test struct success
	structSlice, err := godash.Intersection([]str{{name: "first"}, {name: "second"}, {name: "third"}}, []str{{name: "second"}, {name: "third"}, {name: "fourth"}})
	structExpected := []str{{name: "second"}, {name: "third"}}
	if err != nil {
		t.Errorf("Expected Intersection to return no error, but got %v", err)
	}
	if reflect.TypeOf(structSlice).Kind() != reflect.Slice {
		t.Error("Expected Intersection to return slice")
	}
	if !reflect.DeepEqual(structSlice, structExpected) {
		t.Errorf("Expected Intersection to return %v, but it returned %v", structExpected, structSlice)
	}

	// test failure
	result, err := godash.Intersection(1, []int{1, 2})
	if err == nil {
		t.Error("Expected Intersection to return error")
	}
	if result != nil {
		t.Errorf("Expected Intersection to return nil result, but got %v", result)
	}
	result, err = godash.Intersection([]int{1, 2}, 1)
	if err == nil {
		t.Error("Expected Intersection to return error")
	}
	if result != nil {
		t.Errorf("Expected Intersection to return nil result, but got %v", result)
	}
	result, err = godash.Intersection([]float32{1, 2}, []int{1, 2})
	if err == nil {
		t.Error("Expected Intersection to return error")
	}
	if result != nil {
		t.Errorf("Expected Intersection to return nil result, but got %v", result)
	}

}

func TestIntersectionBy(t *testing.T) {

	// test float success
	fn := func(x interface{}) interface{} {
		return math.Floor(x.(float64))
	}
	floatSlice, err := godash.IntersectionBy([]float64{2.16, 1.23, 5.4}, []float64{5.78, 2.49, 3.7, 2.3}, fn)
	floatExpected := []float64{2.16, 5.4}
	if err != nil {
		t.Errorf("Expected IntersectionBy to return no error, but got %v", err)
	}
	if reflect.TypeOf(floatSlice).Kind() != reflect.Slice {
		t.Error("Expected IntersectionBy to return slice")
	}
	if !reflect.DeepEqual(floatSlice, floatExpected) {
		t.Errorf("Expected IntersectionBy to return %v, but it returned %v", floatExpected, floatSlice)
	}

	// test struct success
	fn = func(x interface{}) interface{} {
		return x.(str).name
	}
	str1 := []str{{name: "apple", foo: "bar"}, {name: "orange", foo: "bar"}, {name: "apple", foo: ""}}
	str2 := []str{{name: "banana", foo: "barz"}, {name: "apple", foo: "barz"}}
	structSlice, err := godash.IntersectionBy(str1, str2, fn)
	structExpected := []str{{name: "apple", foo: "bar"}}
	if err != nil {
		t.Errorf("Expected IntersectionBy to return no error, but got %v", err)
	}
	if reflect.TypeOf(structSlice).Kind() != reflect.Slice {
		t.Error("Expected IntersectionBy to return slice")
	}
	if !reflect.DeepEqual(structSlice, structExpected) {
		t.Errorf("Expected IntersectionBy to return %v, but it returned %v", structExpected, structSlice)
	}

	// test failure
	fn = func(x interface{}) interface{} {
		return x
	}
	result, err := godash.IntersectionBy(1, []int{1, 2}, fn)
	if err == nil {
		t.Error("Expected IntersectionBy to return error")
	}
	if result != nil {
		t.Errorf("Expected IntersectionBy to return nil result, but got %v", result)
	}
	result, err = godash.IntersectionBy([]int{1, 2}, 1, fn)
	if err == nil {
		t.Error("Expected IntersectionBy to return error")
	}
	if result != nil {
		t.Errorf("Expected IntersectionBy to return nil result, but got %v", result)
	}
	result, err = godash.IntersectionBy([]float32{1, 2}, []int{1, 2}, fn)
	if err == nil {
		t.Error("Expected IntersectionBy to return error")
	}
	if result != nil {
		t.Errorf("Expected IntersectionBy to return nil result, but got %v", result)
	}

}
