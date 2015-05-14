package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pilu/lyricfind"
)

const useragent = "LyricFind Golang Client"

type commandFunc func(*lyricfind.Client)

var commands = map[string]commandFunc{
	"search": search,
	"lyric":  lyric,
}

func getenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		fmt.Printf("you must set set `%s` env variable", k)
		os.Exit(1)
	}
	return v
}

func dumpTrack(t *lyricfind.Track) {
	fmt.Printf("AMG: %v\n", t.Amg)
	fmt.Printf("INSTRUMENTAL: %v\n", t.Instrumental)
	fmt.Printf("VIEWABLE: %v\n", t.Viewable)
	fmt.Printf("HASLRC: %v\n", t.HasLrc)
	fmt.Printf("TITLE: %v\n", t.Title)
	fmt.Printf("ARTIST: %v\n", t.Artist.Name)
	fmt.Printf("LYRICS: \n\n%v\n", t.Lyrics)
}

func search(c *lyricfind.Client) {
	territory := getenv("LYRICFIND_SEARCH_TERRITORY")

	so := &lyricfind.SearchOptions{}
	var list bool

	flag.StringVar(&so.Track, "t", "", "query to execute on track name")
	flag.StringVar(&so.Artist, "a", "", "query to execute on artist name")
	flag.StringVar(&so.Album, "b", "", "query to execute on album name")
	flag.BoolVar(&list, "list", false, "list results")
	flag.BoolVar(&so.OnlyDisplayable, "d", false, "search only on displayable tracks")
	flag.Parse()

	if !list {
		so.OnlyDisplayable = true
	}

	if flag.NFlag() == 0 {
		fmt.Println("At least one of the following flags is required:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	r, err := c.Search(lyricfind.Territory(territory), so)
	if err != nil {
		log.Fatal(err)
	}

	if list {
		for _, t := range r.Tracks {
			fmt.Printf("* %s - by %s [track: %d viewable: %v]\n", t.Title, t.Artist.Name, t.Amg, t.Viewable)
		}
		return
	}

	if len(r.Tracks) < 1 {
		fmt.Println("No results")
		os.Exit(0)
	}

	lr, err := c.Lyric(lyricfind.Territory(territory), useragent, r.Tracks[0].Amg)
	if err != nil {
		log.Fatal(err)
	}

	dumpTrack(lr.Track)
}

func lyric(c *lyricfind.Client) {
	territory := getenv("LYRICFIND_SEARCH_TERRITORY")
	if len(os.Args) < 2 {
		fmt.Printf("Usage: lyricfind lyric AMG\n")
		os.Exit(1)
	}

	amg, _ := strconv.Atoi(os.Args[1])
	lr, err := c.Lyric(lyricfind.Territory(territory), useragent, amg)
	if err != nil {
		log.Fatal(err)
	}

	dumpTrack(lr.Track)
}

func main() {
	searchKey := getenv("LYRICFIND_SEARCH_KEY")
	displayKey := getenv("LYRICFIND_DISPLAY_KEY")

	c := lyricfind.New(&lyricfind.Environment{
		SearchApiKey:  searchKey,
		DisplayApiKey: displayKey,
		Live:          true,
	})

	if len(os.Args) < 2 {
		fmt.Println("Usage: lyricfind COMMAND")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	os.Args = os.Args[1:]

	if cmd, ok := commands[cmdName]; ok {
		cmd(c)
		return
	}

	fmt.Printf("Unknown command `%s`\n", cmdName)
	fmt.Printf("Available commands:\n")
	for k, _ := range commands {
		fmt.Printf("  - %s\n", k)
	}
}
