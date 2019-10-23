package class11

import (
	"strconv"
	"strings"
)

const (
	three    = 3
	five     = 5
	threeStr = "3"
	fiveStr  = "5"
	fizz     = "fizz"
	buzz     = "buzz"
)

func Fizzbuzz(num int) string {
	if num == 0 {
		return strconv.Itoa(num)
	}
	if (mod(num, three) && mod(num, five)) || (containNum(num, threeStr) && containNum(num, fiveStr)) {
		return fizz + buzz
	}
	if (mod(num, three) && containNum(num, fiveStr)) || (containNum(num, threeStr) && mod(num, five)) {
		return fizz + buzz
	}
	if mod(num, three) || containNum(num, threeStr) {
		return fizz
	}
	if mod(num, five) || containNum(num, fiveStr) {
		return buzz
	}
	return strconv.Itoa(num)
}

func mod(num, mod int) bool {
	return num%mod == 0
}

func containNum(num int, str string) bool {
	return strings.Contains(strconv.Itoa(num), str)
}
