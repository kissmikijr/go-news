package news

import (
	"errors"
	"fmt"
)

var baseUrl = "http://newsapi.org/v2"

type newsApi struct {
	apiKey string
}

type HeadlinesParameters struct {
	country  string
	category string
	q        string
	sources  string
	pageSize int
	page     int
}
type Source struct {
	id   string
	name string
}
type Article struct {
	source      Source
	author      string
	title       string
	description string
	url         string
	urlToImage  string
	publishedAt string
	content     string
}
type HeadlinesResponse struct {
	status       string
	totalResults int
	articles     []Article
}

func NewApi(apiKey string) *newsApi {
	return &newsApi{apiKey}
}
func (c *newsApi) createUrl(hp *HeadlinesParameters, path string) (string, error) {
	url := baseUrl + fmt.Sprintf("%s?apiKey=%s", path, c.apiKey)
	if hp.country == "" && hp.q == "" && hp.category == "" {
		return "", errors.New("Required parameters are missing. Please set any of the following parameters and try again: sources, q, language, country, category.")
	}
	if hp.sources != "" && (hp.country != "" || hp.category != "") {
		return "", errors.New("You cant mix sources parameter with neither country nor category")
	}
	if hp.country != "" {
		url = url + fmt.Sprintf("&country=%s", hp.country)
	}
	if hp.category != "" {
		url = url + fmt.Sprintf("&category=%s", hp.category)
	}
	if hp.q != "" {
		url = url + fmt.Sprintf("&q=%s", hp.q)
	}

	return url, nil
}

// func (c *newsApi) topHeadlines(hp *HeadlinesParameters) (*HeadlinesResponse, error) {
// 	url, err := createUrl(hp, "/top-headlines")
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp, err := http.Get(url)
// 	return &HeadlinesResponse{}, nil
// }
func (c *newsApi) everything() {}
func (c *newsApi) sources()    {}
