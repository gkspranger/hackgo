package happyfun_test

import (
	"happyfun"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func books() []happyfun.Book {
	return []happyfun.Book{
		{ID: 1, Title: "For the Love of Go", Author: "greg", Copies: 3},
		{ID: 2, Title: "The Power of Go: Tools", Author: "bob", Copies: 1},
		{ID: 3, Title: "Some other bookb", Author: "billy", Copies: 7},
	}
}

func makeCatalog() happyfun.Catalog {
	catalog := happyfun.Catalog{}
	for _, v := range books() {
		catalog[v.ID] = v
	}
	return catalog
}

func TestCatalogGetAllBooks(t *testing.T) {
	t.Parallel()
	var input = makeCatalog()
	want := books()
	got := input.GetAllBooks()

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCatalogGetBookValid(t *testing.T) {
	t.Parallel()
	catalog := makeCatalog()
	want := books()[0]

	got, err := catalog.GetBook(want.ID)

	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestCatalogGetBookInvalid(t *testing.T) {
	t.Parallel()
	catalog := makeCatalog()

	_, err := catalog.GetBook(999)

	if err == nil {
		t.Errorf("started with %v, should have reported as not existing", catalog)
	}
}
