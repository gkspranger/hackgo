package hello

import (
	"bytes"
	"fmt"
)

func PrintTo(buf *bytes.Buffer) {
	fmt.Fprint(buf, "Hello, World!")
}
