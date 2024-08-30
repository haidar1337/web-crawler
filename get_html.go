package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	if res.StatusCode > 299 {
		return "", fmt.Errorf("response failed with status code: %v", res.StatusCode)
	}

	contentType := res.Header.Get("content-type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("reponse is not in text/html format")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	defer res.Body.Close()

	return string(body), nil
}