# transcriber

## Poetry

`poetry shell`

## Commands

### Convert episodes for WAV

* `python cli.py convert --episode=../_data/episodes/236.mp3 --out=../_data/wav`
* `python cli.py convert --src=../_data/episodes --out=../_data/wav`

or shorter

* `python cli.py co -e ../_data/episodes/236.mp3 -o ../_data/wav`
* `python cli.py co -s ../_data/episodes -o ../_data/wav`

### Transcribe episodes

* `python cli.py transcribe --episode=../_data/wav/236.wav --out=../_data/text`
* `python cli.py transcribe --src=../_data/wav --out=../_data/text`

or shorter

* `python cli.py tr -e ../_data/wav/236.wav -o ../_data/text`
* `python cli.py tr -s ../_data/wav -o ../_data/text`
