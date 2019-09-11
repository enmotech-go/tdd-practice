package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("test known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})

	t.Run("test unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "not found"

		if err == nil {
			t.Fatal("excepted an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
