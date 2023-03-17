package happyfun

import (
	"errors"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
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

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	val, ok := catalog[id]
	if ok {
		return val, nil
	}
	return Book{}, errors.New("book does not exist")
}

func MakeCatalog(catalog []Book) map[int]Book {
	var catalogMap = map[int]Book{}
	for _, b := range catalog {
		catalogMap[b.ID] = b
	}
	return catalogMap
}
