package mocking

import (
	"bytes"
	"fmt"
)

func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}