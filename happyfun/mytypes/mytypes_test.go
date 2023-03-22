package mytypes_test

import (
	"happyfun/mytypes"
	"testing"
)

func TestMyIntTwice(t *testing.T) {
	t.Parallel()
	var want mytypes.MyInt = 4
	input := mytypes.MyInt(2)
	got := input.Twice()

	if want != got {
		t.Errorf("input %v; wanted %v; got %v", input, want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	input := mytypes.MyString("greg")
	want := 4
	got := input.Len()

	if want != got {
		t.Errorf("input %v; wanted %v; got %v", input, want, got)
	}
}
