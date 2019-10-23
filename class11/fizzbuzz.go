package class11

import "strconv"

type fizzBuzz struct {
	Input int
}

func (f *fizzBuzz) Number() string {
	if f.Input%3 == 0 {
		return "Fizz"
	}
	if f.Input%5 == 0 {
		return "Buzz"
	}
	if f.Input%5 == 0 && f.Input%3 == 0 {
		return "FizzBuzz"
	}
	return strconv.Itoa(f.Input)

}

func FizzBuzz(num int) *fizzBuzz {
	buzz := new(fizzBuzz)
	buzz.Input = num
	return buzz
}
