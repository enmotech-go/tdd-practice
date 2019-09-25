package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
