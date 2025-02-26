package main

import (
	"fmt"

	loader "github.com/kyl0b1te/loader/internal"
)

func main() {
	feed, err := loader.NewFeed("https://anchor.fm/s/8e1e5620/podcast/rss")
	if err != nil {
		return
	}

	ep := feed.Episodes[0]
	fmt.Printf("%v", ep)
}
