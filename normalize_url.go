package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (normalizedURL string, err error) {
	parsedUrl, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("could not parse url: %w", err)
	}

	query := parsedUrl.RawQuery
	if query != "" {
		query = "?" + query
	}

	normalizedURL = parsedUrl.Hostname() + parsedUrl.Path + query
	normalizedURL = strings.TrimPrefix(normalizedURL, "www.")
	normalizedURL = strings.TrimSuffix(normalizedURL, "/")
	normalizedURL = strings.ToLower(normalizedURL)

	return
}