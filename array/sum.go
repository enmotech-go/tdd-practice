package array

func SumArray(numbers [5]int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumSlice(numbers []int) int {
	return 0
}
