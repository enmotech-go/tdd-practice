package sum


//changeLog : change the param type []int to slice int
func Sum(numbers []int) (sum int) {
	for _,v := range numbers{
		sum += v
	}
	return
}

//operation the two slice to one addition slice
func SumAll(numbersToSum ...[]int) (sums []int) {
	sums = []int{}
	return
}