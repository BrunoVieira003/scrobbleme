package lastfm

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

type SignatureBuilder struct {
	ApiKey       string
	SharedSecret string
	Method       string
	Token        string
	SessionKey   string
	track        string
	artist       string
	album        string
	albumArtist  string
}

func (sb *SignatureBuilder) SetTrack(track string) {
	sb.track = track
}

func (sb *SignatureBuilder) SetArtist(artist string) {
	sb.artist = artist
}

func (sb *SignatureBuilder) SetAlbum(album string) {
	sb.album = album
}

func (sb *SignatureBuilder) SetAlbumArtist(albumArtist string) {
	sb.albumArtist = albumArtist
}

func (sb *SignatureBuilder) SignatureBase(timestamp string, isScrobble bool) string {
	builder := strings.Builder{}

	if isScrobble && sb.album != "" && sb.albumArtist != ""{
		builder.WriteString("album")
		builder.WriteString(sb.album)

		builder.WriteString("albumArtist")
		builder.WriteString(sb.albumArtist)
	}

	builder.WriteString("api_key")
	builder.WriteString(sb.ApiKey)

	if isScrobble{
		builder.WriteString("artist")
		builder.WriteString(sb.artist)
	}

	builder.WriteString("method")
	builder.WriteString(sb.Method)

	if sb.SessionKey != "" {
		builder.WriteString("sk")
		builder.WriteString(sb.SessionKey)
	}
	
	if isScrobble{
		builder.WriteString("timestamp")
		builder.WriteString(timestamp)

		builder.WriteString("track")
		builder.WriteString(sb.track)
	}

	if sb.Token != "" {
		builder.WriteString("token")
		builder.WriteString(sb.Token)
	}

	builder.WriteString(sb.SharedSecret)

	return builder.String()
}

func (sb *SignatureBuilder) Signature(timestamp string, isScrobble bool) string {
	hasher := md5.New()
	sigBase := sb.SignatureBase(timestamp, isScrobble)
	println("sig_base", sigBase)
	io.WriteString(hasher, sigBase)
	md5String := hex.EncodeToString(hasher.Sum(nil))

	return md5String
}
