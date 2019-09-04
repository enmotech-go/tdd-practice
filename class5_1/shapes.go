package class5_1

type Rectangle struct {
	W float64
	H float64
}

func Perimeter(w, h float64) float64 {
	return 2 * (w + h)
}

func Area(r Rectangle) float64 {
	return r.H * r.W
}
