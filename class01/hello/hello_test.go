package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("want '%s' got '%s' ", want, got)
	}
}
