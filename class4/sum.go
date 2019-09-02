package class4

func Sum(num []int) int {
	var sum int
	for _, v := range num {
		sum += v
	}
	return sum
}

func SumAll(num ...[]int) (sums []int) {
	for _, v := range num {
		sums = append(sums, Sum(v))
	}
	return
}
