package happyfun_test

import (
	"happyfun"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = happyfun.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := happyfun.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	want := 1
	result, err := happyfun.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("started with %d, wanted %d, got %d", b.Copies, want, got)
	}
}

func TestBuyErrorIfOutOfStock(t *testing.T) {
	t.Parallel()
	b := happyfun.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := happyfun.Buy(b)
	if err == nil {
		t.Errorf("started with %d, should have reported as out of stock", b.Copies)
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []happyfun.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}

	want := catalog
	got := happyfun.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
