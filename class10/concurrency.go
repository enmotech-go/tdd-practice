package class10

type WebsiteChecker func(string) bool

type Result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan *Result)

	for _, url := range urls {
		go func(url string) {
			resultChannel <- &Result{url, wc(url)}
		}(url)
	}

	//for i := 0; i < len(urls); i++ {
	//	result := <-resultChannel
	//	results[result.string] = result.bool
	//}

	for i := 0; i < len(urls); i++ {
		select {
		case result1 := <-resultChannel:
			results[result1.string] = result1.bool

		}
	}
	return results
}
