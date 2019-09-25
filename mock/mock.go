package mock

import (
	"bytes"
	"fmt"
)

func CountDown(out *bytes.Buffer) {
	fmt.Fprintf(out, "3")
}
