package class4

func Sum(num [5]int) int {
	var sum int
	for i := 0; i < 5; i++ {
		sum += num[i]
	}
	return sum
}
