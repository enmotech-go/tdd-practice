package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, want, got string) {
		if got != want {
			t.Errorf("want '%s' got '%s' ", want, got)
		}
	}

	t.Run("Say hello world as default testing case", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(t, want, got)
	})

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, want, got)
	})
}
