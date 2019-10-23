package fizzbuzz

import "fmt"

type FizzBuzz struct {
	input int
}

func (fb FizzBuzz)Number()  string {
	if fb.input%3==0&&fb.input%5==0{
		return "FizzBuzz"
	}
	if fb.input%3==0{
		return "Fizz"
	}
	if fb.input%5==0{
		return "Buzz"
	}
	return fmt.Sprintf("%d",fb.input)
}

