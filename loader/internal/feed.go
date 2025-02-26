package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

const workers = 3

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
	pool := make(chan Episode, len(f.Episodes))
	done := make(chan bool)

	for i := range workers {
		// starts worker
		go func() {
			for {
				ep, more := <-pool
				if more {
					fmt.Printf("w: %d take ep: %d\n", i, ep.ID)
					ep.SaveAudio(folder)
				} else {
					done <- true
					return
				}
			}
		}()
	}

	// fills the pool
	for _, ep := range f.Episodes {
		pool <- ep
	}
	close(pool)

	// waiting for complete
	<-done
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
