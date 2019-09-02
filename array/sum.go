package array

func Sum(numbers [5]int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}
