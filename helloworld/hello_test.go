package helloworld

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf(got, want)
		}
	}

	got := hello("Chris", "English")
	want := "Hello Chris"

	assertCorrectMessage(t, got, want)

	t.Run("say hello to Spanish", func(t *testing.T) {
		got := hello("Chris", "Spanish")
		want := "Hola, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to French", func(t *testing.T) {
		got := hello("Chris", "French")
		want := "Bonjour, Chris"

		assertCorrectMessage(t, got, want)
	})
}
