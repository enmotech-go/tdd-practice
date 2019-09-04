package structinterface

type Rectangle struct {
	width  float64
	length float64
}

func Perimeter(r Rectangle) float64 {
	return (r.width + r.length) * 2
}

func Area(r Rectangle) float64 {
	return r.width * r.length
}
