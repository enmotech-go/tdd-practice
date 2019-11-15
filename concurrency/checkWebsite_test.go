package concurrency

import (
	"testing"
	"reflect"
)

func BenchmarkCheckWebsites(b *testing.B)  {
	urls := make([]string, 100)
	count := 100
	for i := 0; i < count; i++ {
		urls[i] = "a url"
	}

	// urls := []string{
	// 	"a.url",
	// 	"b.url",
	// 	"c.url",
	// 	"d.url",
	// }
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}
