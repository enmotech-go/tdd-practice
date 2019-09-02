package array

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sum = make([]int, len(numbers))
	for i, n := range numbers {
		sum[i] = Sum(n)
	}
	return sum
}

func SumAllTails(numbers ...[]int) []int {
	var sum = make([]int, len(numbers))
	for i, n := range numbers {
		if len(n) > 0 {
			sum[i] = Sum(n[1:])
		}
	}
	return sum
}
