package lyricfind

import (
  "testing"
  "net/url"
  assert "github.com/pilu/miniassert"
)

func TestSearchUrl(t *testing.T) {
  c := &client{}
  expectedUrl := "http://test.lyricfind.com/api_service/search.do"
  assert.Equal(t, expectedUrl, c.SearchUrl())
}


func TestMergeDefaultRequestParams(t *testing.T) {
  c := &client{}
  params := make(url.Values)
  params.Set("artist", "foo")
  expectedQueryString := "artist=foo&output=json"
  values := c.MergeDefaultRequestParams(params)
  assert.Equal(t, expectedQueryString, values.Encode())
}

func TestMergeSearchRequestParams(t *testing.T) {
  c := &client{}
  params := make(url.Values)
  params.Set("artist", "foo")
  params.Set("track", "bar")
  expectedQueryString := "artist=foo&output=json&reqtype=default&searchtype=track&track=bar"
  values := c.MergeSearchRequestParams(params)
  assert.Equal(t, expectedQueryString, values.Encode())
}
