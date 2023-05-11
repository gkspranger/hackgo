package happyfun_test

import (
	"happyfun"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
	sort.Slice(want, func(i, j int) bool {
		return want[i].ID < want[j].ID
	})
	got := happyfun.GetAllBooks(happyfun.MakeCatalog(want))

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(happyfun.Book{})) {
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
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(happyfun.Book{})) {
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

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(happyfun.Book{})) {
		t.Error(cmp.Diff(want, got))
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := happyfun.Book{
		Author:          "greg",
		Title:           "wawa",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	var want = 3000
	got := b.NetPriceCents()

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := happyfun.Book{
		Author:          "greg",
		Title:           "wawa",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	err := b.SetPriceCents(3000)

	if err != nil {
		t.Fatal(err)
	}

	if want != b.PriceCents {
		t.Errorf("wanted %v, got %v", want, b.PriceCents)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := happyfun.Book{
		Author:          "greg",
		Title:           "wawa",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	err := b.SetPriceCents(-1)

	if err == nil {
		t.Fatal("an error should have been returned for a bad value")
	}
}

func TestBookSetCategory(t *testing.T) {
	t.Parallel()
	b := happyfun.Book{
		Author:          "greg",
		Title:           "wawa",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	cats := []happyfun.Category{
		happyfun.CategoryBiography,
		happyfun.CategoryRomance,
		happyfun.CategoryScience,
	}

	for _, cat := range cats {
		want := cat
		err := b.SetCategory(cat)
		got := b.Category()

		if err != nil {
			t.Fatal(err)
		}

		if want != got {
			t.Errorf("wanted %v, got %v", want, got)
		}
	}
}

func TestBookSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := happyfun.Book{
		Author:          "greg",
		Title:           "wawa",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	var category happyfun.Category = -1
	err := b.SetCategory(category)

	if err == nil {
		t.Errorf("should have failed using category %v", category)
	}
}

func TestBigger(t *testing.T) {
	t.Parallel()
	want := 2
	got := happyfun.Bigger(1, 2)

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestXorTrue(t *testing.T) {
	t.Parallel()
	want := true
	got := happyfun.Xor(true, false)

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestXorFalse(t *testing.T) {
	t.Parallel()
	want := false

	if got := happyfun.Xor(false, false); want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}

	if got := happyfun.Xor(true, true); want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestGreetBob(t *testing.T) {
	t.Parallel()
	want := "hello bob"
	got := happyfun.Greet("bob")

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestGreetAlice(t *testing.T) {
	t.Parallel()
	want := "hey alice"
	got := happyfun.Greet("alice")

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestGreetDefault(t *testing.T) {
	t.Parallel()
	want := "hi stranger"
	got := happyfun.Greet("greg")

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestTotal(t *testing.T) {
	t.Parallel()
	want := 6
	got := happyfun.Total([]int{1, 2, 3})

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestEvens(t *testing.T) {
	t.Parallel()
	happyfun.Evens()
}

func TestNonNegative(t *testing.T) {
	t.Parallel()
	want := []int{0, 2}
	got := happyfun.NonNegative([]int{-1, 0, -4, 2})

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestWithdrawValid(t *testing.T) {
	t.Parallel()
	want := 9
	got, err := happyfun.Withdraw(10, 1)

	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestWithdrawIValid(t *testing.T) {
	t.Parallel()
	_, err := happyfun.Withdraw(10, 11)

	if err == nil {
		t.Fatal(err)
	}
}

func TestApply(t *testing.T) {
	t.Parallel()
	want := 2
	got := happyfun.Apply(1, func(x int) int {
		return x * 2
	})

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestAddManyZero(t *testing.T) {
	t.Parallel()
	want := 0
	got := happyfun.AddMany()

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestAddManyOne(t *testing.T) {
	t.Parallel()
	want := 1
	got := happyfun.AddMany(1)

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestAddManyThree(t *testing.T) {
	t.Parallel()
	want := 6
	got := happyfun.AddMany(1, 2, 3)

	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}
