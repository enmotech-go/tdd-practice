package class4

func Sum(num []int) int {
	var sum int
	for _, v := range num {
		sum += v
	}
	return sum
}

func SumAll(num ...[]int) (sums []int) {
	lengthOfNumbers := len(num)
	sums = make([]int, lengthOfNumbers)

	for k, v := range num {
		sums[k] = Sum(v)
	}
	return
}
