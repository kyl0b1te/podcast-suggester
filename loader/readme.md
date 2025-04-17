## loader

## Build

mac

`docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app -e GOOS=darwin -e GOARCH=arm64 golang:1.23-alpine go build -o bin/loader ./cmd/...`

or windows

`docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app -e GOOS=windows -e GOARCH=amd64 golang:1.23-alpine go build -o bin/loader ./cmd/...`

## Commands

* `loader cache -rss https://anchor.fm/s/8e1e5620/podcast/rss -out ../data`
* `loader audio -rss https://anchor.fm/s/8e1e5620/podcast/rss -out ../data -all`
* `loader audio -rss https://anchor.fm/s/8e1e5620/podcast/rss -out ../data -cache ../data/metadata.json`
