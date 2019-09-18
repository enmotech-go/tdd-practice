package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello,%s", name)
}

func main() {
	Greet(os.Stdout, "errio")
}
