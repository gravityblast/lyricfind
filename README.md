# lyricfind

Lyricfind client for the Go language

## Usage

### Import

    import ("github.com/pilu/lyricfind")

### Initializing a Lyricfind client

    client := lyricfind.NewClient()
    client.SearchApiKey   = "LYRICFIND_SEARCH_API_KEY"
    client.DisplayApiKey  = "LYRICFIND_DISPLAY_API_KEY"

### Search

    res, err := client.SearchByArtistAndTrack(artist, track)

    if err != nil {
      fmt.Printf("Error: %v\n", err)
    } else {
      for index, track := range  res.Tracks {
        fmt.Printf("Result %d:\n", index + 1)
        fmt.Printf("  Amg: %v\n", track.Amg)
        fmt.Printf("  Instrumental: %v\n", track.Instrumental)
        fmt.Printf("  Viewable: %v\n", track.Viewable)
        fmt.Printf("  Has lrc: %v\n", track.Has_lrc)
        fmt.Printf("  Title: %v\n", track.Title)
        fmt.Printf("  Artist: %v\n", track.Artist.Name)
        fmt.Printf("  Snippet: %v\n", track.Snippet)
        fmt.Printf("  Last update: %v\n", track.Last_update)
        fmt.Printf("  Score: %v\n", track.Score)
      }
    }

### Get lyrics

    res, err := client.GetLyrics(id, userAgent)

    if err != nil {
      fmt.Printf("Error: %v\n", err)
    } else {
      track := res.Track
      fmt.Printf("Amg: %v\n", track.Amg)
      fmt.Printf("Instrumental: %v\n", track.Instrumental)
      fmt.Printf("Viewable: %v\n", track.Viewable)
      fmt.Printf("Has lrc: %v\n", track.Has_lrc)
      fmt.Printf("Title: %v\n", track.Title)
      fmt.Printf("Artist: %v\n", track.Artist.Name)
      fmt.Printf("Last update: %v\n", track.Last_update)
      fmt.Printf("Copyright: %v\n", track.Copyright)
      fmt.Printf("Writer: %v\n", track.Writer)
      fmt.Printf("Lyrics:\n%v\n", track.Lyrics)
    }

### Search and Get Lyrics

    res, err := client.SearchAndGetLyrics(artist, track, userAgent)

    if err != nil {
      fmt.Printf("Error: %v\n", err)
    } else {
      track := res.Track
      fmt.Printf("Amg: %v\n", track.Amg)
      fmt.Printf("Instrumental: %v\n", track.Instrumental)
      fmt.Printf("Viewable: %v\n", track.Viewable)
      fmt.Printf("Has lrc: %v\n", track.Has_lrc)
      fmt.Printf("Title: %v\n", track.Title)
      fmt.Printf("Artist: %v\n", track.Artist.Name)
      fmt.Printf("Last update: %v\n", track.Last_update)
      fmt.Printf("Copyright: %v\n", track.Copyright)
      fmt.Printf("Writer: %v\n", track.Writer)
      fmt.Printf("Lyrics:\n%v\n", track.Lyrics)
    }

#### Authors

* Andrea Franz (http://gravityblast.com)
