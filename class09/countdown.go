package class09

import (
	"bytes"
	"fmt"
)

func Countdown(buffer *bytes.Buffer) {
	fmt.Fprint(buffer, "3")
}
