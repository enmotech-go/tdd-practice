package class5

import "math"

type Rectangle struct {
	W float64
	H float64
}

type Circle struct {
	Radius float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.W + r.H)
}

//noinspection GoUnusedCallResult
func Area(r Rectangle) float64 {
	return r.W * r.H
}
func (r *Rectangle) Area() float64 {
	return r.W * r.H
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
