package class5

type Rectangle struct {
	W float64
	H float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.W + r.H)
}

//noinspection GoUnusedCallResult
func Area(r Rectangle) float64 {
	return r.W * r.H
}
