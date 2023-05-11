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

func TestCounterWithOneNext(t *testing.T) {
	t.Parallel()
	c := hello.NewCounter()
	want := 0
	got := c.Next()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestCounterWithManyNext(t *testing.T) {
	t.Parallel()
	c := hello.NewCounter()
	want := 2
	c.Next()
	c.Next()
	got := c.Next()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestCounterWithSeedAndManyNext(t *testing.T) {
	t.Parallel()
	c := hello.NewCounter()
	c.Value = 100
	want := 103
	c.Next()
	c.Next()
	got := c.Next()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
