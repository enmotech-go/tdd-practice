package concurrency

type WebsiteChecker func(string) bool

type Result struct {
	Url       string
	Reachable bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChan := make(chan Result)
	for _, url := range urls {
		go func(u string) {
			resultsChan <- Result{Url: u, Reachable: wc(u)}
		}(url)
	}

	for i := len(urls); i > 0; i-- {
		result := <-resultsChan
		results[result.Url] = result.Reachable
	}
	return results
}
