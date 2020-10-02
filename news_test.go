package news

import (
	"fmt"
	"testing"
)

func TestCreatUrlReturnsError(t *testing.T) {
	apiKey := "test-api-key"
	newsApi := NewApi(apiKey)
	headlinesParams := &HeadlinesParameters{}
	path := "/test-path"
	_, err := newsApi.createUrl(headlinesParams, path)

	if err == nil {
		t.Error()
	}
}

func TestCreateUrlReturnsErrorWhenSourcesUsedWithWrongParams(t *testing.T) {
	apiKey := "valid-key"
	path := "/test-path"
	newsApi := NewApi(apiKey)
	urltests := []struct {
		in   *HeadlinesParameters
		name string
	}{
		{&HeadlinesParameters{sources: "validsource1,validsource2", category: "valid-category"}, "source used with category"},
		{&HeadlinesParameters{sources: "validsource1,validsource2", country: "valid-country"}, "source used with country"},
	}

	for _, tt := range urltests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newsApi.createUrl(tt.in, path)
			if err == nil {
				t.Error()
			}
		})
	}
}

func TestCreateUrl(t *testing.T) {
	testPath := "/test-path"
	apiKey := "test-api-key"
	urltests := []struct {
		in  *HeadlinesParameters
		out string
	}{
		{&HeadlinesParameters{country: "test"}, fmt.Sprintf("%s?apiKey=%s&country=test", testPath, apiKey)},
		{&HeadlinesParameters{q: "query-string"}, fmt.Sprintf("%s?apiKey=%s&q=query-string", testPath, apiKey)},
		{&HeadlinesParameters{category: "test"}, fmt.Sprintf("%s?apiKey=%s&category=test", testPath, apiKey)},
		{&HeadlinesParameters{country: "test", q: "query-string"}, fmt.Sprintf("%s?apiKey=%s&country=test&q=query-string", testPath, apiKey)},
		{&HeadlinesParameters{country: "test", category: "test-category"}, fmt.Sprintf("%s?apiKey=%s&country=test&category=test-category", testPath, apiKey)},
		{&HeadlinesParameters{country: "test", q: "query-string", category: "test-category"}, fmt.Sprintf("%s?apiKey=%s&country=test&category=test-category&q=query-string", testPath, apiKey)},
		{&HeadlinesParameters{category: "test-category", q: "query-string"}, fmt.Sprintf("%s?apiKey=%s&category=test-category&q=query-string", testPath, apiKey)},
	}

	for _, tt := range urltests {
		newsApi := NewApi(apiKey)
		t.Run("testname", func(t *testing.T) {
			u, err := newsApi.createUrl(tt.in, testPath)
			if err != nil {
				t.Error(err)
			}
			expected := fmt.Sprintf("http://newsapi.org/v2%s", tt.out)
			if u != expected {
				t.Errorf("got %q, want %q", u, expected)
			}
		})
	}
}
