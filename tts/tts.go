package tts

import (
	"errors"
	"io"
	"os"
	"path"
	"strings"

	"github.com/matthew-balzan/dca"
	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
)

type TTS interface {
	Name() string
	Voices() []string
	Run(request TTSRequest) (dca.OpusReader, error)
}

type TTSRequest struct {
	Text  string
	Voice string
}

var tts []TTS
var db *bbolt.DB

func Start(database *bbolt.DB) *[]TTS {
	db = database
	tts = nil
	tts = append(tts, new(GTranslateTTS))
	tts = append(tts, new(DectalkTTS))
	return &tts
}

func FindTTS(search string) (TTS, error) {
	for _, method := range tts {
		if strings.ToLower(search) == method.Name() {
			return method, nil
		}
	}

	return nil, errors.New("TTS method not found")
}

// EncodeToOpus uses FFmpeg to encode the input into Opus (defer .Cleanup())
func EncodeToOpus(r io.Reader) (*dca.EncodeSession, error) {
	options := dca.StdEncodeOptions
	options.Application = dca.AudioApplicationVoip
	encode, err := dca.EncodeMem(r, options)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode speech")
		return nil, err
	}
	log.Trace().Interface("stats", encode.Stats()).Msg("Encoding with FFmpeg")
	return encode, nil
}

// EncodeFileToOpus uses FFmpeg to encode the input into Opus (defer .Cleanup())
func EncodeFileToOpus(r string) (*dca.EncodeSession, error) {
	options := dca.StdEncodeOptions
	options.Application = dca.AudioApplicationVoip
	encode, err := dca.EncodeFile(r, options)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode speech")
		return nil, err
	}
	return encode, nil
}

// getTempDir finds the temp files folder. Creates a new one if it does not exist.
func getTempDir() (string, error) {
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

// DeleteTempDir deletes the temp files folder.
func DeleteTempDir() error {
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
