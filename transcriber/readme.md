# transcriber

## Development

`docker build -t podsu/transcriber:latest .`
`docker run --rm -it -v $PWD:/app -v $PWD/../data:/data -v $PWD/.cache:/root/.cache podsu/transcriber:latest bash`

## Commands

### Convert episodes for WAV

* `python cli.py convert --src=../data/episodes/1.mp3`
* `python cli.py convert --src=../data/episodes`

or shorter

* `python cli.py co -s ../data/episodes/1.mp3`
* `python cli.py co -s ../data/episodes`

### Transcribe episodes

* `python cli.py transcribe --ep=../_data/wav/236.wav --out=../_data/text`
* `python cli.py transcribe --src=../_data/wav --out=../_data/text`

or shorter

* `python cli.py tr -e ../_data/wav/236.wav -o ../_data/text`
* `python cli.py tr -s ../_data/wav -o ../_data/text`
