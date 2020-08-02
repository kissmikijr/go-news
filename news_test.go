package news

import (
	"fmt"
	"testing"
)

func TestCreatUrlReturnsError(t *testing.T) {
	headlinesParams := &HeadlinesParameters{}
	path := "/test-path"
	apiKey := "test-api-key"
	_, err := createUrl(headlinesParams, path, apiKey)

	if err == nil {
		t.Error()
	}
}

var testPath = "/test-path"
var apiKey = "test-api-key"
var urltests = []struct {
	in  *HeadlinesParameters
	out string
}{
	{&HeadlinesParameters{country: "test"}, fmt.Sprintf("%s?apiKey=%s&country=test", testPath, apiKey)},
	{&HeadlinesParameters{q: "test"}, fmt.Sprintf("%s?apiKey=%s&q=test", testPath, apiKey)},
	{&HeadlinesParameters{category: "test"}, fmt.Sprintf("%s?apiKey=%s&category=test", testPath, apiKey)},
	{&HeadlinesParameters{language: "test"}, fmt.Sprintf("%s?apiKey=%s&language=test", testPath, apiKey)},
	{&HeadlinesParameters{country: "test", q: "query-string"}, fmt.Sprintf("%s?apiKey=%s&lcountry=test&q=query-string", testPath, apiKey)},
}

func TestCreateUrl(t *testing.T) {
	for _, tt := range urltests {
		t.Run("testname", func(t *testing.T) {
			u, _ := createUrl(tt.in, testPath, apiKey)
			expected := fmt.Sprintf("http://newsapi.org/v2%s", tt.out)
			if u != expected {
				t.Errorf("got %q, want %q", u, expected)
			}
		})
	}
}
