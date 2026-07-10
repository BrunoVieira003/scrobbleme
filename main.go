//go:generate goversioninfo -icon=assets/icon.ico

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"scrobbleme/internal"
	"scrobbleme/internal/lastfm"

	"github.com/gen2brain/beeep"
)

func main() {
	args := os.Args
	if len(args) < 2{
		fmt.Println("Scrobbleme")
		fmt.Println("Usage: scrobbleme <file-path>")
		return
	}

	beeep.AppName = "Scrobbleme"
	
	config_dir, _ := os.UserConfigDir()
	logFilepath := path.Join(config_dir, "Scrobbleme", "logs.txt")

	logFile, err := os.OpenFile(logFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()

    log.SetOutput(logFile)

	config, loaded := internal.LoadConfig()
	if loaded{
		println(config.Session.Key)
		if config.Session.Key == ""{
			auth := lastfm.AuthenticateLastFM()
			config.Session = auth.Session
			internal.SaveConfig(config)
		}

		targetFile := args[1]
		title, artistTag, album, albumArtist := internal.ReadTagsFromFile(targetFile)

		lastfm.Scrobble(config.Session.Key, title, artistTag, album, albumArtist)

		beeep.Notify("Scrobbled", title+" | "+artistTag, "")
		log.Println("Scrobble", "track:", title, "artist:", artistTag, "album:", album, "albumArtist:", albumArtist)
	}




	
}
