package lastfm

import (
	"log"
	"net/http"
	"net/url"
	"scrobbleme/internal"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func Scrobble(sessionKey string, tags internal.AudioTags) {
	now := time.Now().UTC()
	unixSecs := now.Unix() - int64(tags.Duration)
	timestampStr := strconv.FormatInt(unixSecs, 10)

	api_sig := GenerateSigForScrobble(sessionKey, timestampStr, tags.Title, tags.Artist, tags.Album, tags.AlbumArtist)

	form := url.Values{
		"method":    {"track.scrobble"},
		"api_key":   {internal.LASTFM_KEY},
		"artist":    {tags.Artist},
		"track":     {tags.Title},
		"timestamp": {timestampStr},
		"sk":        {sessionKey},
		"api_sig":   {api_sig},
		"format":    {"json"},
	}

	if tags.Album != "" && tags.Artist != "" {
		form.Set("album", tags.Album)
		form.Set("albumArtist", tags.Artist)
	}

	resp, err := http.PostForm("https://ws.audioscrobbler.com/2.0", form)
	if err != nil {
		beeep.Notify("Failed to scrobble", err.Error(), "")
		log.Fatal("Failed to scrobble")
	}
	defer resp.Body.Close()
}
