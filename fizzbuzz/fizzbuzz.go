package fizzbuzz

import "fmt"

const (
	WordFizz = "Fizz"
	WordBuzz = "Buzz"
)

type FizzBuzz struct{}

func (FizzBuzz) Convert(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return WordFizz + WordBuzz
	}
	if n%3 == 0 {
		return WordFizz
	}
	if n%5 == 0 {
		return WordBuzz
	}
	return fmt.Sprintf("%d", n)
}
