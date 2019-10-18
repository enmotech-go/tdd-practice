package fizzbuzz

import (
	"fmt"
	"io"
)

const (
	WordFizz = "Fizz"
	WordBuzz = "Buzz"
)

func Converter(i int) string {
	var result string
	if i%3 == 0 {
		result += WordFizz
	}
	if i%5 == 0 {
		result += WordBuzz
	}
	if result == "" {
		result += fmt.Sprintf("%d", i)
	}
	return result
}

func FizzBuzz(w io.Writer, n int) {
	for i := 1; i < n+1; i++ {
		fmt.Fprintf(w, "%s\n", Converter(i))
	}
}
