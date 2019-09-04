package struct_tdd

type Rectangle struct {
	Width float64
	Height float64
}

//Perimeter 计算周长函数
func Perimeter(rectangle Rectangle)(perimeter float64){
	perimeter = 2*(rectangle.Width + rectangle.Height)
	return
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}