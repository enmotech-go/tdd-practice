package class11

import (
	"strconv"
)

const (
	three = 3
	five  = 5
	fizz  = "fizz"
	buzz  = "buzz"
)

func Fizzbuzz(num int) string {
	if num == 0 {
		return strconv.Itoa(num)
	}
	if mod(num, three) && mod(num, five) {
		return fizz + buzz
	}
	if mod(num, three) {
		return fizz
	}
	if mod(num, five) {
		return buzz
	}
	return strconv.Itoa(num)
}

func mod(num, mod int) bool {
	return num%mod == 0
}
