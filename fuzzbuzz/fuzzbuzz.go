package fuzzbuzz

import "fmt"

const(
	Fizz = "fizz"
	Buzz = "buzz"
	FizzBuzzStr = "FizzBuzz"
)
func FizzBuzz(s int) string {
	if s%3 == 0 && s%5 == 0{
		return FizzBuzzStr
	}
	if s%3 == 0{
		return Fizz
	}
	if s%5 == 0{
		return Buzz
	}
	return fmt.Sprintf("%d",s)
}