package news

import (
	"fmt"
	"testing"
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
		{&HeadlinesParameters{Sources: "validsource1,validsource2", Category: "valid-category"}, "source used with category"},
		{&HeadlinesParameters{Sources: "validsource1,validsource2", Country: "valid-country"}, "source used with country"},
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
		{&HeadlinesParameters{Country: "test"}, fmt.Sprintf("/top-headlines?apiKey=%s&country=test", apiKey)},
		{&HeadlinesParameters{Q: "query-string"}, fmt.Sprintf("/top-headlines?apiKey=%s&q=query-string", apiKey)},
		{&HeadlinesParameters{Category: "test"}, fmt.Sprintf("/top-headlines?apiKey=%s&category=test", apiKey)},
		{&HeadlinesParameters{Country: "test", Q: "query-string"}, fmt.Sprintf("/top-headlines?apiKey=%s&country=test&q=query-string", apiKey)},
		{&HeadlinesParameters{Country: "test", Category: "test-category"}, fmt.Sprintf("/top-headlines?apiKey=%s&country=test&category=test-category", apiKey)},
		{&HeadlinesParameters{Country: "test", Q: "query-string", Category: "test-category"}, fmt.Sprintf("/top-headlines?apiKey=%s&country=test&category=test-category&q=query-string", apiKey)},
		{&HeadlinesParameters{Category: "test-category", Q: "query-string"}, fmt.Sprintf("/top-headlines?apiKey=%s&category=test-category&q=query-string", apiKey)},
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
		{&EverythingParameters{Q: "test"}, fmt.Sprintf("/everything?apiKey=%s&q=test", apiKey)},
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

func TestCreateSourcesUrl(t *testing.T) {
	client := NewApi("api-key")
	url, err := client.createSourcesUrl(&SourcesParameters{Category: CATEGORY_BUSINESS, Language: LANG_DE, Country: COUNTRY_AE})
	if err != nil {
		t.Error(err)
	}
	if url != "http://newsapi.org/v2/sources?apiKey=api-key&category=business&language=de&country=ae" {
		t.Errorf("got %s", url)
	}
}
func TestCreateSourcesUrlThrowsError(t *testing.T) {

}
