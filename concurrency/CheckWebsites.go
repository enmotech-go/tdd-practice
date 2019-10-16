package concurrency

type WebsiteChecker func(string) bool

type Result struct {
	Url       string
	Reachable bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan Result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- Result{
				Url:       u,
				Reachable: wc(u),
			}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.Url] = result.Reachable
	}
	//time.Sleep(2 * time.Second)
	return results
}
