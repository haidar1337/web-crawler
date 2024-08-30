package main

import (
	"strings"
	"testing"
)


func TestNormalizeURL(t *testing.T) {
	tests := map[string]struct {
		inputURL      string
		expected      string
		errorContains string
	}{
			"remove scheme": {inputURL: "https://google.com", expected: "google.com",},
			"remove slashes": {inputURL: "youtube.com/", expected: "youtube.com"},
			"url with path": {inputURL: "https://en.wikipedia.org/wiki/Go_(programming_language)", expected: "en.wikipedia.org/wiki/go_(programming_language)"},
			"url with query": {inputURL: "https://google.com/search?q=dogs", expected: "google.com/search?q=dogs"},
			"url with www": {inputURL: "http://www.google.com", expected: "google.com"},
			"invalid url": {inputURL: ":\\invalidURL", expected: "", errorContains: "could not parse"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test - '%s' FAIL: unexpected error: %v", name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test - '%s' FAIL: unexpected error: %v", name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test - '%s' FAIL: expected error: %v, got none.", name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test - %s FAIL: expected URL: %v, actual: %v", name, tc.expected, actual)
			}
		})
	}
}