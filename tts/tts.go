package tts

import (
	"io"
	"os"
	"path"

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

func detectTempDir() (string, error) {
	tempdir := path.Join(os.TempDir(), "VEGA")
	_, err := os.Stat(tempdir)
	if os.IsNotExist(err) {
		log.Trace().Str("tempdir", tempdir).Msg("Creating temp directory")
		err := os.Mkdir(tempdir, 0666)
		if err != nil {
			return "", err
		}
	}

	return tempdir, nil
}

func ClearTempDir() error {
	tempdir := path.Join(os.TempDir(), "VEGA")
	_, err := os.Stat(tempdir)
	if err == nil {
		err = os.RemoveAll(tempdir)
		if err != nil {
			log.Error().Err(err).Msg("Error deleting config file")
			return err
		}
	}
	return nil
}
