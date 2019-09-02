package array_slice

func Sum(numbers []int)int{
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}


func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}