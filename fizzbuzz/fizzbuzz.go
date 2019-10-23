package fizzbuzz

import (
	"fmt"
	"strings"
)

type FizzBuzz struct {
	input int
}

func (fb FizzBuzz) Number() string {
	if isRelatedTo(fb.input, 3) && isRelatedTo(fb.input, 5) {
		return "FizzBuzz"
	}
	if isRelatedTo(fb.input, 3) {
		return "Fizz"
	}
	if isRelatedTo(fb.input, 5) {
		return "Buzz"
	}
	return fmt.Sprintf("%d", fb.input)
}

func isRelatedTo(input, n int) bool {
	return input%n == 0 || strings.Contains(fmt.Sprintf("%d", input), fmt.Sprintf("%d", n))
}
