package main

import (
	"errors"
	"fmt"
	"net/http"
)

var baseUrl = "http://newsapi.org/v2"

type newsApi struct {
	apiKey string
}

type HeadlinesParameters struct {
	country  string
	category string
	q        string
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
func createUrl(params *HeadlinesParameters, path string) (string, error) {
	if params.q != "" && (params.country != "" || params.category != "") {
		return "", errors.New("You cant mix sources parameter with neither country nor category")
	}
	if params.country != "" {
		baseUrl = baseUrl + fmt.Sprintf("country=%s", params.country)
	}
	if params.category != "" {
		baseUrl = baseUrl + fmt.Sprintf("category=%s", params.category)
	}
	if params.q != "" {
		baseUrl = baseUrl + fmt.Sprintf("q=%s", params.q)
	}

	return baseUrl, nil
}

func (c *newsApi) topHeadlines(params *HeadlinesParameters) (*HeadlinesResponse, error) {
	url, err := createUrl(params, "/top-headlines")
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
}
func (c *newsApi) everything() {}
func (c *newsApi) sources()    {}

func main() {}
