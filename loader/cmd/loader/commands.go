package main

import (
	"flag"
	"fmt"
	"os"

	loader "github.com/kyl0b1te/loader/internal"
)

func cache() {
	cmd := flag.NewFlagSet("cache", flag.ExitOnError)
	url := cmd.String("rss", "", "show feed URL")
	out := cmd.String("out", "", "cache output folder path")

	cmd.Parse(os.Args[2:])
	if len(*url) == 0 {
		fmt.Println("error: rss url cannot be empty")
		os.Exit(1)
	}

	if len(*out) == 0 {
		fmt.Println("error: output path cannot be empty")
		os.Exit(1)
	}

	feed, err := loader.NewFeedFromURL(*url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cache, err := feed.Cache(*out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("cache has been stored in `%s` file\n", cache)
}

func audio() {
	cmd := flag.NewFlagSet("audio", flag.ExitOnError)
	url := cmd.String("rss", "", "show feed URL")
	out := cmd.String("out", "", "cache output folder path")
	all := cmd.Bool("all", false, "download all episodes or only latest missing (false by default)")
	cache := cmd.String("cache", "", "path to cached metadata file (used for upload latest)")

	cmd.Parse(os.Args[2:])
	if len(*url) == 0 {
		fmt.Println("error: rss url cannot be empty")
		os.Exit(1)
	}

	if len(*out) == 0 {
		fmt.Println("error: output path cannot be empty")
		os.Exit(1)
	}

	if !*all && len(*cache) == 0 {
		fmt.Println("error: path to cache file cannot be empty")
		os.Exit(1)
	}

	feed, err := loader.NewFeedFromURL(*url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *all {
		feed.SaveEpisodes(feed.Episodes, *out)
		feed.Cache(*out)
		os.Exit(0)
	}

	cacheFeed, err := loader.NewFeedFromCache(*cache)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	miss := feed.LastEpisodeId - cacheFeed.LastEpisodeId
	if miss == 0 {
		fmt.Println("there are no new episodes")
		os.Exit(0)
	}

	feed.SaveEpisodes(feed.Episodes[0:miss], *out)
	feed.Cache(*out)
	os.Exit(0)
}
