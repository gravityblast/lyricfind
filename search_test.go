package lyricfind

import (
  "testing"
  "net/url"
  "fmt"
  assert "github.com/pilu/miniassert"
)

func TestSearchUrl(t *testing.T) {
  c := &Client{ SearchApiKey: "XYZ" }
  expectedUrl := "http://test.lyricfind.com/api_service/search.do"
  assert.Equal(t, expectedUrl, c.SearchUrl())
}

func TestMergeSearchRequestParams(t *testing.T) {
  c := &Client{ SearchApiKey: "XYZ" }
  params := make(url.Values)
  params.Set("artist", "foo")
  params.Set("track", "bar")
  expectedQueryString := "apikey=XYZ&artist=foo&output=json&reqtype=default&searchtype=track&track=bar"
  values := c.MergeSearchRequestParams(params)
  assert.Equal(t, expectedQueryString, values.Encode())
}

func TestBuildSearchUrl(t *testing.T) {
  c := &Client{ SearchApiKey: "XYZ" }
  params := make(url.Values)
  params.Set("artist", "foo")
  params.Set("track", "bar")
  queryString := "apikey=XYZ&artist=foo&output=json&reqtype=default&searchtype=track&track=bar"
  expectedUrl := fmt.Sprintf("%s?%s", c.SearchUrl(), queryString)
  assert.Equal(t, expectedUrl, c.BuildSearchUrl(params))
}

var searchResponseBodyFixture = []byte(`{
  "response": {
    "code": 100,
    "description": "SUCCESS"
  },
  "totalresults": 181,
  "totalpages": 19,
  "tracks": [
    {
      "amg": 1,
      "instrumental": false,
      "viewable": true,
      "has_lrc": true,
      "title": "Title 1",
      "artist": {
        "name": "Artist 1"
      },
      "snippet": "snippet 1",
      "last_update": "2012-03-26 08:00:00",
      "score": 18.9
    },
    {
      "amg": 2,
      "instrumental": false,
      "viewable": true,
      "has_lrc": true,
      "title": "Title 2",
      "artist": {
        "name": "Artist 2"
      },
      "snippet": "snippet 2",
      "last_update": "2012-04-26 08:00:00",
      "score": 15
    }
  ]
}`)

/* func TestSearchByArtistAndTrack(t *testing.T) { */
/*   fakeClient := &FakeHttpClient{} */
/*   c := &Client{ SearchApiKey: "XXX", httpClient: fakeClient } */
/*   _, err := c.SearchByArtistAndTrack("foo", "bar") */
/*   assert.Nil(t, err) */
/* } */

func TestParseSearchResponseBody(t *testing.T) {
  c := Client{}
  searchResponse, err := c.ParseSearchResponseBody(searchResponseBodyFixture)
  assert.Nil(t, err)
  assert.Equal(t, 100, searchResponse.Response.Code)
  assert.Equal(t, "SUCCESS", searchResponse.Response.Description)
  assert.Equal(t, 181, searchResponse.Totalresults)
  assert.Equal(t, 19, searchResponse.Totalpages)
  assert.Equal(t, 2, len(searchResponse.Tracks))
}

func TestParseSearchResponseBody_TrackFields(t *testing.T) {
  c := Client{}
  searchResponse, err := c.ParseSearchResponseBody(searchResponseBodyFixture)
  assert.Nil(t, err)

  var track Track

  track = searchResponse.Tracks[0]
  assert.Equal(t, 1, track.Amg)
  assert.False(t, track.Instrumental)
  assert.True(t, track.Viewable)
  assert.True(t, track.Has_lrc)
  assert.Equal(t, "Title 1", track.Title)
  assert.Equal(t, "Artist 1", track.Artist.Name)
  assert.Equal(t, "snippet 1", track.Snippet)
  assert.Equal(t, "2012-03-26 08:00:00", track.Last_update)
  assert.Equal(t, 18.9, track.Score)

  track = searchResponse.Tracks[1]
  assert.Equal(t, 2, track.Amg)
  assert.False(t, track.Instrumental)
  assert.True(t, track.Viewable)
  assert.True(t, track.Has_lrc)
  assert.Equal(t, "Title 2", track.Title)
  assert.Equal(t, "Artist 2", track.Artist.Name)
  assert.Equal(t, "snippet 2", track.Snippet)
  assert.Equal(t, "2012-04-26 08:00:00", track.Last_update)
  assert.Equal(t, 15.0, track.Score)
}

