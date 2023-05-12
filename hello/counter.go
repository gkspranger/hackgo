package hello

import (
	"fmt"
	"io"
	"os"
)

type Counter struct {
	Value         int
	Output        io.Writer
	MaxIterations int
}

func (c *Counter) Next() int {
	c.Value++
	return c.Value
}

func (c *Counter) Print() {
	fmt.Fprint(c.Output, c.Value, "\n")
}

func (c *Counter) Run() {
	for i := 0; i < c.MaxIterations; i++ {
		c.Next()
		c.Print()
	}
}

func NewCounter() *Counter {
	return &Counter{
		Value:         -1,
		Output:        os.Stdout,
		MaxIterations: 1_000_000_000_000,
	}
}
