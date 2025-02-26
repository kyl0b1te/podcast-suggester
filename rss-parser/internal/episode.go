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
	URL string `xml:"href,attr"`
}

type Audio struct {
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Episode struct {
	ID          int
	Link        string `xml:"link"`
	Title       string `xml:"title"`
	Description string `xml:"summary"`
	PublishedAt string `xml:"pubDate"`
	Image       Image  `xml:"image"`
	Audio       Audio  `xml:"enclosure"`
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
