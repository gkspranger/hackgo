package mytypes

import "strings"

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (i *MyInt) Double() {
	// *i *= 2
	*i = *i * 2
}

type MyString string

func (s MyString) Len() int {
	return len(s)
}

type MyStrBuilder strings.Builder

func (m MyStrBuilder) Hello() string {
	return "Hello, Gophers!"
}

type MyBuilder struct {
	Contents strings.Builder
}

type StringUppercaser struct {
	Contents strings.Builder
}

func (s StringUppercaser) ToUpper() string {
	return strings.ToUpper(s.Contents.String())
}

func Double(input int) int {
	return input * 2
}

func DoubleP(input *int) {
	*input *= 2
}
