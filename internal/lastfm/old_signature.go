package lastfm

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
	"strings"
)

type OldSignatureBuilder struct {
	ApiKey       string
	SharedSecret string
	Method       string
	Token        string
	SessionKey   string
	tracks       []string
	artists      []string
}

func (sb *OldSignatureBuilder) AddTrack(track string, artist string) {
	sb.tracks = append(sb.tracks, track)
	sb.artists = append(sb.artists, artist)
}

func (sb *OldSignatureBuilder) SignatureBase(timestamp string) string {
	builder := strings.Builder{}

	builder.WriteString("api_key")
	builder.WriteString(sb.ApiKey)

	for index, value := range sb.artists {
		builder.WriteString("artist[" + strconv.Itoa(index) + "]")
		builder.WriteString(value)
	}

	builder.WriteString("method")
	builder.WriteString(sb.Method)

	if sb.SessionKey != "" {
		builder.WriteString("sk")
		builder.WriteString(sb.SessionKey)
	}

	for index := range sb.tracks {
		builder.WriteString("timestamp[" + strconv.Itoa(index) + "]")
		builder.WriteString(timestamp)
	}

	for index, value := range sb.tracks {
		builder.WriteString("track[" + strconv.Itoa(index) + "]")
		builder.WriteString(value)
	}

	if sb.Token != "" {
		builder.WriteString("token")
		builder.WriteString(sb.Token)
	}

	builder.WriteString(sb.SharedSecret)

	return builder.String()
}

func (sb *OldSignatureBuilder) Signature(timestamp string) string {
	hasher := md5.New()
	io.WriteString(hasher, sb.SignatureBase(timestamp))
	md5String := hex.EncodeToString(hasher.Sum(nil))

	return md5String
}
