package struct_method_interface

func Perimeter(rect Rectangle) float64 {
	return (rect.Width + rect.Height) * 2
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}
