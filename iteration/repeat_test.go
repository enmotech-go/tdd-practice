package iteration

import (
	"testing"
)
func TestReapt(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"
	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
