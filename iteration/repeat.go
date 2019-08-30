package iteration

func Repeat(char string, times int) string {
	var result string
	for index := 0; index < times; index++ {
		result += char
	}
	return result
}
