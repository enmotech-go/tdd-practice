package concurrency

import (
	"time"

)

/**
class_10
demand: 测试网站是否能访问，go 提高并发速度，channel 解决竞态条件
*/

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool  {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _,url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
			// results[u] = wc(u)
		}(url)
		// results[url] = wc(url)

	}
	
	for i := 0; i < len(urls); i++ {
		result := <- resultChannel
		results[result.string] = result.bool
	}
	return results
}

func slowStubWebsiteChecker(_ string) bool  {
	time.Sleep(time.Millisecond*20)
	return true
}