package tts

import (
	"io"

	"github.com/matthew-balzan/dca"
	"github.com/rs/zerolog/log"
)

// EncodeToOpus uses FFmpeg to encode the input into Opus (defer .Cleanup())
func EncodeToOpus(r io.Reader) (*dca.EncodeSession, error) {
	options := dca.StdEncodeOptions
	options.Application = dca.AudioApplicationVoip
	encode, err := dca.EncodeMem(r, options)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode speech")
		return nil, err
	}
	return encode, nil
}
