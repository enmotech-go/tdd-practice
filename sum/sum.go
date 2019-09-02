package sum


//changeLog : change the param type []int to slice int
func Sum(numbers []int) (sum int) {
	for _,v := range numbers{
		sum += v
	}
	return
}