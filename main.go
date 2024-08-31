package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	websiteURL := os.Args[1]
	parsedBaseURL, err := url.Parse(websiteURL)
	if err != nil {
		fmt.Printf("failed to parse provided url: %w", err)
		return
	}

	fmt.Printf("\n=========== CRAWLING '%v' STARTED ===========\n", websiteURL)
	config := config{
		pages: make(map[string]int, 0),
		baseURL: parsedBaseURL,
		mu: &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg: &sync.WaitGroup{},
	}

	config.wg.Add(1)
	go config.crawlPage(websiteURL)
	config.wg.Wait()

	fmt.Println("DONE")

	for k, v := range config.pages {
		fmt.Printf("%d - %s\n", v, k)
	}
}