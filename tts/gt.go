package tts

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"unicode/utf8"

	"github.com/rs/zerolog/log"
)

// NewGTSpeech uses the Google Translate API to generate a speech (mp3)
func NewGTSpeech(text string) (io.ReadCloser, error) {
	// Limit text size to 200
	if utf8.RuneCountInString(text) > 200 {
		return nil, errors.New("text is too long")
	}

	host, err := url.Parse("https://translate.google.com/translate_tts")
	if err != nil {
		return nil, err
	}

	q := host.Query()
	q.Add("tl", "en-US")
	q.Add("q", text)
	q.Add("client", "tw-ob")
	host.RawQuery = q.Encode()

	log.Trace().Str("url", host.String()).Msg("GTranslate request")
	res, err := http.Get(host.String())
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		return res.Body, nil
	} else {
		return nil, errors.New("status: " + res.Status)
	}
}
