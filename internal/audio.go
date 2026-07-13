package internal

import (
	"io"
	"log"
	"os"

	"github.com/dhowden/tag"
	"github.com/gen2brain/beeep"
	"github.com/lizc2003/audioduration"
)

type AudioTags struct {
	Title string
	Artist string
	Album string
	AlbumArtist string
	Duration float64
}

func ReadTagsFromFile(audiofilePath string) (AudioTags, *tag.Picture) {
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

	duration, err := audioduration.Duration(file, audioduration.TypeMp3)
	if err != nil{
		duration = 0
		log.Fatal("Error while opening file: ", err)
	}

	audioTags := AudioTags{
		Title: tags.Title(),
		Artist: tags.Artist(),
		Album: tags.Album(),
		AlbumArtist: tags.AlbumArtist(),
		Duration: duration,
	}

	return audioTags, tags.Picture()
}