package creditcard_test

import (
	"happyfun/creditcard"
	"testing"
)

func TestNumber(t *testing.T) {
	t.Parallel()
	want := "123"
	c, err := creditcard.New(want)

	if err != nil {
		t.Fatal(err)
	}

	got := c.Number()

	if want != got {
		t.Errorf("started with %v, wanted %v, got %v", c, want, got)
	}
}

func TestNumberInvalid(t *testing.T) {
	t.Parallel()
	_, err := creditcard.New("")

	if err == nil {
		t.Fatal("this should have failed because number was empty")
	}
}
