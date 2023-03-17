package happyfun_test

import (
	"happyfun"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testableCatalog() []happyfun.Book {
	return []happyfun.Book{
		{ID: 1, Title: "For the Love of Go", Author: "greg", Copies: 3},
		{ID: 2, Title: "The Power of Go: Tools", Author: "bob", Copies: 1},
		{ID: 3, Title: "Some other bookb", Author: "billy", Copies: 7},
	}
}

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
	want := testableCatalog()
	got := happyfun.GetAllBooks(testableCatalog())

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookValid(t *testing.T) {
	t.Parallel()
	catalogMap := happyfun.MakeCatalog(testableCatalog())

	var book = happyfun.Book{}

	for k := range catalogMap {
		book = catalogMap[k]
		break
	}

	want := book
	got, err := happyfun.GetBook(catalogMap, book.ID)

	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestGetBookInvalid(t *testing.T) {
	t.Parallel()
	catalogMap := happyfun.MakeCatalog(testableCatalog())

	_, err := happyfun.GetBook(catalogMap, 999)

	if err == nil {
		t.Errorf("started with %v, should have reported as not existing", catalogMap)
	}
}

func TestMakeCatalogMap(t *testing.T) {
	t.Parallel()
	catalog := testableCatalog()
	var want = map[int]happyfun.Book{}
	want[catalog[0].ID] = catalog[0]
	want[catalog[1].ID] = catalog[1]
	want[catalog[2].ID] = catalog[2]

	got := happyfun.MakeCatalog(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
		t.Errorf("wanted %v, got %v", want, got)
	}
}
