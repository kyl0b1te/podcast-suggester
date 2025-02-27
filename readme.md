# podcast-suggester

## Build

`go build -o bin/loader ./cmd/...`

## Commands

`loader cache -rss https://anchor.fm/s/8e1e5620/podcast/rss -out ./`
`loader audio -rss https://anchor.fm/s/8e1e5620/podcast/rss -out ./ -all`
`loader audio -rss https://anchor.fm/s/8e1e5620/podcast/rss -out ./ -cache ./meta.json`
