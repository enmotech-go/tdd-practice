package fuzzbuzz

import (
	"fmt"
	"strings"
)

const(
	Fizz = "fizz"
	Buzz = "buzz"
	FizzBuzzStr = "FizzBuzz"
)
func FizzBuzz(s int) string {
	if isRelatedTo(s,3) && isRelatedTo(s,5){
		return FizzBuzzStr
	}
	if isRelatedTo(s,3){
		return Fizz
	}
	if isRelatedTo(s,5){
		return Buzz
	}
	return fmt.Sprintf("%d",s)
}


func isRelatedTo(num , relate int) bool{
	if num%relate == 0{
		return true
	}
	if strings.Contains(fmt.Sprintf("%d",num),fmt.Sprintf("%d",relate)){
		return true
	}
	return false
}