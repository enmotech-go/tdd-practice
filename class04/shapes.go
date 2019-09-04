package class04

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (a Circle) Perimeter() float64 {
	return 2 * math.Pi * a.Radius
}

func (a Circle) Area() float64 {
	return math.Pi * a.Radius * a.Radius
}
