package main

func Sum(numbers [5]int) (ret int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
