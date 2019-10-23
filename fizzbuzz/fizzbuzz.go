package fizzbuzz

import "fmt"

func FizzBuzz(s int64) string {
	if s%15 == 0 {
		return "fizzbuzz"
	} else if s%3 == 0 {
		return "fizz"
	} else if s%5 == 0 {
		return "buzz"
	}
	return fmt.Sprintf("%d", s)
}
