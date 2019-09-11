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
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")
	want := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Errorf("should Add err %s", err)
	}

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
