package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

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

func TestPrintsNextMessageToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	c := hello.NewCounter()
	c.Output = fakeTerminal
	c.MaxIterations = 1
	c.Sleep = 0
	c.Run()
	want := "0\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestPrintsManyNextMessageToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	c := hello.NewCounter()
	c.Output = fakeTerminal
	c.MaxIterations = 3
	c.Sleep = 0
	c.Run()
	want := "0\n1\n2\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
