package struct_tdd

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

//Perimeter 计算周长函数
func Perimeter(rectangle Rectangle)(perimeter float64){
	perimeter = 2*(rectangle.Width + rectangle.Height)
	return
}

func (r Rectangle) Area() (area float64) {
	area = r.Width * r.Height
	return
}

func (t Triangle) Area() (area float64) {
	area = t.Base * t.Height / 2
	return
}


func (c Circle) Area() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}