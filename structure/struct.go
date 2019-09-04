package structure

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter2(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

func Area2(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}

type Circle struct {
	Radius float64
}
