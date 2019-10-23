package fizzbuzz

import "fmt"

type FizzBuzz struct {
	input int
}

func (fb FizzBuzz) Number() string {
	return fmt.Sprintf("%d", fb.input)
}
