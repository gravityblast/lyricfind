package lyricfind

import (
  "testing"
  "net/url"
  "fmt"
  assert "github.com/pilu/miniassert"
)

func TestLyricsUrl(t *testing.T) {
  c := &client{ SearchApiKey: "XYZ" }
  expectedUrl := "http://test.lyricfind.com/api_service/lyric.do"
  assert.Equal(t, expectedUrl, c.LyricsUrl())
}

func TestMergeLyricsRequestParams(t *testing.T) {
  c := &client{ SearchApiKey: "XXX", DisplayApiKey: "YYY" }
  params := make(url.Values)
  params.Set("trackid", "amg:1")
  params.Set("useragent", "mozilla")
  expectedQueryString := "apikey=YYY&output=json&reqtype=default&trackid=amg%3A1&useragent=mozilla"
  values := c.MergeLyricsRequestParams(params)
  assert.Equal(t, expectedQueryString, values.Encode())
}

func TestBuildLyricsUrl(t *testing.T) {
  c := &client{ SearchApiKey: "XXX", DisplayApiKey: "YYY" }
  params := make(url.Values)
  params.Set("trackid", "amg:1")
  queryString := "apikey=YYY&output=json&reqtype=default&trackid=amg%3A1"
  expectedUrl := fmt.Sprintf("%s?%s", c.LyricsUrl(), queryString)
  assert.Equal(t, expectedUrl, c.BuildLyricsUrl(params))
}

var lyricsResponseBodyFixture = []byte(`{
  "response": {
    "code": 101,
    "description": "SUCCESS: LICENSE, LYRICS"
  },
  "track": {
    "amg": 2,
    "instrumental": false,
    "viewable": true,
    "has_lrc": true,
    "title": "The Title",
    "artist": {
      "name": "The Artist"
    },
    "last_update": "2010-04-01 12:31:00",
    "lyrics": "Lorem ipsum dolor sit amet",
    "copyright": "Lyrics copy",
    "writer": "Writer 1 / Writer 2"
  }
}`)

/* func TestGetLyrics(t *testing.T) { */
/*   fakeClient := &FakeHttpClient{} */
/*   c := &client{ SearchApiKey: "XXX", httpClient: fakeClient } */
/*   c.GetLyrics("amg:1") */
/*   assert.Nil(t, err) */
/* } */

func TestParseLyricsResponseBody(t *testing.T) {
  c := client{}
  searchResponse, err := c.ParseLyricsResponseBody(lyricsResponseBodyFixture)
  assert.Nil(t, err)
  assert.Equal(t, 101, searchResponse.Response.Code)
  assert.Equal(t, "SUCCESS: LICENSE, LYRICS", searchResponse.Response.Description)
  assert.Equal(t, 2, searchResponse.Track.Amg)
  assert.False(t, searchResponse.Track.Instrumental)
  assert.True(t, searchResponse.Track.Viewable)
  assert.True(t, searchResponse.Track.Has_lrc)
  assert.Equal(t, "The Artist", searchResponse.Track.Artist.Name)
  assert.Equal(t, "2010-04-01 12:31:00", searchResponse.Track.Last_update)
  assert.Equal(t, "Lorem ipsum dolor sit amet", searchResponse.Track.Lyrics)
  assert.Equal(t, "Lyrics copy", searchResponse.Track.Copyright)
  assert.Equal(t, "Writer 1 / Writer 2", searchResponse.Track.Writer)
}

