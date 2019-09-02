package sum


//changeLog : change the param type []int to slice int
func Sum(numbers []int) (sum int) {
	for _,v := range numbers{
		sum += v
	}
	return
}

//operation the two slice to one addition slice
//changeLog : change the sum opreation to append
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}