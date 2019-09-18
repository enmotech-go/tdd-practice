package injection

import (
	"bytes"
	"fmt"
)

func Greet(buffer *bytes.Buffer, name string) {
	fmt.Fprintf(buffer, "Hello ,%s", name)
}
