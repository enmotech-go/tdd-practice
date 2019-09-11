package mapp

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just test"}
	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just test"
		assertStrings(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatalf("expected to get a error")
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
