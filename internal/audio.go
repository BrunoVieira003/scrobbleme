package internal

import (
	"io"
	"log"
	"os"

	// "github.com/bogem/id3v2"
	"github.com/dhowden/tag"
)

func ReadTagsFromFile(audiofilePath string) (string, string, string, string) {
	file, err := os.Open(audiofilePath)
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
	defer file.Close()

	var reader io.ReadSeeker = file
	tags, err := tag.ReadFrom(reader)
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}

	
	return tags.Title(), tags.Artist(), tags.Album(), tags.AlbumArtist()
}