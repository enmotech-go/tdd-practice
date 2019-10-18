package fizzbuzz

import "fmt"

const (
	WordFizz = "Fizz"
	WordBuzz = "Buzz"
)

type FizzBuzz struct{}

func (FizzBuzz) Convert(n int) string {
	var result string
	if n%3 == 0 {
		result += WordFizz
	}
	if n%5 == 0 {
		result += WordBuzz
	}
	if result == "" {
		result = fmt.Sprintf("%d", n)
	}
	return result
}
