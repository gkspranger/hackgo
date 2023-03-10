package happyfun

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) Book {
	b.Copies--
	return b
}

func Return(b Book) Book {
	b.Copies++
	return b
}
