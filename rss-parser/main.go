package main

import (
	"fmt"

	"github.com/kyl0b1te/rss-parser/internal"
)

func main() {
	feed, err := internal.NewFeed("https://anchor.fm/s/8e1e5620/podcast/rss")
	if err != nil {
		return
	}

	ep := feed.Episodes[0]
	fmt.Printf("%v", ep)
}
