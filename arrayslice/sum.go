package sum

func Sum(nums []int) int {
	var sum int
	for _, i := range nums {
		sum += i
	}
	return sum
}

func sumAll(numList [][]int) []int {
	length := len(numList)
	rst := make([]int, length)
	for i := range numList {
		rst[i] = Sum(numList[i])
	}
	return rst
}

func sumAllTails(numberList [][]int) []int {
	var rst []int
	for _, num := range numberList {
		if len(num) == 0 {
			rst = append(rst, 0)
		} else {
			rst = append(rst, Sum(num[1:]))
		}
	}
	return rst
}
