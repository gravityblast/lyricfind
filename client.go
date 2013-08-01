package lyricfind

import (
  "fmt"
  "net/url"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

const BASE_URL = "http://test.lyricfind.com/api_service"

type HttpClient interface {
  Get(string) (*http.Response, error)
}

type client struct {
  SearchApiKey string
  httpClient HttpClient
}

type Artist struct {
  Name string
}

type Track struct {
  Amg int
  Instrumental bool
  Viewable bool
  Has_lrc bool
  Title string
  Artist Artist
  Snippet string
  Last_update string
  Score int
}

type Response struct {
  Code int
  Description string
}

type SearchResponse struct {
  Response Response
  Totalresults int
  Totalpages int
  Tracks []Track
}

func (c client) SearchUrl() string {
  return fmt.Sprintf("%s/%s", BASE_URL, "search.do")
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

func (c client) BuildSearchUrl(params url.Values) string {
  params = c.MergeSearchRequestParams(params)
  url := fmt.Sprintf("%s?%s", c.SearchUrl(), params.Encode())
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

func (c client) SearchByArtistAndTrack(artist, track string) (SearchResponse, error) {
  params := make(url.Values)
  params.Set("artist", artist)
  params.Set("track", track)
  url := c.BuildSearchUrl(params)
  resp, err := c.Get(url)
  if err != nil {
    return SearchResponse{}, err
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return SearchResponse{}, err
  }
  searchResponse, err := c.ParseSearchResponseBody(body)
  return searchResponse, err
}

func NewClient() *client {
  c := &client{}
  c.httpClient = &http.Client{}
  return c
}
