package fizzbuzz

import (
	"fmt"
	"io"
)

const (
	WordFizz     = "Fizz"
	WordBuzz     = "Buzz"
	WordFizzBuzz = WordFizz + WordBuzz
)

func Converter(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return WordFizzBuzz
	} else if i%3 == 0 {
		return WordFizz
	} else if i%5 == 0 {
		return WordBuzz
	} else {
		return fmt.Sprintf("%d", i)
	}
}

func FizzBuzz(w io.Writer, n int) {
	for i := 1; i < n+1; i++ {
		fmt.Fprintf(w, "%s\n", Converter(i))
	}
}
