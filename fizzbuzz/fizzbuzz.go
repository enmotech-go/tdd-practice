package fizzbuzz

import "fmt"

type FizzBuzz struct {
	input int
}

func (fb FizzBuzz)Number()  string {
	if fb.input%3==0{
		return "Fizz"
	}
	return fmt.Sprintf("%d",fb.input)
}

