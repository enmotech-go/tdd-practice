package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}
