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
		return "fizzbuzz"
	}
	if isRelatedTo(fb.input, 3) {
		return "fizz"
	}
	if isRelatedTo(fb.input, 5) {
		return "buzz"
	}
	return fmt.Sprintf("%d", fb.input)
}

func isRelatedTo(input, n int) bool {
	if input%n == 0 || strings.Contains(fmt.Sprintf("%d", input), fmt.Sprintf("%d", n)) {
		return true
	}
	return false
}
