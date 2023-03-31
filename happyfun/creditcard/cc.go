package creditcard

import (
	"errors"
)

type creditcard struct {
	number string
}

func New(number string) (creditcard, error) {
	if number == "" {
		return creditcard{}, errors.New("CC number is not valid because it is empty")
	}
	return creditcard{number}, nil
}

func (c creditcard) Number() string {
	return c.number
}
