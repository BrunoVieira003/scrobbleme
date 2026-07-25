//go:generate goversioninfo -icon=assets/icon.ico

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"scrobbleme/internal"
	"scrobbleme/internal/lastfm"
	"strings"

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

    log.SetOutput(logFile)

	config, loaded := internal.LoadConfig()
	if loaded{
		println(config.Session.Key)
		if config.Session.Key == ""{
			auth := lastfm.AuthenticateLastFM()
			config.Session = auth.Session
			internal.SaveConfig(config)
		}

		var targetPaths []string

		for _, arg := range args[1:]{
			info, _ := os.Stat(arg)
			if(info.IsDir()){
				items, _ := os.ReadDir(arg)
				for _, item := range items{
					if !item.IsDir() && strings.HasSuffix(item.Name(), ".mp3"){
						targetPaths = append(targetPaths, path.Join(arg, item.Name()))
					}
				}
			}else{
				targetPaths = append(targetPaths, arg)
			}
		}

		

		scrobbler := lastfm.Scrobbler{
			ApiKey:       internal.LASTFM_KEY,
			SharedSecret: internal.LASTFM_SECRET,
			SessionKey: config.Session.Key,

		}

		var ntfyPicture []byte

		for _, path := range targetPaths{
			tags, picture := internal.ReadTagsFromFile(path)
			scrobbler.AddTrack(tags.Title, tags.Artist, tags.Album, tags.AlbumArtist, tags.Duration)

			if picture != nil{
				ntfyPicture = picture.Data
			}
		}

		scrobbler.Scrobble()

		if(scrobbler.Lenght() > 1){
			beeep.Notify("Scrobbled", fmt.Sprintf("Scrobbled %d songs", scrobbler.Lenght()), ntfyPicture)
		}else {
			tags := scrobbler.First()
			beeep.Notify("Scrobbled", tags.Title+" | "+tags.Artist, ntfyPicture)
			log.Println("Scrobble", "track:", tags.Title, "artist:", tags.Artist, "album:", tags.Album, "albumArtist:", tags.AlbumArtist)
		}

	}




	
}
