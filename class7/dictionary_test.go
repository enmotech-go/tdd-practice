package class7

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertStrings(t, err.Error(), ErrNotFound.Error())

	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	value, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if value != definition {
		t.Errorf("definition %s value %s", definition, value)
	}
}
