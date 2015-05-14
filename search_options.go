package lyricfind

import (
	"fmt"
	"net/url"
)

type SearchOptions struct {
	Artist          string
	Album           string
	Track           string
	Lyrics          string
	Meta            string
	All             string
	OnlyDisplayable bool
	Offset          int
	Limit           int
}

func (o *SearchOptions) values() url.Values {
	v := url.Values{}

	if o.Artist != "" {
		v.Set("artist", o.Artist)
	}

	if o.Album != "" {
		v.Set("album", o.Album)
	}

	if o.Track != "" {
		v.Set("track", o.Track)
	}

	if o.Lyrics != "" {
		v.Set("lyrics", o.Lyrics)
	}

	if o.Meta != "" {
		v.Set("meta", o.Meta)
	}

	if o.All != "" {
		v.Set("all", o.All)
	}

	if o.OnlyDisplayable {
		v.Set("alltracks", "no")
	}

	if o.Offset != 0 {
		v.Set("offset", fmt.Sprintf("%d", o.Offset))
	}

	if o.Limit != 0 {
		v.Set("limit", fmt.Sprintf("%d", o.Limit))
	}

	return v
}
