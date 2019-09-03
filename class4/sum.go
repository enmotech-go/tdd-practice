package class4

func Sum(num []int) int {
	var sum int
	for _, v := range num {
		sum += v
	}
	return sum
}

func SumAll(nums ...[]int) (sums []int) {
	for _, v := range nums {
		sums = append(sums, Sum(v))
	}
	return
}

func SumTails(nums ...[]int) (sums []int) {
	for _, v := range nums {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(v[1:]))
		}
	}
	return
}
