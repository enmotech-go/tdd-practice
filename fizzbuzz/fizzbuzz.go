package fizzbuzz

import "fmt"

const Fizz = "Fizz"
const Buzz = "Buzz"
const FizzBuzz = Fizz + Buzz

//GetSayText get say text
func GetSayText(input int) (outPut string) {
	canBeDivisibleByThree := false
	canBeDivisibleByFive := false
	if canBeDivisibleBySet(3, input) {
		canBeDivisibleByThree = true
	}

	if canBeDivisibleBySet(5, input) {
		canBeDivisibleByFive = true
	}

	if canBeDivisibleByFive && canBeDivisibleByThree {
		outPut = FizzBuzz
		return
	}

	if canBeDivisibleByThree {
		outPut = Fizz
		return
	}

	if canBeDivisibleByFive {
		outPut = Buzz
		return
	}
	return fmt.Sprintf("%d", input)
}

//canBeDivisibleBySet can be divisible by set
func canBeDivisibleBySet(set, input int) bool {
	if set == 0 {
		return false
	}
	remainder := input / set
	if remainder*set == input {
		return true
	}
	return false
}
