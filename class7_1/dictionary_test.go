package class7_1

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(got, want, t)
	})
	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(err, ErrNotFound, t)
	})

}

func assertString(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}
}

func assertError(got, want error, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func TestAdd(t *testing.T) {

	t.Run("new world ", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		_ = dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("exist word ", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(err, ErrWordExists, t)
		assertDefinition(t, dictionary, word, definition)
	})
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

func TestUpdate(t *testing.T) {

	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"
	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)
}
