package searchme_test

import (
	"bytes"
	"searchme"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPass(t *testing.T) {
	t.Parallel()
	want := 1
	got := 1
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSearch(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	wHello :=
		[]string{
			"hello world",
			"i am hello",
		}
	woHello :=
		[]string{
			"i dont have that word",
			"i refuse to say that word",
		}

	// want := strings.Join(wHello[:], "\n")

	input := append(wHello, woHello...)

	search := "hello"

	s, err := searchme.NewSearch(
		searchme.WithInput(bytes.NewBufferString(strings.Join(input[:], "\n"))),
		searchme.WithOutput(fakeTerminal),
	)
	if err != nil {
		t.Fatal(err)
	}

	s.Search(search)

	gotStr := strings.TrimSpace(fakeTerminal.String())
	got := strings.Split(gotStr, "\n")

	if !cmp.Equal(wHello, got) {
		t.Error(cmp.Diff(wHello, got))
	}
}
