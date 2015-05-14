package lyricfind

import (
	"net/url"
	"testing"

	assert "github.com/pilu/miniassert"
)

func TestSearchOptions_values(t *testing.T) {
	data := []struct {
		o *SearchOptions
		v url.Values
	}{
		{
			&SearchOptions{
				Artist: "artist-foo",
			},
			url.Values{
				"artist": []string{"artist-foo"},
			},
		},
		{
			&SearchOptions{
				Album: "album-foo",
			},
			url.Values{
				"album": []string{"album-foo"},
			},
		},
		{
			&SearchOptions{
				Track: "track-foo",
			},
			url.Values{
				"track": []string{"track-foo"},
			},
		},
		{
			&SearchOptions{
				Lyrics: "lyrics-foo",
			},
			url.Values{
				"lyrics": []string{"lyrics-foo"},
			},
		},
		{
			&SearchOptions{
				Meta: "meta-foo",
			},
			url.Values{
				"meta": []string{"meta-foo"},
			},
		},
		{
			&SearchOptions{
				All: "all-foo",
			},
			url.Values{
				"all": []string{"all-foo"},
			},
		},
		{
			&SearchOptions{
				OnlyDisplayable: true,
			},
			url.Values{
				"alltracks": []string{"no"},
			},
		},
		{
			&SearchOptions{
				Offset: 11,
			},
			url.Values{
				"offset": []string{"11"},
			},
		},
		{
			&SearchOptions{
				Limit: 12,
			},
			url.Values{
				"limit": []string{"12"},
			},
		},
	}

	for _, d := range data {
		assert.Equal(t, d.v, d.o.values())
	}
}
