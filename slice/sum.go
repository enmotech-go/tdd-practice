package main

func Sum(numbers [5]int) (ret int) {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
