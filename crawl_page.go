package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	reached := cfg.checkPagesLimitReached()
	cfg.concurrencyControl<-struct{}{}
	defer func ()  {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	if reached {
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("failed to parse current url", err)	
		return
	}

	if cfg.baseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	isFirst := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirst {
		return
	}
	fmt.Printf("crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("failed to get the HTML body of the current url", err)
		return
	}

	urls, err := getURLsFromHTML(html, rawCurrentURL)
	if err != nil {
		fmt.Println("failed to get the urls of the current page", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if cfg.pages[normalizedURL] != 0 {
		cfg.pages[normalizedURL]++
		return
	}

	isFirst = true
	cfg.pages[normalizedURL] = 1

	return 
}

func (cfg *config) checkPagesLimitReached() (reached bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if len(cfg.pages) >= cfg.maxPages {
		reached = true
		return
	}

	return
}