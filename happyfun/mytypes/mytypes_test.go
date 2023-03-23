package mytypes_test

import (
	"happyfun/mytypes"
	"strings"
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

func TestStringsBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := sb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d", sb.String(),
			wantLen, gotLen)
	}
}
