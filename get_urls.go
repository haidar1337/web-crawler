package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)


func getURLsFromHTML(htmlBody, rawURL string) ([]string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %w", err)
	}
	reader := strings.NewReader(htmlBody)
	parsedHTML, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var out []string
	getURLSFromHTMLNodeRecursive(*parsedURL, parsedHTML, &out)
	
	return out, nil
}

func getURLSFromHTMLNodeRecursive(baseURL url.URL,htmlNode *html.Node, anchorTags *[]string) {
	if htmlNode.Type == html.ElementNode && htmlNode.Data == "a" {
		for _, a := range htmlNode.Attr {
			if a.Key == "href" {
				href, err := url.Parse(a.Val)
				if err != nil {
					fmt.Println("failed to parse href url: %w", err)
					continue
				}

				*anchorTags = append(*anchorTags, baseURL.ResolveReference(href).String())
				break
			}
		}
	}

	for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
		getURLSFromHTMLNodeRecursive(baseURL, child, anchorTags)
	}

}