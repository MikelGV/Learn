package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
    Fetch(url string) (body string, urls []string, err error)
}

type SafeCounter struct {
    v map[string]bool
    mu sync.Mutex
    wg sync.WaitGroup
}

var c SafeCounter = SafeCounter{v: make(map[string]bool)}

func (s SafeCounter) checkvisited(url string) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    _, ok := s.v[url]
    if ok == false {
        s.v[url] = true
        return false
    }
    return true
}

func Crawl(url string, depth int, fetcher Fetcher) {
    // TODO: Fetch Urls in parallel.
    // TODO: Don't fetch the same URL twice.
    // This implementation doesn't do either:
    defer c.wg.Done()
    if depth <= 0 {
        return
    }

    if c.checkvisited(url) {
        return
    }

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Found: %s %q\n", url, body)
    for _, u := range urls {
        c.wg.Add(1)
        go Crawl(u, depth-1, fetcher)
    }
    return

}

func main() {
    c.wg.Add(1)
    Crawl("https://golang.org/", 4, fetcher)
    c.wg.Wait()
}

// FakeFetcher is Fetcher that returns canned results.
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
