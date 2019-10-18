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

func FizzBuzz(w io.Writer, n int) {
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Fprintf(w, "%s\n", WordFizzBuzz)
		} else if i%3 == 0 {
			fmt.Fprintf(w, "%s\n", WordFizz)
		} else if i%5 == 0 {
			fmt.Fprintf(w, "%s\n", WordBuzz)
		} else {
			fmt.Fprintf(w, "%d\n", i)
		}
	}
}
