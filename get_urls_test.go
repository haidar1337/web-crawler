package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetUrlsFromHTML(t *testing.T) {
	tests := map[string]struct{
		inputURL string
		inputBody string
		errorContains string
		expected []string
	}{
		"absolute and relative paths": {
			inputURL: "https://google.com",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>google.com</span>
		</a>
		<a href="https://other.com/path/one">
			<span>google.com</span>
		</a>
	</body>
</html>
`,
		expected: []string{"https://google.com/path/one", "https://other.com/path/one"},
		},
		"none": {
			inputURL: "https://google.com",
			inputBody: `
<html>
	<body>
		<p>Hello, World!</p>
	</body>
</html>
`,
			expected: nil,
		},
		"invalid base url": { 
			inputURL: ":\\google.com",
			inputBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev></span>
	</a>
</html body>
`,
			expected: nil,
			errorContains: "failed to parse base url",
		},
		"invalid href url": {
			inputURL: "https://google.com",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">invalid url</a>
	</body>
</html>
`,
		expected: nil,
		},
		"no href": {
			inputURL: "https://google.com",
			inputBody: `
<html>
	<body>
		<a>no ref</a>
	</body>
</html>
`,
			expected: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test - %s FAIL unexpected error: %v", name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test - %s FAIL unexpected error: %v", name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test - %s FAIL expected error to contain: %v, got none.", name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test - %s FAIL expected: %v, actual: %v", name, tc.expected, actual)
			}
		})
	}
}