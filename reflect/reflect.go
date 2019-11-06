package reflect

func walk(x interface{}, fn func(input string)) {

	fn("11")
}
