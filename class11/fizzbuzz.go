package class11

type fizzBuzz struct {
	Input string
}

func (f *fizzBuzz) Number() string {

	return "1"
}

func FizzBuzz(num int) *fizzBuzz {
	buzz := new(fizzBuzz)
	return buzz
}
