package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Image struct {
	URL string `xml:"href,attr" json:"url"`
}

type Audio struct {
	URL    string `xml:"url,attr" json:"url"`
	Length string `xml:"length,attr" json:"length"`
	Type   string `xml:"type,attr" json:"type"`
}

type Episode struct {
	ID          int    `json:"id"`
	Link        string `xml:"link" json:"link"`
	Title       string `xml:"title" json:"title"`
	Description string `xml:"summary" json:"description"`
	PublishedAt string `xml:"pubDate" json:"publishedAt"`
	Image       Image  `xml:"image" json:"image"`
	Audio       Audio  `xml:"enclosure" json:"audio"`
}

func (e *Episode) getAudioFilePath(folder string) string {
	src := strings.Split(e.Audio.URL, ".")
	return filepath.Join(folder, fmt.Sprintf("%d.%s", e.ID, src[len(src)-1]))
}

func (e *Episode) SaveAudio(folder string) (string, error) {
	resp, err := http.Get(e.Audio.URL)
	if err != nil {
		return "", fmt.Errorf("failed to download episode audio from URL '%s': %w", e.Audio.URL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("audio resource is not accessible '%d': %w", resp.StatusCode, err)
	}

	audioFilePath := e.getAudioFilePath(folder)

	dist, err := os.Create(audioFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to create a file ('%s') in folder '%s': %w", audioFilePath, folder, err)
	}
	defer dist.Close()

	_, err = io.Copy(dist, resp.Body)
	return audioFilePath, err
}
