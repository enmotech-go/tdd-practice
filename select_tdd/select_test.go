package select_tdd

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://github.com/"
	fastURL := "http://www.163.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}



