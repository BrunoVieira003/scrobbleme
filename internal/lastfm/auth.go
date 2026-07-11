package lastfm

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"scrobbleme/internal"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/pkg/browser"
)

type GetTokenResponse struct {
	Token string `json:"token"`
}

type AuthorizationResponse struct {
	Session struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	} `json:"session"`
}

type LastFmErrorResponse struct {
	Message string `json:"message"`
	Error   int    `json:"error"`
}

func AuthenticateLastFM() AuthorizationResponse {
	token := getToken()

	authorizeUrl, err := url.Parse("http://www.last.fm/api/auth")
	if err != nil {
		log.Fatal(err)
	}

	params := authorizeUrl.Query()
	params.Add("api_key", internal.LASTFM_KEY)
	params.Add("token", token)

	authorizeUrl.RawQuery = params.Encode()

	err = browser.OpenURL(authorizeUrl.String())
	if err != nil {
		log.Fatal(err)
	}

	authorization, err := CheckAuthorization(token, 5)
	if err != nil {
		beeep.Notify("Unable to check auth.api_key", internal.LASTFM_KEY, "")
		log.Fatal(err.Error())
	}

	return authorization

}

func CheckAuthorization(token string, retries int8) (AuthorizationResponse, error) {
	var authResponse AuthorizationResponse
	var errResponse LastFmErrorResponse

	for range retries {
		url, err := url.Parse("https://ws.audioscrobbler.com/2.0")
		if err != nil {
			log.Fatal("Unable to parse URL")
		}

		params := url.Query()
		params.Add("api_key", internal.LASTFM_KEY)
		params.Add("api_sig", GenerateSigForSession(token))
		params.Add("format", "json")
		params.Add("method", "auth.getSession")
		params.Add("token", token)

		url.RawQuery = params.Encode()

		resp, err := http.Get(url.String())
		if err != nil {
			beeep.Notify("Authentication failed", err.Error(), "./assets/icon.ico")
			log.Fatal("Unable to make request")
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			time.Sleep(5 * time.Second)
			err = json.NewDecoder(resp.Body).Decode(&errResponse)
			if err != nil {
				log.Fatal("Unable to parse error", err.Error())
			}

			if errResponse.Error == 14 {
				continue
			} else {
				println("error_code", errResponse.Error, errResponse.Message)
				log.Fatal(err.Error())
			}
		}

		err = json.NewDecoder(resp.Body).Decode(&authResponse)
		if err != nil {
			log.Fatal(err.Error())
		}

		return authResponse, nil
	}

	return authResponse, errors.New("User not authorized")
}

func GenerateSigForSession(token string) string {
	sigBuilder := SignatureBuilder{
		ApiKey:       internal.LASTFM_KEY,
		SharedSecret: internal.LASTFM_SECRET,
		Method:       "auth.getSession",
		Token:        token,
	}

	return sigBuilder.Signature("", false)
}

func GenerateSigForScrobble(sk string, timestamp string, track string, artist string, album string, albumArtist string) string {
	sigBuilder := SignatureBuilder{
		Method:       "track.scrobble",
		ApiKey:       internal.LASTFM_KEY,
		SharedSecret: internal.LASTFM_SECRET,
		SessionKey:   sk,
	}

	sigBuilder.SetTrack(track)
	sigBuilder.SetArtist(artist)

	sigBuilder.SetAlbum(album)
	sigBuilder.SetAlbumArtist(albumArtist)

	return sigBuilder.Signature(timestamp, true)
}

func getToken() string {
	authenticateUrl, err := url.Parse("https://ws.audioscrobbler.com/2.0")
	if err != nil {
		log.Fatal(err)
	}

	params := authenticateUrl.Query()
	params.Add("method", "auth.getToken")
	params.Add("api_key", internal.LASTFM_KEY)
	params.Add("format", "json")

	authenticateUrl.RawQuery = params.Encode()

	resp, err := http.Get(authenticateUrl.String())
	if err != nil {
		log.Fatal("Error during auth process. Contact the support team")
	}
	defer resp.Body.Close()

	var tokenResponse GetTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		log.Fatal("Error during auth process. Contact the support team")
	}

	return tokenResponse.Token
}
