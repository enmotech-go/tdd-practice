package mapp

import "testing"

func TestSearch(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just test"
		err := dictionary.Add(word, definition)
		asserError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("exist key", func(t *testing.T) {
		word := "test"
		definition := "this is just test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "xxxx")
		asserError(t, err, ExistError)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, d Dictionary, key, want string) {
	t.Helper()
	got, err := d.Search(key)
	if err != nil {
		t.Fatal("show find add word", err)
	}
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func asserError(t *testing.T, got, want error) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
