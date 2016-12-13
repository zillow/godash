package godash_test

import (
	"reflect"
	"testing"

	"github.com/zillow/godash"
)

func TestUniq(t *testing.T) {

	// test for string success
	stringSlice, err := godash.Uniq([]string{"apple", "orange", "apple", "banana", "orange"})
	stringExpected := []string{"apple", "orange", "banana"}
	if err != nil {
		t.Errorf("Expected Uniq to return no error, but got %v", err)
	}
	if reflect.TypeOf(stringSlice).Kind() != reflect.Slice {
		t.Error("Expected Uniq to return slice")
	}
	if !reflect.DeepEqual(stringSlice, stringExpected) {
		t.Errorf("Expected Uniq to return %v, but it returned %v", stringExpected, stringSlice)
	}

	// test for struct success
	structSlice, err := godash.Uniq([]str{{name: "apple", foo: "bar"}, {name: "orange", foo: "bar"}, {name: "apple", foo: "-"}, {name: "orange", foo: "bar"}})
	structExpected := []str{{name: "apple", foo: "bar"}, {name: "orange", foo: "bar"}, {name: "apple", foo: "-"}}
	if err != nil {
		t.Errorf("Expected Uniq to return no error, but got %v", err)
	}
	if reflect.TypeOf(structSlice).Kind() != reflect.Slice {
		t.Error("Expected Uniq to return slice")
	}
	if !reflect.DeepEqual(structSlice, structExpected) {
		t.Errorf("Expected Uniq to return %v, but it returned %v", structExpected, structSlice)
	}

	// test for no changes
	intSlice, err := godash.Uniq([]int{1, 2, 3, 4, 5})
	intExpected := []int{1, 2, 3, 4, 5}
	if err != nil {
		t.Errorf("Expected Uniq to return no error, but got %v", err)
	}
	if reflect.TypeOf(intSlice).Kind() != reflect.Slice {
		t.Error("Expected Uniq to return slice")
	}
	if !reflect.DeepEqual(intSlice, intExpected) {
		t.Errorf("Expected Uniq to return %v, but it returned %v", intExpected, intSlice)
	}

	// test for failure
	fail, err := godash.Uniq(str{name: "one"})
	if err == nil {
		t.Error("Expected Uniq to return error")
	}
	if fail != nil {
		t.Errorf("Expected Uniq to return nil result, but got %v", fail)
	}

}
