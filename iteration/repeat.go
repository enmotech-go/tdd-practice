package iteration

func Repeat(word string) string {
	var repeat string
	for i := 0; i < 5; i++ {
		repeat += word
	}
	return repeat
}
