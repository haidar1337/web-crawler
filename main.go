package main

import (
	"fmt"
	"os"
)

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
	fmt.Printf("\n=========== CRAWLING '%v' STARTED ===========\n", websiteURL)
	html, err := getHTML(websiteURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(html)

}