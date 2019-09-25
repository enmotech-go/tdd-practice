package mocking

import (
	"bytes"
	"fmt"
)

func Countdown(b *bytes.Buffer) {
	fmt.Fprintf(b, "3")
}
