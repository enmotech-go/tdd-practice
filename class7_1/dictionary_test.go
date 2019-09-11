package class7_1

import "testing"

func TestSearch(t *testing.T) {
	dictionory := map[string]string{"test": "this is just a test"}
	got := Search(dictionory, "test")
	want := "this is just a test"
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}
}
