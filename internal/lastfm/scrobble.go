package lastfm

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Scrobble(sessionKey string, track string, artist string, album string, albumArtist string) {
	now := time.Now().UTC()
	unixSecs := now.Unix()
	timestampStr := strconv.FormatInt(unixSecs, 10)

	api_sig := GenerateSigForScrobble(sessionKey, timestampStr, track, artist, album, albumArtist)

	form := url.Values{
		"method":      {"track.scrobble"},
		"album":       {album},
		"albumArtist": {albumArtist},
		"api_key":     {API_KEY},
		"artist":      {artist},
		"track":       {track},
		"timestamp":   {timestampStr},
		"sk":          {sessionKey},
		"api_sig":     {api_sig},
		"format":      {"json"},
	}

	resp, err := http.PostForm("https://ws.audioscrobbler.com/2.0", form)
	if err != nil {
		log.Fatal("Failed to scrobble")
	}

	println(resp.StatusCode)

	defer resp.Body.Close()
}
