package class5_1

import "math"

type Rectangle struct {
	W float64
	H float64
}

type Circle struct {
	R float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.W + r.H)
}

func (r Rectangle) Area() float64 {
	return r.H * r.W
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}
