package structinterface

import "math"

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Length) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
