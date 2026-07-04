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
}

func (sb *SignatureBuilder) SetTrack(track string) {
	sb.track = track
}

func (sb *SignatureBuilder) SetArtist(artist string) {
	sb.artist = artist
}

func (sb *SignatureBuilder) SignatureBase(timestamp string) string {
	builder := strings.Builder{}

	builder.WriteString("api_key")
	builder.WriteString(sb.ApiKey)

	builder.WriteString("artist")
	builder.WriteString(sb.artist)

	builder.WriteString("method")
	builder.WriteString(sb.Method)

	if sb.SessionKey != "" {
		builder.WriteString("sk")
		builder.WriteString(sb.SessionKey)
	}

	builder.WriteString("timestamp")
	builder.WriteString(timestamp)
	
	builder.WriteString("track")
	builder.WriteString(sb.track)

	if sb.Token != "" {
		builder.WriteString("token")
		builder.WriteString(sb.Token)
	}

	builder.WriteString(sb.SharedSecret)

	return builder.String()
}

func (sb *SignatureBuilder) Signature(timestamp string) string {
	hasher := md5.New()
	sigBase := sb.SignatureBase(timestamp)
	println("sig_base", sigBase)
	io.WriteString(hasher, sigBase)
	md5String := hex.EncodeToString(hasher.Sum(nil))

	return md5String
}
