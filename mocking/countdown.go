package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(b io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(b, i)
	}
	fmt.Fprintf(b, "go!")
}

func main() {
	Countdown(os.Stdout)
}
