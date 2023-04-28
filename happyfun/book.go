package happyfun

import (
	"errors"
	"fmt"
	"sort"
)

type Category int

const (
	CategoryScience Category = iota
	CategoryRomance
	CategoryBiography
)

var validCategory = map[Category]bool{
	CategoryBiography: true,
	CategoryRomance:   true,
	CategoryScience:   true,
}

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        Category
}

func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving

}

func (b *Book) SetPriceCents(price int) error {
	if price <= 0 {
		return fmt.Errorf("the price was less <= 0")
	}
	b.PriceCents = price
	return nil
}

func (b *Book) SetCategory(category Category) error {
	if !validCategory[category] {
		return fmt.Errorf("category is invalid")
	}
	b.category = category
	return nil
}

func (b Book) Category() Category {
	return b.category
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("book is out of stock")
	}
	b.Copies--
	return b, nil
}

func Return(b Book) Book {
	b.Copies++
	return b
}

func sortBooksByID(catalog []Book) []Book {
	sort.Slice(catalog, func(i, j int) bool {
		return catalog[i].ID < catalog[j].ID
	})
	return catalog
}

func GetAllBooks(catalogMap map[int]Book) []Book {
	var books = []Book{}
	for _, b := range catalogMap {
		books = append(books, b)
	}
	return sortBooksByID(books)
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	val, ok := catalog[id]
	if ok {
		return val, nil
	}
	return Book{}, fmt.Errorf("ID %d doesn't exist", id)
}

func MakeCatalog(catalog []Book) map[int]Book {
	var catalogMap = map[int]Book{}
	for _, b := range catalog {
		catalogMap[b.ID] = b
	}
	return catalogMap
}

func NetPriceCents(book Book) int {
	saving := book.PriceCents * book.DiscountPercent / 100
	return book.PriceCents - saving
}
