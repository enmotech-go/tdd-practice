package class7_1

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"
	assertString(got, want, t)
}

func assertString(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}
}
