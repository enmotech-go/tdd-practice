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
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	word := "test"
	def := "this is a test"

	t.Run("test add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(word, def)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, def)
	})

	t.Run("test add exist word", func(t *testing.T) {
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrExist)
		assertDefinition(t, dictionary, word, def)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	def := "this is a test"

	t.Run("test exist word", func(t *testing.T) {
		dictionary := Dictionary{word: def}
		newDef := "updated test"
		err := dictionary.Update(word, newDef)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDef)
	})

	t.Run("test non-exist word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update(word, def)

		assertError(t, err, ErrNotExist)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertDefinition(t *testing.T, dict Dictionary, word, def string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, def)
}
