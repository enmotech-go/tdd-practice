package fizzbuzz

import "fmt"

type FizzBuzz struct {
	input int
}

func (fb FizzBuzz) Number() string {
	if isDivisibleBy(fb.input, 3) && isDivisibleBy(fb.input, 5) {
		return "FizzBuzz"
	}
	if isDivisibleBy(fb.input, 3) {
		return "Fizz"
	}
	if isDivisibleBy(fb.input, 5) {
		return "Buzz"
	}
	return fmt.Sprintf("%d", fb.input)
}

func isDivisibleBy(input, n int) bool {
	return input%n == 0
}
