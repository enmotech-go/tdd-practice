package main

import (
	"fmt"
	"io"
	"os"
)

const finalword = "go!"
const countstart = 3

func Countdown(b io.Writer) {
	for i := countstart; i > 0; i-- {
		fmt.Fprintln(b, i)
	}
	fmt.Fprintf(b, finalword)
}

func main() {
	Countdown(os.Stdout)
}
