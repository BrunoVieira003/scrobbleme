//go:generate goversioninfo -icon=assets/icon.ico

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"scrobbleme/internal"
	"scrobbleme/internal/lastfm"
	// "sync"

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

	// wg := sync.WaitGroup{}

    // log.SetOutput(logFile)

	config, loaded := internal.LoadConfig()
	if loaded{
		println(config.Session.Key)
		if config.Session.Key == ""{
			auth := lastfm.AuthenticateLastFM()
			config.Session = auth.Session
			internal.SaveConfig(config)
		}

		targetFile := args[1]
		tags, picture := internal.ReadTagsFromFile(targetFile)

		lastfm.Scrobble(config.Session.Key, tags)

		var ntfyPicture []byte
		if picture != nil{
			ntfyPicture = picture.Data
		}


		beeep.Notify("Scrobbled", tags.Title+" | "+tags.Artist, ntfyPicture)
		log.Println("Scrobble", "track:", tags.Title, "artist:", tags.Artist, "album:", tags.Album, "albumArtist:", tags.AlbumArtist)
	}




	
}
