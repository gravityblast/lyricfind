package lyricfind

import (
  "testing"
  "net/url"
  "net/http"
  assert "github.com/pilu/miniassert"
)

func TestNewClient(t *testing.T) {
  c := NewClient()
  assert.Type(t, "*http.Client", c.httpClient)
}

func TestMergeDefaultRequestParams(t *testing.T) {
  c := &Client{ SearchApiKey: "XYZ" }
  params := make(url.Values)
  params.Set("artist", "foo")
  expectedQueryString := "artist=foo&output=json"
  values := c.MergeDefaultRequestParams(params)
  assert.Equal(t, expectedQueryString, values.Encode())
}

type FakeHttpClient struct {
  getCount int
  lastUrl string
}

func (httpClient *FakeHttpClient) Get(url string) (*http.Response, error) {
  httpClient.getCount++
  httpClient.lastUrl = url
  response := &http.Response{}
  return response, nil
}

func TestGet(t *testing.T) {
  fakeClient := &FakeHttpClient{}
  c := &Client{ httpClient: fakeClient }
  c.Get("http://foo.bar")
  assert.Equal(t, 1, fakeClient.getCount)
  assert.Equal(t, "http://foo.bar", fakeClient.lastUrl)
}

