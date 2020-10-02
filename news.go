package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	sources  string
	pageSize int
	page     int
}
type Source struct {
	id   string
	name string
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}
type HeadlinesResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}
type NewsApiError struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
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

func (c *newsApi) topHeadlines(hp *HeadlinesParameters) (*HeadlinesResponse, error) {
	url, err := c.createUrl(hp, "/top-headlines")
	fmt.Println(url)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		var t *NewsApiError
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &t); err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("Received status code: %d error: %s", resp.StatusCode, data))
	}

	var newsApiResponse *HeadlinesResponse
	err = json.NewDecoder(resp.Body).Decode(&newsApiResponse)
	if err != nil {
		return nil, err
	}
	return newsApiResponse, nil
}

func (c *newsApi) everything() {}
func (c *newsApi) sources()    {}
