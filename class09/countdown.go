package class09

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
