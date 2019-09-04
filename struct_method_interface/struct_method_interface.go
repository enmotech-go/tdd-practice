package struct_method_interface

import "math"

func Perimeter(rect Rectangle) float64 {
	return (rect.Width + rect.Height) * 2
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}
