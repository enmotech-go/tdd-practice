package premeter

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Permeter() float64 {
	return (r.width + r.height) * 2
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
