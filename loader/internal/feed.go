package internal

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const workers = 3

type Feed struct {
	RSS           xml.Name  `xml:"rss" json:"-"`
	URL           string    `xml:"-" json:"url"`
	Title         string    `xml:"channel>title" json:"title"`
	LastEpisodeId int       `json:"lastEpisodeId"`
	Episodes      []Episode `xml:"channel>item" json:"episodes"`
}

func (f *Feed) init() {
	count := len(f.Episodes)
	for index := range f.Episodes {
		f.Episodes[index].ID = count
		count -= 1
	}

	f.LastEpisodeId = len(f.Episodes)
}

func (f *Feed) Cache(folder string) (string, error) {
	cacheFilePath := filepath.Join(folder, "metadata.json")

	json, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to json feed: %w", err)
	}

	err = os.WriteFile(cacheFilePath, json, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save cache file: %w", err)
	}

	return cacheFilePath, nil
}

func (f *Feed) SaveEpisodes(episodes []Episode, out string) {
	pool := make(chan Episode, len(episodes))
	done := make(chan bool)

	for i := range workers {
		// starts worker
		go func() {
			for {
				ep, more := <-pool
				if more {
					fmt.Printf("w: %d take ep: %d\n", i, ep.ID)
					ep.SaveAudio(out)
				} else {
					done <- true
					return
				}
			}
		}()
	}

	// fills the pool
	for _, ep := range episodes {
		pool <- ep
	}
	close(pool)

	// waiting for complete
	<-done
}

func NewFeedFromURL(url string) (Feed, error) {
	feed := Feed{}
	feed.URL = url

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

func NewFeedFromCache(path string) (Feed, error) {
	feed := Feed{}

	file, err := os.ReadFile(path)
	if err != nil {
		return feed, fmt.Errorf("failed to read metadata file: %w", err)
	}

	err = json.Unmarshal(file, &feed)
	if err != nil {
		return feed, fmt.Errorf("failed to parse metadata: %w", err)
	}

	return feed, nil
}
