package lyricfind

import (
  "fmt"
  "net/url"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

const BASE_URL = "http://test.lyricfind.com/api_service"

type client struct {
  SearchApiKey string
  DisplayApiKey string
  httpClient HttpClient
}

type ResponseError struct {
  Code int
  Description string
  Message string
}

type HttpClient interface {
  Get(string) (*http.Response, error)
}

func (c client) SearchUrl() string {
  return fmt.Sprintf("%s/%s", BASE_URL, "search.do")
}

func (c client) LyricsUrl() string {
  return fmt.Sprintf("%s/%s", BASE_URL, "lyric.do")
}

func (c client) MergeDefaultRequestParams(params url.Values) url.Values  {
  params.Set("output", "json")
  return params
}

func (c client) MergeSearchRequestParams(params url.Values) url.Values {
  params.Set("reqtype", "default")
  params.Set("searchtype", "track")
  params.Set("apikey", c.SearchApiKey)
  return c.MergeDefaultRequestParams(params)
}

func (c client) MergeLyricsRequestParams(params url.Values) url.Values {
  params.Set("reqtype", "default")
  params.Set("apikey", c.DisplayApiKey)
  return c.MergeDefaultRequestParams(params)
}

func (c client) BuildSearchUrl(params url.Values) string {
  params = c.MergeSearchRequestParams(params)
  url := fmt.Sprintf("%s?%s", c.SearchUrl(), params.Encode())
  return url
}

func (c client) BuildLyricsUrl(params url.Values) string {
  params = c.MergeLyricsRequestParams(params)
  url := fmt.Sprintf("%s?%s", c.LyricsUrl(), params.Encode())
  return url
}

func (c client) Get(url string) (*http.Response, error) {
  return c.httpClient.Get(url)
}

func (c client) ParseSearchResponseBody(body []byte) (SearchResponse, error) {
  response := SearchResponse{}
  err := json.Unmarshal(body, &response)
  return response, err
}

func (c client) ParseLyricsResponseBody(body []byte) (LyricsResponse, error) {
  response := LyricsResponse{}
  err := json.Unmarshal(body, &response)
  return response, err
}

func (responseError ResponseError) Error() string {
  return fmt.Sprintf("%v - %v - %v", responseError.Code, responseError.Description, responseError.Message)
}

func (c client) ReadSearchResponse(res *http.Response) (SearchResponse, error) {
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return SearchResponse{}, err
  }
  searchResponse, err := c.ParseSearchResponseBody(body)
  if searchResponse.Response.Code != 100 {
    err = ResponseError{searchResponse.Response.Code, searchResponse.Response.Description, searchResponse.Response.Message}
  }

  return searchResponse, err
}

func (c client) ReadLyricsResponse(res *http.Response) (LyricsResponse, error) {
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return LyricsResponse{}, err
  }
  lyricsResponse, err := c.ParseLyricsResponseBody(body)
  if lyricsResponse.Response.Code < 101 || lyricsResponse.Response.Code > 111 {
    err = ResponseError{lyricsResponse.Response.Code, lyricsResponse.Response.Description, lyricsResponse.Response.Message}
  }

  return lyricsResponse, err
}

func (c client) SearchByArtistAndTrack(artist, track string) (SearchResponse, error) {
  params := make(url.Values)
  params.Set("artist", artist)
  params.Set("track", track)
  url := c.BuildSearchUrl(params)
  httpResponse, err := c.Get(url)
  if err != nil {
    return SearchResponse{}, err
  }

  return c.ReadSearchResponse(httpResponse)
}

func (c client) GetLyrics(trackId string, userAgent string) (LyricsResponse, error) {
  params := make(url.Values)
  params.Set("trackid", trackId)
  params.Set("useragent", userAgent)
  url := c.BuildLyricsUrl(params)
  httpResponse, err := c.Get(url)
  if err != nil {
    return LyricsResponse{}, err
  }

  return c.ReadLyricsResponse(httpResponse)
}

func NewClient() *client {
  c := &client{}
  c.httpClient = &http.Client{}
  return c
}
