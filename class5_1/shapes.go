package class5_1

type Rectangle struct {
	W float64
	H float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.W + r.H)
}

func Area(r Rectangle) float64 {
	return r.H * r.W
}
