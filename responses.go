package lyricfind

type Artist struct {
	Name string `json:"name"`
}

type Track struct {
	Amg          int     `json:"amg"`
	Instrumental bool    `json:"instrumental"`
	Viewable     bool    `json:"viewable"`
	HasLrc       bool    `json:"has_lrc"`
	Title        string  `json:"title"`
	Artist       *Artist `json:"artist"`
	Snippet      string  `json:"snippet"`
	LastUpdate   string  `json:"last_update"`
	Score        float64 `json:"score"`
	Lyrics       string  `json:"lyrics"`
	Copyright    string  `json:"copyright"`
	Writer       string  `json:"writer"`
}

type Response struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Message     string `json:"message"`
}

type SearchResponse struct {
	Response     *Response `json:"response"`
	TotalResults int       `json:"total_results"`
	TotalPages   int       `json:"total_pages"`
	Tracks       []*Track  `json:"tracks"`
}

type LyricsResponse struct {
	Response *Response `json:"response"`
	Track    *Track    `json:"track"`
}
