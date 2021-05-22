package news

import (
	"fmt"
	"testing"
	"log"
)

func TestCreatUrlReturnsError(t *testing.T) {
	apiKey := "test-api-key"
	newsApi := NewApi(apiKey)
	headlinesParams := &HeadlinesParameters{}
	_, err := newsApi.createHeadlinesUrl(headlinesParams)

	if err == nil {
		t.Error()
	}
}

func TestCreateUrlReturnsErrorWhenSourcesUsedWithWrongParams(t *testing.T) {
	apiKey := "valid-key"
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
			_, err := newsApi.createHeadlinesUrl(tt.in)
			if err == nil {
				t.Error()
			}
		})
	}
}

func TestCreateHeadlinesUrl(t *testing.T) {
	apiKey := "test-api-key"
	urltests := []struct {
		in  *HeadlinesParameters
		out string
	}{
		{&HeadlinesParameters{country: "test"}, fmt.Sprintf("/top-headlines?apiKey=%s&country=test", apiKey)},
		{&HeadlinesParameters{q: "query-string"}, fmt.Sprintf("/top-headlines?apiKey=%s&q=query-string", apiKey)},
		{&HeadlinesParameters{category: "test"}, fmt.Sprintf("/top-headlines?apiKey=%s&category=test", apiKey)},
		{&HeadlinesParameters{country: "test", q: "query-string"}, fmt.Sprintf("/top-headlines?apiKey=%s&country=test&q=query-string", apiKey)},
		{&HeadlinesParameters{country: "test", category: "test-category"}, fmt.Sprintf("/top-headlines?apiKey=%s&country=test&category=test-category",  apiKey)},
		{&HeadlinesParameters{country: "test", q: "query-string", category: "test-category"}, fmt.Sprintf("/top-headlines?apiKey=%s&country=test&category=test-category&q=query-string", apiKey)},
		{&HeadlinesParameters{category: "test-category", q: "query-string"}, fmt.Sprintf("/top-headlines?apiKey=%s&category=test-category&q=query-string", apiKey)},
	}

	for _, tt := range urltests {
		newsApi := NewApi(apiKey)
		t.Run("testname", func(t *testing.T) {
			u, err := newsApi.createHeadlinesUrl(tt.in)
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

func TestCreateEverythinUrl(t *testing.T) {
	apiKey := "test-api-key"
	urltests := []struct {
		in  *EverythingParameters
		out string
	}{
		{&EverythingParameters{q: "test"}, fmt.Sprintf("/everything?apiKey=%s&q=test", apiKey)},
	}

	for _, tt := range urltests {
		newsApi := NewApi(apiKey)
		t.Run("testname", func(t *testing.T) {
			u, err := newsApi.createEverythingUrl(tt.in)
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

func TestEverything(t *testing.T) {
	client := NewApi("e1750e47dd844d54ac301a7f99c8cdf5")
	resp, err := client.Everything(&EverythingParameters{})
	log.Println(resp, "@@@@@")
	if err != nil {
		t.Error(err)
	}
}