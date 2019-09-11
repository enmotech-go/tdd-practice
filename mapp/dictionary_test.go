package mapp

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just test"}
	got := dictionary.Search("test")
	want := "this is just test"
	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
