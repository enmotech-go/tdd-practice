package sum

func Sum(nums []int) int {
	var sum int
	for _, i := range nums {
		sum += i
	}
	return sum
}
