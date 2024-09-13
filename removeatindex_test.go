package removeatindex

import (
	"reflect"
	"testing"
)

func TestRemoveAtIndex(t *testing.T) {
	newSlice := Slice{1, "two", 3, "bar"}
	got := RemoveAtIndex(newSlice, 2)
	want := Slice{1, "two", "bar"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveAtIndex() = %v, want %v", got, want)
	}
}
