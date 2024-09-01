package scraper

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScrapeFeed(t *testing.T) {
	tests := map[string]struct {
		url      string
		expected RssFeed
	}{
		"basic": {
			url:      "https://blog.boot.dev/index.xml",
			expected: RssFeed{},
		},
	}

	for name, tc := range tests {
		got, err := ScrapeFeed(tc.url)
		fmt.Println(got)
		if err != nil {
			t.Fatalf("Test %v failed\nGot error: %v", name, err)
		}
		if !reflect.DeepEqual(tc.expected, got) {
			t.Fatalf("Test %v failed\nExpected: %v\nGot: %v\n", name, tc.expected, got)
		}
	}
}
