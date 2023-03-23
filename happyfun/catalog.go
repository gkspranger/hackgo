package happyfun

import (
	"fmt"
	"sort"
)

type Catalog map[int]Book

func (c Catalog) GetAllBooks() []Book {
	books := []Book{}
	for _, book := range c {
		books = append(books, book)
	}
	sort.Slice(books, func(i, j int) bool {
		return books[i].ID < books[j].ID
	})

	return books
}

func (c Catalog) GetBook(ID int) (Book, error) {
	val, ok := c[ID]
	if ok {
		return val, nil
	}
	return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
}
