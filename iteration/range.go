package iteration

const count = 5

func Repeat(str string) string {
	repeat := ""
	for i := 0; i < count; i++ {
		repeat += str
	}
	return repeat
}
