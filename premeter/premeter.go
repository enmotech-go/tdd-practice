package premeter

type Rectangle struct {
	width  float64
	height float64
}

func Permeter(rectangle Rectangle) float64 {
	return (rectangle.width + rectangle.height) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
