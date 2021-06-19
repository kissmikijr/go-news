package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var baseUrl = "http://newsapi.org/v2"

type NewsApi struct {
	apiKey string
}

type HeadlinesParameters struct {
	Country  string
	Category string
	Q        string
	Sources  string
	PageSize int
	Page     int
}
type ArticleSource struct {
	Id   string
	Name string
}

type Article struct {
	Source      ArticleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	UrlToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
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

func (nae *NewsApiError) Error() string {
	return fmt.Sprintf("received code: %s message: %s", nae.Code, nae.Message)
}

func NewApi(apiKey string) *NewsApi {
	return &NewsApi{apiKey}
}
func (c *NewsApi) createHeadlinesUrl(hp *HeadlinesParameters) (string, error) {
	url := baseUrl + fmt.Sprintf("/top-headlines?apiKey=%s", c.apiKey)
	if hp.Country == "" && hp.Q == "" && hp.Category == "" {
		return "", errors.New("required parameters are missing. Please set any of the following parameters and try again: sources, q, language, country, category")
	}
	if hp.Sources != "" && (hp.Country != "" || hp.Category != "") {
		return "", errors.New("you cant mix sources parameter with neither country nor category")
	}
	if hp.Country != "" {
		url = url + fmt.Sprintf("&country=%s", hp.Country)
	}
	if hp.Category != "" {
		url = url + fmt.Sprintf("&category=%s", hp.Category)
	}
	if hp.Q != "" {
		url = url + fmt.Sprintf("&q=%s", hp.Q)
	}

	return url, nil
}

func (c *NewsApi) createEverythingUrl(ep *EverythingParameters) (string, error) {
	url := baseUrl + fmt.Sprintf("/everything?apiKey=%s", c.apiKey)

	if ep.Q != "" && ep.QinTitle != "" && ep.Sources != "" && ep.Domains != "" {
		return "", errors.New("please set any of the following required parameters and try again: q, qInTitle, sources, domains")
	}

	if ep.Q != "" {
		url = url + fmt.Sprintf("&q=%s", ep.Q)
	}
	if ep.QinTitle != "" {
		url = url + fmt.Sprintf("&qInTitle=%s", ep.QinTitle)
	}
	if ep.Sources != "" {
		url = url + fmt.Sprintf("&sources=%s", ep.Sources)
	}
	if ep.Domains != "" {
		url = url + fmt.Sprintf("&domains=%s", ep.Domains)
	}
	if ep.ExcludeDomains != "" {
		url = url + fmt.Sprintf("&excludeDomains=%s", ep.ExcludeDomains)
	}
	if ep.From != "" {
		url = url + fmt.Sprintf("&from=%s", ep.From)
	}
	if ep.To != "" {
		url = url + fmt.Sprintf("&to=%s", ep.To)
	}
	if ep.Language != "" {
		url = url + fmt.Sprintf("&language=%s", ep.Language)
	}
	if ep.SortBy != "" {
		url = url + fmt.Sprintf("&sortBy=%s", ep.SortBy)
	}
	if ep.PageSize != 0 {
		url = url + fmt.Sprintf("&pageSize=%d", ep.PageSize)
	}
	if ep.Page != 0 {
		url = url + fmt.Sprintf("&page=%d", ep.Page)
	}

	return url, nil
}
func createApiError(body io.Reader) error {
	var nae *NewsApiError
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &nae); err != nil {
		return err
	}
	return nae
}

func (c *NewsApi) TopHeadlines(hp *HeadlinesParameters) (*NewsResponse, error) {
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
		return nil, createApiError(resp.Body)
	}

	var newsResponse *NewsResponse
	err = json.NewDecoder(resp.Body).Decode(&newsResponse)
	if err != nil {
		return nil, err
	}
	return newsResponse, nil
}

type EverythingParameters struct {
	Q              string
	QinTitle       string
	Sources        string
	Domains        string
	ExcludeDomains string
	From           string
	To             string
	Language       string
	SortBy         string
	PageSize       int
	Page           int
}
type EverythingResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func (c *NewsApi) Everything(ep *EverythingParameters) (*EverythingResponse, error) {
	url, err := c.createEverythingUrl(ep)

	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, createApiError(resp.Body)
	}

	var newsApiResponse *EverythingResponse
	err = json.NewDecoder(resp.Body).Decode(&newsApiResponse)
	if err != nil {
		return nil, err
	}
	return newsApiResponse, nil
}

