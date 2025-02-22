package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Image struct {
	URL string `xml:"href,attr"`
}

type Audio struct {
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Episode struct {
	Link        string `xml:"link"`
	Title       string `xml:"title"`
	Description string `xml:"summary"`
	PublishedAt string `xml:"pubDate"`
	Image       Image  `xml:"image"`
	Audio       Audio  `xml:"enclosure"`
}

type Feed struct {
	RSS      xml.Name  `xml:"rss"`
	Episodes []Episode `xml:"channel>item"`
}

func GetFeed(url string) (Feed, error) {
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

	return feed, nil
}
