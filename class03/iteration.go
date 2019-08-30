package class03

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 4; i++ {
		repeated += character
	}
	return repeated
}
