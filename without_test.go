package godash_test

import (
	"reflect"
	"testing"

	"github.com/zillow/godash"
)

func TestWithout(t *testing.T) {

	stringSource := []string{"one", "two", "three", "four", "five", "six"}
	intSource := []int{1, 2, 3, 4, 5, 6}
	structSource := []str{{name: "first"}, {name: "second"}, {name: "third"}}

	// test for string success
	stringDest, err := godash.Without(stringSource, "two", "four")
	stringExpected := []string{"one", "three", "five", "six"}
	if err != nil {
		t.Errorf("Expected Without to return no error, but got %v", err)
	}
	if reflect.TypeOf(stringDest).Kind() != reflect.Slice {
		t.Error("Expected Without to return slice")
	}
	if !reflect.DeepEqual(stringDest, stringExpected) {
		t.Errorf("Expected Without to return %v, but it returned %v", stringExpected, stringDest)
	}

	// test for int success
	intDest, err := godash.Without(intSource, 2, 4)
	intExpected := []int{1, 3, 5, 6}
	if err != nil {
		t.Errorf("Expected Without to return no error, but got %v", err)
	}
	if reflect.TypeOf(intDest).Kind() != reflect.Slice {
		t.Error("Expected Without to return slice")
	}
	if !reflect.DeepEqual(intDest, intExpected) {
		t.Errorf("Expected Without to return %v, but it returned %v", intExpected, intDest)
	}

	// test for struct success
	structDest, err := godash.Without(structSource, str{name: "second"})
	structExpected := []str{{name: "first"}, {name: "third"}}
	if err != nil {
		t.Errorf("Expected Without to return no error, but got %v", err)
	}
	if reflect.TypeOf(structDest).Kind() != reflect.Slice {
		t.Error("Expected Without to return slice")
	}
	if !reflect.DeepEqual(structDest, structExpected) {
		t.Errorf("Expected Without to return %v, but it returned %v", structExpected, structDest)
	}

	// test for failure
	failDest, err := godash.Without(stringSource, 1, 2)
	if err == nil {
		t.Error("Expected Without to return error")
	}
	if failDest != nil {
		t.Errorf("Expected Without to return nil result, but got %v", failDest)
	}
	failDest, err = godash.Without(1, 2, 3)
	if err == nil {
		t.Error("Expected Without to return error")
	}
	if failDest != nil {
		t.Errorf("Expected Without to return nil result, but got %v", failDest)
	}

	// test for empty slice
	var emptySource, emptyVals []interface{}
	emptyDest, err := godash.Without(emptySource, emptyVals...)
	if err != nil {
		t.Errorf("Expected Without to return no error, but got %v", err)
	}
	if reflect.TypeOf(emptyDest).Kind() != reflect.Slice {
		t.Error("Expected Without to return slice")
	}
	if reflect.ValueOf(emptyDest).Len() > 0 {
		t.Errorf("Expected Without to return empty slice, but got %v", emptyDest)
	}

}

func TestWithoutString(t *testing.T) {

	source := []string{"one", "two", "three", "four", "five", "six"}
	expected := []string{"one", "three", "five", "six"}

	dest, err := godash.WithoutString(source, "two", "four")

	if err != nil {
		t.Errorf("Expected WithoutString to return no error, but got %v", err)
	}
	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("Expected WithoutString to return %v, but it returned %v", expected, dest)
	}

}

func TestWithoutInt(t *testing.T) {

	source := []int{1, 2, 3, 4, 5, 6}
	expected := []int{1, 3, 5, 6}

	dest, err := godash.WithoutInt(source, 2, 4)

	if err != nil {
		t.Errorf("Expected WithoutInt to return no error, but got %v", err)
	}
	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("Expected WithoutInt to return %v, but it returned %v", expected, dest)
	}

}

func TestWithoutInt8(t *testing.T) {

	source := []int8{1, 2, 3, 4, 5, 6}
	expected := []int8{1, 3, 5, 6}

	dest, err := godash.WithoutInt8(source, int8(2), int8(4))

	if err != nil {
		t.Errorf("Expected WithoutInt8 to return no error, but got %v", err)
	}
	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("Expected WithoutInt8 to return %v, but it returned %v", expected, dest)
	}

}

func TestWithoutFloat32(t *testing.T) {

	source := []float32{1.4, 2.0, 3.3, 4.7, 5.1, 6.0}
	expected := []float32{1.4, 3.3, 5.1, 6.0}

	dest, err := godash.WithoutFloat32(source, float32(2.0), float32(4.7))

	if err != nil {
		t.Errorf("Expected WithoutFloat32 to return no error, but got %v", err)
	}
	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("Expected WithoutFloat32 to return %v, but it returned %v", expected, dest)
	}

}

func TestWithoutBy(t *testing.T) {

	source := []int{1, 2, 3, 4, 5, 6}
	fn := func(x interface{}) bool {
		i := x.(int)
		return i%2 == 0
	}
	expected := []int{1, 3, 5}

	// test for success
	successDest, err := godash.WithoutBy(source, fn)
	if err != nil {
		t.Errorf("Expected WithoutBy to return no error, but got %v", err)
	}
	if !reflect.DeepEqual(successDest, expected) {
		t.Errorf("Expected WithoutBy to return %v, but it returned %v", expected, successDest)
	}

	// test for failure
	failDest, err := godash.WithoutBy(1, fn)
	if err == nil {
		t.Error("Expected WithoutBy to return error")
	}
	if failDest != nil {
		t.Errorf("Expected WithoutBy to return nil result, but got %v", failDest)
	}

}
