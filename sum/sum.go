package sum


//Sum changeLog : change the param type []int to slice int
func Sum(numbers []int) (sum int) {
	for _,v := range numbers{
		sum += v
	}
	return
}

//SumAll operation the two slice to one addition slice
//changeLog : change the sum opreation to append
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

//SumAllTails get the slice addition slice without the prefix dom
func SumAllTails(numbersToSum ...[]int) (sums []int){
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}