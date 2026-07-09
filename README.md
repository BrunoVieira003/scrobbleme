# Scrobbleme
Scroobleme is a simple software to add a Right-Click interaction to scrobble a mp3 file to your lastfm account

## How to use
### Installing
1. You can download the Windows installer from the [releases](https://github.com/BrunoVieira003/scrobbleme/releases) page and install

2. After installing, right-click some mp3 file and click on Scrobble me

3. If this is your first time, a browser tab will open asking for access to your account. Click in "Yes, allow access"
4. Right-click on any mp3 file and click on the option "Scrobble me", wait a few seconds and it's done

# Development
## Requirements
- Go 1.25+

## Building
```bash
go build -ldflags="-H=windowsgui" -o scrobbleme.exe
```
