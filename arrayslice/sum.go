package sum

func Sum(nums []int) int {
	var sum int
	for _, i := range nums {
		sum += i
	}
	return sum
}

func sumAll(numList [][]int) []int {
	var rst []int
	for _, num := range numList {
		rst = append(rst, Sum(num))
	}
	return rst
}
