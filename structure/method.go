package structure

import "math"

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * 2
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
