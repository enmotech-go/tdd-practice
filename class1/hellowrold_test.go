package class1

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		if got != want {
			t.Errorf("got '%s' ,want '%s'", got, want)
		}
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		if got != want {
			t.Errorf("got '%s' ,want '%s'", got, want)
		}
	})
}
