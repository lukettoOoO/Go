package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlCache struct {
	mu sync.Mutex
	wg sync.WaitGroup
	urls map[string]int
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *UrlCache) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	cache.mu.Lock()
	if cache.urls[url] > 0 {
		cache.wg.Done()
		cache.mu.Unlock()
		return
	}
	cache.urls[url]++
	cache.mu.Unlock()
	body, urls, err := fetcher.Fetch(url)
	
	if err != nil {
		fmt.Println(err)
		cache.wg.Done()
		return
	}
	
	fmt.Printf("found: %s %q\n", url, body)
	cache.mu.Lock()
	for _, u := range urls {
		if cache.urls[u] < 1 {
			cache.wg.Add(1)
			go Crawl(u, depth-1, fetcher, cache)
		}
	}
	cache.mu.Unlock()
	cache.wg.Done()
	return
}

func main() {
	cache := UrlCache{urls: make(map[string]int)}
	cache.wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &cache)
	cache.wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
