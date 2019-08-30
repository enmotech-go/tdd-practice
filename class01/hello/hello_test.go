package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello world as default testing case", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("want '%s' got '%s' ", want, got)
		}
	})

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("want '%s' got '%s' ", want, got)
		}
	})
}
