package class10

import (
	"reflect"
	"testing"
	"time"
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

	t.Log(checkWebsites)
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

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(200 * time.Microsecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
