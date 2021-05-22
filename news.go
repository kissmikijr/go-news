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
type NewsResponse struct {
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
func (c *newsApi) createHeadlinesUrl(hp *HeadlinesParameters) (string, error) {
	url := baseUrl + fmt.Sprintf("/top-headlines?apiKey=%s", path, c.apiKey)
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

func (c *newsApi) createEverythingUrl(ep *EverythingParameters) (string, error) {
	url := baseUrl + fmt.Sprintf("%s?apiKey=%s", path, c.apiKey)

	if ep.q != "" && ep.qInTitle != "" && ep.sources != "" && ep.domains != "" {
		return "", errors.New("Please set any of the following required parameters and try again: q, qInTitle, sources, domains.")
	}

	if ep.q != "" {
		url = url + fmt.Sprintf("&q=%s", ep.q)
	}
	if ep.qInTitle != "" {
		url = url + fmt.Sprintf("&qInTitle=%s", ep.qInTitle)
	}
	if ep.sources != "" {
		url = url + fmt.Sprintf("&sources=%s", ep.sources)
	}
	if ep.domains != "" {
		url = url + fmt.Sprintf("&domains=%s", ep.domains)
	}
	if ep.excludeDomains != "" {
		url = url + fmt.Sprintf("&excludeDomains=%s", ep.excludeDomains)
	}
	if ep.from != "" {
		url = url + fmt.Sprintf("&from=%s", ep.from)
	}
	if ep.to != "" {
		url = url + fmt.Sprintf("&to=%s", ep.to)
	}
	if ep.language != "" {
		url = url + fmt.Sprintf("&language=%s", ep.language)
	}
	if ep.sortBy != "" {
		url = url + fmt.Sprintf("&sortBy=%s", ep.sortBy)
	}
	if ep.pageSize != 0 {
		url = url + fmt.Sprintf("&pageSize=%d", ep.pageSize)
	}
	if ep.page != 0 {
		url = url + fmt.Sprintf("&page=%d", ep.page)
	}

	return url, nil
}

func (c *newsApi) TopHeadlines(hp *HeadlinesParameters) (*NewsResponse, error) {
	url, err := c.createHeadlinesUrl(hp)

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

	var newsResponse *NewsResponse
	err = json.NewDecoder(resp.Body).Decode(&newsResponse)
	if err != nil {
		return nil, err
	}
	return newsResponse, nil
}

type EverythingParameters struct {
	q              string
	qInTitle       string
	sources        string
	domains        string
	excludeDomains string
	from           string
	to             string
	language       string
	sortBy         string
	pageSize       int
	page           int
}
type EverythingResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func (c *newsApi) Everything(ep *EverythingParameters) (*EverythingResponse, error) {
	url, err := c.createEverythingUrl(ep, "/everything")

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

	var newsApiResponse *EverythingResponse
	err = json.NewDecoder(resp.Body).Decode(&newsApiResponse)
	if err != nil {
		return nil, err
	}
	return newsApiResponse, nil
}
type SourcesParameters struct {
	category string
	language string
	country string
}
SUPPORTED_CATEGORIES := map[string]bool{
	"business": true,
	"entertainment": true,
	"general": true,
	"health": true,
	"science": true,
	"sports": true,
	"technology": true
} 

func CreateSourcesParameters(category string, language string, country string)  (&SourcesParameters, error){

	return &SourcesParameters{category, language, country}, nil
}

func (c *newsApi) createSourcesUrl(sp *SourcesParameters) (string, err) {
	url := baseUrl + fmt.Sprintf("/sources?apiKey=%s", c.apiKey)

	if sp.category != "" {
		url = url + fmt.Sprintf("&category=%s", sp.category)
	}

	if sp.language != "" {
		url = url + fmt.Sprintf("&language=%s", sp.language)
	}

	if sp.country != "" {
		url = url + fmt.Sprintf("&country=%s", sp.country)
	}

	return url, nil
}
func (c *newsApi) Sources(sp *SourcesParameters) {
	url, err := c.createSourcesUrl(sp)
}
