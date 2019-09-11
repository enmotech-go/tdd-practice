package mapp

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just test"
	dictionary.Add(word, definition)
	assertDefinition(t, dictionary, word, definition)
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
