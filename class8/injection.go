package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(buffer io.Writer, name string) {
	fmt.Fprintf(buffer, "Hello ,%s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
