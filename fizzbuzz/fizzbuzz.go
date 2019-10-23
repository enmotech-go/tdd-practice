package fizzbuzz

import (
	"fmt"
	"strings"
)

func FizzBuzz(s int64) string {
	if IsRelatedTo(s, 3) && IsRelatedTo(s, 5) {
		return "fizzbuzz"
	} else if IsRelatedTo(s, 3) {
		return "fizz"
	} else if IsRelatedTo(s, 5) {
		return "buzz"
	}
	return fmt.Sprintf("%d", s)
}

func IsRelatedTo(s, n int64) bool {
	if s%n == 0 {
		return true
	}
	return strings.Contains(fmt.Sprintf("%d", s), fmt.Sprintf("%d", n))
}
