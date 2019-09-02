package sum

func Sum(number [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += number[i]
	}
	return
}
