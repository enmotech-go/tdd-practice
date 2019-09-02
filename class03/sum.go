package class03

func Sum(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}
