package tts

import (
	"errors"
	"net/http"
	"net/url"
	"unicode/utf8"

	"github.com/matthew-balzan/dca"
	"github.com/rs/zerolog/log"
)

type GTranslateTTS struct{}

// Name implements TTS.
func (*GTranslateTTS) Name() string {
	return "gt"
}

// Voices implements TTS.
func (*GTranslateTTS) Voices() []string {
	return []string{"af-ZA", "ar-XA", "eu-ES", "bn-IN", "bg-BG", "ca-ES", "yue-HK", "cs-CZ", "da-DK", "nl-BE", "nl-NL", "en-AU", "en-IN", "en-GB", "en-US", "fil-PH", "fi-FI", "fr-CA", "fr-FR", "gl-ES", "de-DE", "el-GR", "gu-IN", "he-IL", "hi-IN", "hu-HU", "is-IS", "id-ID", "it-IT", "ja-JP", "kn-IN", "ko-KR", "lv-LV", "lt-LT", "ms-MY", "ml-IN", "cmn-CN", "cmn-TW", "mr-IN", "nb-NO", "pl-PL", "pt-BR", "pt-PT", "pa-IN", "ro-RO", "ru-RU", "sr-RS", "sk-SK", "es-ES", "es-US", "sv-SE", "ta-IN", "te-IN", "th-TH", "tr-TR", "uk-UA", "vi-VN"}
}

// Run implements TTS.
func (*GTranslateTTS) Run(request TTSRequest) (dca.OpusReader, error) {
	// Limit text size to 200
	if utf8.RuneCountInString(request.Text) > 200 {
		return nil, errors.New("text is too long")
	}

	host, err := url.Parse("https://translate.google.com/translate_tts")
	if err != nil {
		return nil, err
	}

	q := host.Query()
	if request.Voice != "" {
		q.Add("tl", request.Voice)
	} else {
		q.Add("tl", "en-US")
	}
	q.Add("q", request.Text)
	q.Add("client", "tw-ob")
	host.RawQuery = q.Encode()

	log.Trace().Str("url", host.String()).Msg("GTranslate request")
	res, err := http.Get(host.String())
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("status: " + res.Status)
	}

	encoded, err := EncodeToOpus(res.Body)
	if err != nil {
		return nil, err
	}

	return encoded, nil
}

var _ TTS = (*GTranslateTTS)(nil)
