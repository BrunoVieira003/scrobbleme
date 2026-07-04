# Scrobbleme
Scroobleme is a simple software to add a Right-Click interaction to scrobble a mp3 file to your lastfm account

## How to use
### Installing
1. You can download the Windows installer from the releases page and execute it

2. After installing, right-click some mp3 file and click on Scrobble me

3. If this is your first time, a browser tab will open asking for access to your account. Click in "Yes, allow access"

# Development
## Compiling to exe
```bash
go build -ldflags="-H=windowsgui" -o scrobbleme.exe
```