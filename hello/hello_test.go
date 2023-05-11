package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintsHelloMessageToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	hello.PrintTo(fakeTerminal)
	want := "Hello, World!"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
