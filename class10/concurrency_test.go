package class10

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	checkWebsites := CheckWebsites(mockWebsiteChecker, websites)
	want := len(websites)
	got := len(checkWebsites)
	if got != want {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
	expectedResults := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}
	if !reflect.DeepEqual(checkWebsites, expectedResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, checkWebsites)
	}
}
