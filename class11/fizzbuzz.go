package class11

import "strconv"

type fizzBuzz struct {
	Input int
}

func (f *fizzBuzz) Number() string {
	return strconv.Itoa(f.Input)

}

func FizzBuzz(num int) *fizzBuzz {
	buzz := new(fizzBuzz)
	buzz.Input = num
	return buzz
}
