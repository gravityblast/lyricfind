package lyricfind

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	assert "github.com/pilu/miniassert"
)

type fakeHttpClient struct {
	fixture string
	request *http.Request
}

func newFakeHttpClient(fixture string) *fakeHttpClient {
	return &fakeHttpClient{
		fixture: fixture,
	}
}

func (c *fakeHttpClient) Do(req *http.Request) (*http.Response, error) {
	if c.request != nil {
		log.Fatal("fakeHttpClient has already received a request")
	}
	c.request = req
	path := fmt.Sprintf("fixtures/%s.json", c.fixture)
	b, _ := os.Open(path)
	res := &http.Response{
		StatusCode: 200,
		Body:       b,
	}

	return res, nil
}

func TestEnvironment_Url(t *testing.T) {
	e := &Environment{Live: false}
	assert.Equal(t, "http://test.lyricfind.com/api_service", e.Url().String())

	e = &Environment{Live: true}
	assert.Equal(t, "http://api.lyricfind.com", e.Url().String())
}

func TestNew(t *testing.T) {
	c := New(&Environment{
		SearchApiKey:  "foo",
		DisplayApiKey: "bar",
	})
	assert.Equal(t, "foo", c.searchApiKey)
	assert.Equal(t, "bar", c.displayApiKey)
}

func TestClient_newRequest(t *testing.T) {
	c := New(&Environment{
		SearchApiKey:  "foo",
		DisplayApiKey: "bar",
	})
	v := url.Values{}
	v.Set("p1", "v1")
	r, err := c.newRequest("/foo", v)
	assert.Nil(t, err)
	expectedURI := "http://test.lyricfind.com/api_service/foo?output=json&p1=v1"
	assert.Equal(t, expectedURI, r.URL.String())
}

func TestClient_newSearchRequest(t *testing.T) {
	c := New(&Environment{
		SearchApiKey:  "foo",
		DisplayApiKey: "bar",
	})
	r, err := c.newSearchRequest(IPAddress("1.2.3.4"), &SearchOptions{Artist: "artist-foo"})
	assert.Nil(t, err)

	v := r.URL.Query()
	expected := url.Values{
		"output":     []string{"json"},
		"apikey":     []string{"foo"},
		"reqtype":    []string{"default"},
		"searchtype": []string{"track"},
		"artist":     []string{"artist-foo"},
		"ipaddress":  []string{"1.2.3.4"},
		"displaykey": []string{"bar"},
	}
	assert.Equal(t, expected, v)
}

func TestClient_newLyricRequest(t *testing.T) {
	c := New(&Environment{
		SearchApiKey:  "foo",
		DisplayApiKey: "bar",
	})
	r, err := c.newLyricRequest(IPAddress("1.2.3.4"), "Test Client", 1234)
	assert.Nil(t, err)

	v := r.URL.Query()
	expected := url.Values{
		"output":    []string{"json"},
		"apikey":    []string{"bar"},
		"reqtype":   []string{"default"},
		"trackid":   []string{"amg:1234"},
		"ipaddress": []string{"1.2.3.4"},
		"useragent": []string{"Test Client"},
	}
	assert.Equal(t, expected, v)
}

func TestClient_Search(t *testing.T) {
	c := New(&Environment{
		SearchApiKey:  "foo",
		DisplayApiKey: "bar",
		Live:          true,
	})
	fc := newFakeHttpClient("search")
	c.httpClient = fc
	sr, err := c.Search(IPAddress("1.2.3.4"), &SearchOptions{Artist: "artist-foo"})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(sr.Tracks))
	assert.Equal(t, "/search.do", fc.request.URL.Path)
}

func TestClient_Lyric(t *testing.T) {
	c := New(&Environment{
		SearchApiKey:  "foo",
		DisplayApiKey: "bar",
		Live:          true,
	})
	fc := newFakeHttpClient("lyric")
	c.httpClient = fc
	sr, err := c.Lyric(IPAddress("1.2.3.4"), "test client", 1234)
	assert.Nil(t, err)
	assert.Equal(t, "Lorem Ipsum", sr.Track.Title)
	assert.Equal(t, "/lyric.do", fc.request.URL.Path)
}
