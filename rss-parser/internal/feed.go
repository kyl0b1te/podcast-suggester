package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Feed struct {
	RSS      xml.Name  `xml:"rss"`
	Title    string    `xml:"channel>title"`
	Episodes []Episode `xml:"channel>item"`
}

func (f *Feed) init() {
	count := len(f.Episodes)
	for index := range f.Episodes {
		f.Episodes[index].ID = count
		count -= 1
	}
}

func (f *Feed) SaveLatest(folder string) error {
	_, err := f.Episodes[0].SaveAudio(folder)
	return err
}

func (f *Feed) SaveAll(folder string) {
	load := make(chan int, 2)

	// todo : refactor
	for i := range 2 {
		go func() {
			_, err := f.Episodes[i].SaveAudio(folder)
			if err != nil {
				panic(err)
			}
			load <- i
		}()
	}

	for range 2 {
		fmt.Printf("processed: %d\n", <-load)
	}
}

func NewFeed(url string) (Feed, error) {
	feed := Feed{}

	resp, err := http.Get(url)
	if err != nil {
		return feed, fmt.Errorf("failed to load feed from URL '%s': %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return feed, fmt.Errorf("failed to read body: %w", err)
	}

	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return Feed{}, fmt.Errorf("failed to parse xml: %w", err)
	}

	feed.init()
	return feed, nil
}
