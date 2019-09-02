package sum

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}
