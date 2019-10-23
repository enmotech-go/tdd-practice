package fizzbuzz

import (
	"fmt"
	"strconv"
	"strings"
)

type FizzBuzz struct {
	input int
}

func (fb FizzBuzz) Number(num int) string {
	if idRelatedNum(num, 3) && idRelatedNum(num , 5)  {
		return "FizzBuzz"
	}
	if idRelatedNum(num, 3)  {
		return "Fizz"
	}
	if idRelatedNum(num , 5)  {
		return "Buzz"
	}
	return fmt.Sprintf("%d", fb.input)
}
func idRelatedNum(raw , needNum int) bool {
	return raw % needNum == 0 || strings.Contains(strconv.Itoa(raw), strconv.Itoa(needNum))
}