package lastfm

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func Scrobble(sessionKey string, track string, artist string, album string, albumArtist string) {
	now := time.Now().UTC()
	unixSecs := now.Unix()
	timestampStr := strconv.FormatInt(unixSecs, 10)

	api_sig := GenerateSigForScrobble(sessionKey, timestampStr, track, artist, album, albumArtist)

	form := url.Values{
		"method":      {"track.scrobble"},
		"api_key":     {API_KEY},
		"artist":      {artist},
		"track":       {track},
		"timestamp":   {timestampStr},
		"sk":          {sessionKey},
		"api_sig":     {api_sig},
		"format":      {"json"},
	}

	if album != "" && albumArtist != ""{
		form.Set("album", album)
		form.Set("albumArtist", albumArtist)
	}

	resp, err := http.PostForm("https://ws.audioscrobbler.com/2.0", form)
	if err != nil {
		beeep.Notify("Failed to scrobble", err.Error(), "")
		log.Fatal("Failed to scrobble")
	}
	defer resp.Body.Close()
}
