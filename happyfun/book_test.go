package happyfun_test

import (
	"happyfun"
	"testing"
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
	result := happyfun.Buy(b)
	got := result.Copies
	if want != got {
		t.Errorf("started with %d, wanted %d, got %d", b.Copies, want, got)
	}
}
