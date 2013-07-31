package lyricfind

import (
  "fmt"
  "net/url"
)

const BASE_URL = "http://test.lyricfind.com/api_service"

type client struct {

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
  return c.MergeDefaultRequestParams(params)
}
