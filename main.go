package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	maxPages		   int
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
	} else if len(args) < 2 {
		fmt.Println("no max concurrency provided")
		os.Exit(1)
	} else if len(args) < 3 {
		fmt.Println("no max pages provided")
		os.Exit(1)
	} else if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	websiteURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("max concurrency must be an integer")
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("max pages must be an integer")
	}
	parsedBaseURL, err := url.Parse(websiteURL)
	if err != nil {
		fmt.Printf("failed to parse provided url: %w", err)
		return
	}

	fmt.Printf("\n=========== CRAWLING '%v' STARTED ===========\n", websiteURL)
	config := config{
		maxPages: maxPages,
		pages: make(map[string]int, 0),
		baseURL: parsedBaseURL,
		mu: &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg: &sync.WaitGroup{},
	}

	config.wg.Add(1)
	go config.crawlPage(websiteURL)
	config.wg.Wait()

	printReport(config.pages, websiteURL)
}