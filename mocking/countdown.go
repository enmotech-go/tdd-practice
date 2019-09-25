package mocking

import (
	"fmt"
	"io"
)

func Countdown(b io.Writer) {
	fmt.Fprintf(b, "3")
}
