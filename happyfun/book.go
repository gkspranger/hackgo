package happyfun

import (
	"errors"
)

type Book struct {
	Title  string
	Author string
	Copies int
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

func GetAllBooks(c []Book) []Book {
	return c
}
