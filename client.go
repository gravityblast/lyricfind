package lyricfind

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const (
	TEST_BASE_URI = "http://test.lyricfind.com/api_service"
	LIVE_BASE_URI = "http://api.lyricfind.com"
)

type location interface {
	keyValue() (string, string)
}

type IPAddress string

func (i IPAddress) keyValue() (string, string) {
	return "ipaddress", string(i)
}

type Territory string

func (t Territory) keyValue() (string, string) {
	return "territory", string(t)
}

type apiHTTPError struct {
	statusCode int
	status     string
}

func (e *apiHTTPError) Error() string {
	return fmt.Sprintf("http error %d - %s", e.statusCode, e.status)
}

type apiResponseError struct {
	code        int
	description string
}

func (e *apiResponseError) Error() string {
	return fmt.Sprintf("api response error %d - %s", e.code, e.description)
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Environment struct {
	SearchApiKey  string
	DisplayApiKey string
	Live          bool
}

func (e *Environment) Url() *url.URL {
	var uri string
	if e.Live {
		uri = LIVE_BASE_URI
	} else {
		uri = TEST_BASE_URI
	}

	u, _ := url.Parse(uri)

	return u
}

type Client struct {
	searchApiKey  string
	displayApiKey string
	httpClient    httpClient
	baseURL       *url.URL
}

func New(env *Environment) *Client {
	c := &Client{
		searchApiKey:  env.SearchApiKey,
		displayApiKey: env.DisplayApiKey,
		httpClient:    &http.Client{},
		baseURL:       env.Url(),
	}

	return c
}

func (c *Client) newRequest(uripath string, values url.Values) (*http.Request, error) {
	values.Set("output", "json")
	u := *c.baseURL
	u.Path = path.Join(u.Path, uripath)
	u.RawQuery = values.Encode()

	return http.NewRequest("GET", u.String(), nil)
}

func (c *Client) newSearchRequest(loc location, options *SearchOptions) (*http.Request, error) {
	v := options.values()
	v.Set("apikey", c.searchApiKey)
	v.Set("searchtype", "track")
	v.Set("reqtype", "default")
	v.Set("displaykey", c.displayApiKey)

	key, value := loc.keyValue()
	v.Set(key, value)

	return c.newRequest("search.do", v)
}

func (c *Client) newLyricRequest(loc location, useragent string, amg int) (*http.Request, error) {
	v := url.Values{}
	v.Set("useragent", useragent)
	v.Set("apikey", c.displayApiKey)
	v.Set("reqtype", "default")
	v.Set("trackid", fmt.Sprintf("amg:%v", amg))

	key, value := loc.keyValue()
	v.Set(key, value)

	return c.newRequest("lyric.do", v)
}

func (c *Client) doRequest(r *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, &apiHTTPError{statusCode: resp.StatusCode, status: resp.Status}
	}

	return resp, nil
}

func (c *Client) Search(loc location, options *SearchOptions) (*SearchResponse, error) {
	req, err := c.newSearchRequest(loc, options)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sr SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		return nil, err
	}

	if sr.Response == nil {
		return nil, errors.New("response json doesn't have a `response` key")
	}

	if sr.Response.Code < 100 || sr.Response.Code > 199 {
		return nil, &apiResponseError{code: sr.Response.Code, description: sr.Response.Description}
	}

	return &sr, nil
}

func (c *Client) Lyric(loc location, useragent string, amg int) (*LyricsResponse, error) {
	req, err := c.newLyricRequest(loc, useragent, amg)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lr LyricsResponse
	err = json.NewDecoder(resp.Body).Decode(&lr)
	if err != nil {
		return nil, err
	}

	if lr.Response == nil {
		return nil, errors.New("response json doesn't have a `response` key")
	}

	if lr.Response.Code < 100 || lr.Response.Code > 199 {
		return nil, &apiResponseError{code: lr.Response.Code, description: lr.Response.Description}
	}

	return &lr, nil

}
