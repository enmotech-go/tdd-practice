package class11

import "strconv"

type fizzBuzz struct {
	Input int
}

const (
	Fizz        = "Fizz"
	Buzz        = "Buzz"
	FizzBuzzStr = "FizzBuzz"
)

func (f *fizzBuzz) Number() string {
	if f.Input%5 == 0 && f.Input%3 == 0 {
		return FizzBuzzStr
	}
	if f.Input%3 == 0 {
		return Fizz
	}
	if f.Input%5 == 0 {
		return Buzz
	}
	return strconv.Itoa(f.Input)

}

func FizzBuzz(num int) *fizzBuzz {
	buzz := new(fizzBuzz)
	buzz.Input = num
	return buzz
}
