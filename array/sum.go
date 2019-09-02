package array

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbers1 []int, numbers2 []int) []int {
	sum1 := Sum(numbers1)
	sum2 := Sum(numbers2)
	return []int{sum1, sum2}
}
