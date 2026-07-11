package internal

import (
	"io"
	"log"
	"os"

	"github.com/dhowden/tag"
	"github.com/gen2brain/beeep"
)

func ReadTagsFromFile(audiofilePath string) (string, string, string, string, *tag.Picture) {
	file, err := os.Open(audiofilePath)
	if err != nil {
		beeep.Notify("Failed to scrobble", err.Error(), "")
		log.Fatal("Error while opening file: ", err)
	}
	defer file.Close()

	var reader io.ReadSeeker = file
	tags, err := tag.ReadFrom(reader)
	if err != nil {
		beeep.Notify("Failed to scrobble", err.Error(), "")
		log.Fatal("Error while opening file: ", err)
	}

	
	return tags.Title(), tags.Artist(), tags.Album(), tags.AlbumArtist(), tags.Picture()
}