type category int

func (c category) String() string {
	return []string{"business", "entertainment", "general", "health", "science", "sports", "technology"}[c]
}

type language int

func (l language) String() string {
	return []string{"ar", "de", "en", "es", "er", "fr", "he", "it", "nl", "no", "pt", "ru", "se", "ud", "zh"}[l]
}

type country int

func (c country) String() string {
	return []string{"ae", "ar", "at", "au", "be", "bg", "br", "ca", "ch", "cn", "co", "cz", "de", "eg", "fr", "gb", "gr", "hk", "hu", "id", "ie", "il", "in", "it", "jp", "kr", "lt", "lv", "ma", "mx", "my", "ng", "nl", "no", "nz", "ph", "pl", "pt", "ro", "rs", "ru", "sa", "se", "sg", "si", "sk", "th", "tr", "tw", "ua", "us", "ve", "za"}[c]
}

type SourcesParameters struct {
	Category category
	Language language
	Country  country
}

const (
	CATEGORY_BUSINESS category = iota
	CATEGORY_ENTERTAINMENT
	CATEGORY_GENERAL
	CATEGORY_HEALTH
	CATEGORY_SCIENCE
	CATEGORY_SPORTS
	CATEGORY_TECHNOLOGY
)

const (
	LANG_AR language = iota
	LANG_DE
	LANG_EN
	LANG_ES
	LANG_FR
	LANG_HE
	LANG_IT
	LANG_NL
	LANG_NO
	LANG_PT
	LANG_RU
	LANG_SE
	LANG_UD
	LANG_ZH
)

const (
	COUNTRY_AE country = iota
	COUNTRY_AR
	COUNTRY_AT
	COUNTRY_AU
	COUNTRY_BE
	COUNTRY_BG
	COUNTRY_BR
	COUNTRY_CA
	COUNTRY_CH
	COUNTRY_CN
	COUNTRY_CO
	COUNTRY_CZ
	COUNTRY_DE
	COUNTRY_EG
	COUNTRY_FR
	COUNTRY_GB
	COUNTRY_GR
	COUNTRY_HK
	COUNTRY_HU
	COUNTRY_ID
	COUNTRY_IE
	COUNTRY_IL
	COUNTRY_IN
	COUNTRY_IT
	COUNTRY_JP
	COUNTRY_KR
	COUNTRY_LT
	COUNTRY_LV
	COUNTRY_MA
	COUNTRY_MX
	COUNTRY_MY
	COUNTRY_NG
	COUNTRY_NL
	COUNTRY_NO
	COUNTRY_NZ
	COUNTRY_PH
	COUNTRY_PL
	COUNTRY_PT
	COUNTRY_RO
	COUNTRY_RS
	COUNTRY_RU
	COUNTRY_SA
	COUNTRY_SE
	COUNTRY_SG
	COUNTRY_SI
	COUNTRY_SK
	COUNTRY_TH
	COUNTRY_TR
	COUNTRY_TW
	COUNTRY_UA
	COUNTRY_US
	COUNTRY_VE
	COUNTRY_ZA
)

func (c *NewsApi) createSourcesUrl(sp *SourcesParameters) (string, error) {
	url := baseUrl + fmt.Sprintf("/sources?apiKey=%s", c.apiKey)

	if sp.Category.String() != "" {
		url = url + fmt.Sprintf("&category=%s", sp.Category)
	}

	if sp.Language.String() != "" {
		url = url + fmt.Sprintf("&language=%s", sp.Language)
	}

	if sp.Country.String() != "" {
		url = url + fmt.Sprintf("&country=%s", sp.Country)
	}

	return url, nil
}

type Source struct {
	Id          string
	Name        string
	Description string
	Url         string
	Category    string
	Language    string
	Country     string
}
type SourcesResponse struct {
	Status  string
	Sources []Source
}

func (c *NewsApi) Sources(sp *SourcesParameters) (*SourcesResponse, error) {
	url, err := c.createSourcesUrl(sp)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, createApiError(resp.Body)
	}

	var sourcesResponse *SourcesResponse
	err = json.NewDecoder(resp.Body).Decode(&sourcesResponse)
	if err != nil {
		return nil, err
	}
	return sourcesResponse, nil
}
