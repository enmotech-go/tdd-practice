package class03

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaa"

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}

/*
jamesfans-Mac:class03 james$ go test -bench=.
goos: darwin
goarch: amd64
pkg: enmotech-go/tdd-practice/class03
BenchmarkRepeat-8       10000000               125 ns/op
PASS
ok      enmotech-go/tdd-practice/class03        1.389s
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
