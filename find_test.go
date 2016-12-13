package godash_test

import (
	"testing"

	"github.com/zillow/godash"
)

func TestFindBy(t *testing.T) {

	intFn := func(x interface{}) bool {
		i := x.(int)
		return i > 4
	}
	sliceFn := func(x interface{}) bool {
		s := x.(str)
		return s.name == "second" || s.name == "third"
	}

	// test for int success
	val, err := godash.FindBy([]int{1, 2, 3, 4, 5, 6}, intFn)
	if err != nil {
		t.Errorf("Expected FindBy to return no error, but got %v", err)
	}
	if val != 5 {
		t.Errorf("Expected FindBy to return %v, but it returned %v", 5, val)
	}

	// test for struct success
	val, err = godash.FindBy([]str{{name: "first"}, {name: "second"}, {name: "third"}}, sliceFn)
	expectedStruct := str{name: "second"}
	if err != nil {
		t.Errorf("Expected FindBy to return no error, but got %v", err)
	}
	if val != expectedStruct {
		t.Errorf("Expected FindBy to return %v, but it returned %v", expectedStruct, val)
	}

	// test for not found
	val, err = godash.FindBy([]int{1, 2, 3, 1, 2, 3}, intFn)
	if err != nil {
		t.Errorf("Expected FindBy to return no error, but got %v", err)
	}
	if val != nil {
		t.Errorf("Expected FindBy to return no value, but it returned %v", val)
	}

	// test for failure
	val, err = godash.FindBy(str{name: "value"}, sliceFn)
	if err == nil {
		t.Error("Expected FindBy to return error")
	}
	if val != nil {
		t.Errorf("Expected FindBy to return no value, but it returned %v", val)
	}

}

func TestFindLastBy(t *testing.T) {

	intFn := func(x interface{}) bool {
		i := x.(int)
		return i < 6
	}
	sliceFn := func(x interface{}) bool {
		s := x.(str)
		return s.name == "second" || s.name == "third"
	}

	// test for int success
	val, err := godash.FindLastBy([]int{1, 2, 3, 4, 5, 6}, intFn)
	if err != nil {
		t.Errorf("Expected FindLastBy to return no error, but got %v", err)
	}
	if val != 5 {
		t.Errorf("Expected FindLastBy to return %v, but it returned %v", 5, val)
	}

	// test for struct success
	val, err = godash.FindLastBy([]str{{name: "first"}, {name: "second"}, {name: "third"}}, sliceFn)
	expectedStruct := str{name: "third"}
	if err != nil {
		t.Errorf("Expected FindLastBy to return no error, but got %v", err)
	}
	if val != expectedStruct {
		t.Errorf("Expected FindLastBy to return %v, but it returned %v", expectedStruct, val)
	}

	// test for not found
	val, err = godash.FindLastBy([]int{6, 7, 8}, intFn)
	if err != nil {
		t.Errorf("Expected FindLastBy to return no error, but got %v", err)
	}
	if val != nil {
		t.Errorf("Expected FindLastBy to return no value, but it returned %v", val)
	}

	// test for failure
	val, err = godash.FindLastBy(str{name: "value"}, sliceFn)
	if err == nil {
		t.Error("Expected FindLastBy to return error")
	}
	if val != nil {
		t.Errorf("Expected FindLastBy to return no value, but it returned %v", val)
	}

}

