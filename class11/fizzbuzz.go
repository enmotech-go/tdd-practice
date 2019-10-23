package class11

import (
	"strconv"
	"strings"
)

type fizzBuzz struct {
	Input int
}

const (
	Fizz        = "Fizz"
	Buzz        = "Buzz"
	FizzBuzzStr = "FizzBuzz"
)

func (f *fizzBuzz) Number() string {
	if isRelated(f.Input, 3) && isRelated(f.Input, 5) {
		return FizzBuzzStr
	}
	if isRelated(f.Input, 3) {
		return Fizz
	}
	if isRelated(f.Input, 5) {
		return Buzz
	}
	return strconv.Itoa(f.Input)

}

func FizzBuzz(num int) *fizzBuzz {
	buzz := new(fizzBuzz)
	buzz.Input = num
	return buzz
}

func isRelated(input, num int) bool {
	return input%num == 0 || strings.Contains(strconv.Itoa(input), strconv.Itoa(num))

}
