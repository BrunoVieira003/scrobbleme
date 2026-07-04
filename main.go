package main

import (
	"fmt"
	"os"
	"scrobbleme/internal"
	"scrobbleme/internal/lastfm"
)

func main() {
	args := os.Args
	if len(args) < 2{
		fmt.Println("Scrobbleme")
		fmt.Println("Usage: scrobbleme <file-path>")
		return
	}

	config, loaded := internal.LoadConfig()
	if loaded{
		println(config.Session.Key)
		if config.Session.Key == ""{
			auth := lastfm.AuthenticateLastFM()
			config.Session = auth.Session
			internal.SaveConfig(config)
		}

		targetFile := args[1]
		title, artistTag := internal.ReadTagsFromFile(targetFile)

		lastfm.Scrobble(config.Session.Key, title, artistTag)
	}




	
}
