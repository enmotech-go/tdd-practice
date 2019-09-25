package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(b io.Writer) {
	fmt.Fprintf(b, "3")
}

func main() {
	Countdown(os.Stdout)
}
