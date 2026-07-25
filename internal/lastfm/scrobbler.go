package lastfm

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"scrobbleme/internal"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

type Scrobbler struct {
	ApiKey       string
	SharedSecret string
	SessionKey   string
	Token        string
	tracks       []string
	artists      []string
	albums       []string
	albumArtists []string
	durations    []float64
}

func (s Scrobbler) Scrobble() {
	now := time.Now().UTC()
	unixSecs := now.Unix() - int64(s.durations[0])
	timestampStr := strconv.FormatInt(unixSecs, 10)

	form := s.scrobbleForm(timestampStr)

	resp, err := http.PostForm("https://ws.audioscrobbler.com/2.0", form)
	if err != nil {
		beeep.Notify("Failed to scrobble", err.Error(), "")
		log.Fatal("Failed to scrobble")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		beeep.Notify("Failed to scrobble", resp.Status, "")
		log.Fatal("Failed to scrobble")
	}
}

func (s *Scrobbler) AddTrack(track string, artist string, album string, albumArtist string, duration float64) {
	s.tracks = append(s.tracks, track)
	s.artists = append(s.artists, artist)
	s.albums = append(s.albums, album)
	s.albumArtists = append(s.albumArtists, albumArtist)
	s.durations = append(s.durations, duration)
}

func (s *Scrobbler) FlushTracks() {
	s.tracks = make([]string, 0)
	s.artists = make([]string, 0)
	s.albums = make([]string, 0)
	s.albumArtists = make([]string, 0)
	s.durations = make([]float64, 0)
}

func (s Scrobbler) First() internal.AudioTags{
	return internal.AudioTags{
		Title: s.tracks[0],
		Artist: s.artists[0],
		Album: s.artists[0],
		AlbumArtist: s.albumArtists[0],
		Duration: s.durations[0],
	}
}

func (s Scrobbler) Lenght() int{
	return len(s.tracks)
}

func (s Scrobbler) scrobbleSignature(timestamp string) string {
	builder := strings.Builder{}

	for i, albumArtist := range s.albumArtists {
		if albumArtist != "" {
			builder.WriteString(fmt.Sprintf("albumArtist[%d]", i))
			builder.WriteString(albumArtist)
		}
	}

	for i, album := range s.albums {
		if album != "" {
			builder.WriteString(fmt.Sprintf("album[%d]", i))
			builder.WriteString(album)
		}
	}

	builder.WriteString("api_key")
	builder.WriteString(s.ApiKey)

	for i, artist := range s.artists {
		builder.WriteString(fmt.Sprintf("artist[%d]", i))
		builder.WriteString(artist)
	}

	builder.WriteString("method")
	builder.WriteString("track.scrobble")

	builder.WriteString("sk")
	builder.WriteString(s.SessionKey)

	for i := range s.tracks {
		builder.WriteString(fmt.Sprintf("timestamp[%d]", i))
		builder.WriteString(timestamp)
	}

	for i, tracks := range s.tracks {
		builder.WriteString(fmt.Sprintf("track[%d]", i))
		builder.WriteString(tracks)
	}

	builder.WriteString(s.SharedSecret)

	signatureBase := builder.String()
	hasher := md5.New()

	io.WriteString(hasher, signatureBase)
	md5String := hex.EncodeToString(hasher.Sum(nil))

	return md5String
}

func (s Scrobbler) SessionSignature() string {
	builder := strings.Builder{}

	builder.WriteString("api_key")
	builder.WriteString(s.ApiKey)

	builder.WriteString("method")
	builder.WriteString("auth.getSession")

	builder.WriteString("token")
	builder.WriteString(s.Token)

	builder.WriteString(s.SharedSecret)

	signatureBase := builder.String()
	hasher := md5.New()
	io.WriteString(hasher, signatureBase)
	md5String := hex.EncodeToString(hasher.Sum(nil))

	return md5String
}

func (s Scrobbler) scrobbleForm(timestamp string) url.Values {
	form := url.Values{}

	for i, album := range s.albums {
		if album != "" {
			form.Add(fmt.Sprintf("album[%d]", i), album)
		}
	}

	for i, albumArtist := range s.albumArtists {
		if albumArtist != "" {
			form.Add(fmt.Sprintf("albumArtist[%d]", i), albumArtist)
		}
	}

	form.Add("api_key", s.ApiKey)

	for i, artist := range s.artists {
		form.Add(fmt.Sprintf("artist[%d]", i), artist)
	}

	form.Add("method", "track.scrobble")

	form.Add("sk", s.SessionKey)

	for i := range s.tracks {
		form.Add(fmt.Sprintf("timestamp[%d]", i), timestamp)
	}

	for i, track := range s.tracks {
		form.Add(fmt.Sprintf("track[%d]", i), track)
	}

	form.Add("format", "json")
	form.Add("api_sig", s.scrobbleSignature(timestamp))

	return form
}
