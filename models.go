package lyricfind

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
  Score float64
  Lyrics string
  Copyright string
  Writer string
}

type Response struct {
  Code int
  Description string
  Message string
}

type SearchResponse struct {
  Response Response
  Totalresults int
  Totalpages int
  Tracks []Track
}

type LyricsResponse struct {
  Response Response
  Totalresults int
  Totalpages int
  Track Track
}
