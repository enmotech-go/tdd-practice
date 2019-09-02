package sum

func Sum(nums [5]int) int {
	var sum int
	for _, i := range nums {
		sum += i
	}
	return sum
}
