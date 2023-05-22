package count_test

import (
	"bytes"
	"count"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()
	c, err := count.NewCounter(
		count.WithInput(bytes.NewBufferString("1\n2\n3")),
	)
	if err != nil {
		t.Fatal(err)
	}
	// c.Input = bytes.NewBufferString("1\n2\n3")
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
