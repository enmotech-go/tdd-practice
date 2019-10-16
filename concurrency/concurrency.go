package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(url string) {
			results[url] = wc(url)
		}(url)
	}

	time.Sleep(5 * time.Second)
	return results
}
