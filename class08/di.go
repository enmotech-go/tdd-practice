package class08

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}
