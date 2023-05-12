package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintsHelloMessageToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	p := hello.NewPrinter()
	p.Output = fakeTerminal
	p.Print()
	want := "Hello, World!\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
