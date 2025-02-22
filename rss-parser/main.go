package main

import (
	"fmt"

	"github.com/kyl0b1te/rss-parser/internal"
)

func main() {
	feed, err := internal.GetFeed("https://anchor.fm/s/8e1e5620/podcast/rss")
	if err != nil {
		return
	}

	for _, ep := range feed.Episodes {
		fmt.Printf("ep: %v\n", ep)
	}
}
