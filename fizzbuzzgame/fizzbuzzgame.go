package fizzbuzzgame

import (
	"strings"
	"strconv"
)

func FizzBuzzGame(num int)  string {
	if idRelatedNum(num, 3) && idRelatedNum(num , 5)  {
		return "fizzbuzz"
	}
	if idRelatedNum(num, 3)  {
		return "fizz"
	}
	if idRelatedNum(num , 5)  {
		return "buzz"
	}
	return strconv.Itoa(num)
}

func idRelatedNum(raw , needNum int) bool {
	return raw % needNum == 0 || strings.Contains(strconv.Itoa(raw), strconv.Itoa(needNum))
}