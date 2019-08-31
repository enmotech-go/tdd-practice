package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, want, got string) {
		t.Helper()
		if got != want {
			t.Errorf("want '%s' got '%s' ", want, got)
		}
	}

	t.Run("Say hello world as default testing case", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, want, got)
	})

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, want, got)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, want, got)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Rouphen", "Bonjour")
		want := "Bonjour, Rouphen"

		assertCorrectMessage(t, want, got)
	})
}
