package main

import (
	"fmt"
	"sort"
)

type page struct {
	page string
	count int		
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=======================")
	fmt.Printf("REPORT FOR %s\n", baseURL)
	fmt.Println("=======================")

	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		fmt.Printf("Found %v internal links to %s\n", page.count, page.page)
	}
}


func sortPages(pages map[string]int) []page {
	sorted := make([]page, 0)
	for k, v := range pages {
		sorted = append(sorted, page{page: k, count: v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		a := sorted[i].count
		b := sorted[j].count
		return a > b
	})

	return sorted
}

