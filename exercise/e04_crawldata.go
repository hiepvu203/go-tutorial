package main

import (
	"fmt"
	"sync"
)

// Fetcher interface
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// SafeMap để lưu URL đã crawl
type SafeMap struct {
	visited map[string]bool
	mu      sync.Mutex
}

func (s *SafeMap) Visit(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.visited[url] {
		return false // đã visit rồi
	}
	s.visited[url] = true
	return true
}

// Crawl hàm chính
func Crawl(url string, depth int, fetcher Fetcher, s *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// Kiểm tra đã visit chưa
	if !s.Visit(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, s, wg)
	}
}

func main() {
	s := &SafeMap{visited: make(map[string]bool)}
	var wg sync.WaitGroup
	var fetcher Fetcher
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, s, &wg)
	wg.Wait()
}
