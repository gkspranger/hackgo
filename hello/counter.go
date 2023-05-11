package hello

type Counter struct {
	Value int
}

func (c *Counter) Next() int {
	c.Value++
	return c.Value
}

func NewCounter() *Counter {
	return &Counter{
		Value: -1,
	}
}
