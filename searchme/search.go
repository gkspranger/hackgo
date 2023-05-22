package searchme

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type search struct {
	input  io.Reader
	output io.Writer
	// search string
}

func (s *search) Search(text string) {
	scanner := bufio.NewScanner(s.input)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			fmt.Fprint(s.output, scanner.Text(), "\n")
		}
	}
}

type option func(*search) error

func NewSearch(opts ...option) (search, error) {
	s := search{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(&s)
		if err != nil {
			return search{}, err
		}
	}
	return s, nil
}

func WithInput(input io.Reader) option {
	return func(s *search) error {
		if input == nil {
			return errors.New("input was nil")
		}
		s.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(s *search) error {
		if output == nil {
			return errors.New("output was nil")
		}
		s.output = output
		return nil
	}
}
