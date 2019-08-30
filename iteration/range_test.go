package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	want := "aaaaa"

	if repeated != want {
		t.Errorf("want %v, but get %v", want, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