func TestFindIndex(t *testing.T) {

	stringSource := []string{"one", "two", "three", "four", "five", "six"}
	intSource := []int{1, 2, 3, 1, 2, 3}
	structSource := []str{{name: "first"}, {name: "second"}, {name: "third"}}

	// test for string success
	index, err := godash.FindIndex(stringSource, "three")
	if err != nil {
		t.Errorf("Expected FindIndex to return no error, but got %v", err)
	}
	if index != 2 {
		t.Errorf("Expected FindIndex to return %v, but it returned %v", 2, index)
	}

	// test for int success
	index, err = godash.FindIndex(intSource, 1)
	if err != nil {
		t.Errorf("Expected FindIndex to return no error, but got %v", err)
	}
	if index != 0 {
		t.Errorf("Expected FindIndex to return %v, but it returned %v", 0, index)
	}
	index, err = godash.FindIndex(intSource, 3)
	if err != nil {
		t.Errorf("Expected FindIndex to return no error, but got %v", err)
	}
	if index != 2 {
		t.Errorf("Expected FindIndex to return %v, but it returned %v", 2, index)
	}

	// test for struct success
	index, err = godash.FindIndex(structSource, str{name: "second"})
	if err != nil {
		t.Errorf("Expected FindIndex to return no error, but got %v", err)
	}
	if index != 1 {
		t.Errorf("Expected FindIndex to return %v, but it returned %v", 1, index)
	}

	// test for not found
	index, err = godash.FindIndex(intSource, 8)
	if err != nil {
		t.Errorf("Expected FindIndex to return no error, but got %v", err)
	}
	if index != -1 {
		t.Errorf("Expected FindIndex to return %v, but it returned %v", -1, index)
	}

	// test for failure
	index, err = godash.FindIndex(str{name: "value"}, str{name: "value"})
	if err == nil {
		t.Error("Expected FindIndex to return error")
	}
	if index != -1 {
		t.Errorf("Expected FindIndex to return %v, but it returned %v", -1, index)
	}

}

func TestFindIndexBy(t *testing.T) {

	fn := func(x interface{}) bool {
		i := x.(int)
		return i > 4
	}

	// test for success
	index, err := godash.FindIndexBy([]int{1, 2, 3, 4, 5, 6}, fn)
	if err != nil {
		t.Errorf("Expected FindIndexBy to return no error, but got %v", err)
	}
	if index != 4 {
		t.Errorf("Expected FindIndexBy to return %v, but it returned %v", 4, index)
	}

	// test for failure
	index, err = godash.FindIndexBy(5, fn)
	if err == nil {
		t.Error("Expected FindIndexBy to return error")
	}
	if index != -1 {
		t.Errorf("Expected FindIndexBy to return %v, but it returned %v", -1, index)
	}

	// test for not found
	index, err = godash.FindIndexBy([]int{1, 2, 3, 1, 2, 3}, fn)
	if err != nil {
		t.Errorf("Expected FindIndexBy to return no error, but got %v", err)
	}
	if index != -1 {
		t.Errorf("Expected FindIndexBy to return %v, but it returned %v", -1, index)
	}

}

func TestFindLastIndex(t *testing.T) {

	stringSource := []string{"one", "two", "three", "one", "two", "three"}
	intSource := []int{1, 2, 3, 1, 2, 3}
	structSource := []str{{name: "first"}, {name: "second"}, {name: "third"}}

	// test for string success
	index, err := godash.FindLastIndex(stringSource, "three")
	if err != nil {
		t.Errorf("Expected FindLastIndex to return no error, but got %v", err)
	}
	if index != 5 {
		t.Errorf("Expected FindLastIndex to return %v, but it returned %v", 5, index)
	}

	// test for int success
	index, err = godash.FindLastIndex(intSource, 1)
	if err != nil {
		t.Errorf("Expected FindLastIndex to return no error, but got %v", err)
	}
	if index != 3 {
		t.Errorf("Expected FindLastIndex to return %v, but it returned %v", 3, index)
	}

	// test for struct success
	index, err = godash.FindLastIndex(structSource, str{name: "second"})
	if err != nil {
		t.Errorf("Expected FindLastIndex to return no error, but got %v", err)
	}
	if index != 1 {
		t.Errorf("Expected FindLastIndex to return %v, but it returned %v", 1, index)
	}

	// test for not found
	index, err = godash.FindLastIndex(intSource, 8)
	if err != nil {
		t.Errorf("Expected FindLastIndex to return no error, but got %v", err)
	}
	if index != -1 {
		t.Errorf("Expected FindLastIndex to return %v, but it returned %v", -1, index)
	}

	// test for failure
	index, err = godash.FindLastIndex(str{name: "value"}, str{name: "value"})
	if err == nil {
		t.Error("Expected FindLastIndex to return error")
	}
	if index != -1 {
		t.Errorf("Expected FindLastIndex to return %v, but it returned %v", -1, index)
	}

}
