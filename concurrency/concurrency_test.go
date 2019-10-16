package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://not-found.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://baidu.com",
		"http://blog.kliy.es",
		"http://not-found.com",
	}

	gotResults := CheckWebsites(mockWebsiteChecker, websites)
	got := len(gotResults)
	want := len(websites)
	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}

	wantResults := map[string]bool{
		"http://baidu.com":     true,
		"http://blog.kliy.es":  true,
		"http://not-found.com": false,
	}
	if !reflect.DeepEqual(gotResults, wantResults) {
		t.Fatalf("want %v, got %v", wantResults, gotResults)
	}
}
