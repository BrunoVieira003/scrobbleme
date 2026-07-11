# Scrobbleme
Scroobleme is a simple software to add a Right-Click interaction to scrobble a mp3 file to your lastfm account

## How to use
### Installing
1. You can download the Windows installer from the [releases](https://github.com/BrunoVieira003/scrobbleme/releases) page and install

2. Right-click on any mp3 file and click on the option "Scrobble me" and wait a few seconds

3. If this is your first time, a browser tab will open asking for access to your account. Click in "Yes, allow access"
<img width="824" height="294" alt="image" src="https://github.com/user-attachments/assets/1ba58654-d31d-4119-ae6e-4da65e7a424e" />

4. Next time you right-click the same option, it will scrobble the song without asking for authorization again

# Development
## Requirements
- Go 1.25+

## Building
```bash
go build -ldflags="-H=windowsgui -s -w -X scrobbleme/internal.LASTFM_KEY=<your lastfm api key> -X scrobbleme/internal.LASTFM_SECRET=<your lastfm secret>" -o scrobbleme.exe
```
