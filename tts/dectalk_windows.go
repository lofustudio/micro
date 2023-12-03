//go:build windows

package tts

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path"
)

// NewDecSpeech uses DECtalk to generate a speech (wav)
func NewDecSpeech(text string) (io.ReadCloser, error) {
	tempdir := path.Join(os.TempDir(), "VEGA")
	_, err := os.Stat(tempdir)
	if os.IsNotExist(err) {
		log.Trace().Str("tempdir", tempdir).Msg("Creating temp directory")
		err := os.Mkdir(tempdir, 0666)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create temp directory")
			return nil, err
		}
	}

	file := path.Join(tempdir, fmt.Sprint("temp_", rand.Intn(1000), ".wav"))
	log.Trace().Str("file", file).Msg("Temporary wav file")

	cmd := exec.Command("./say.exe", "-w", file, fmt.Sprint("[:phoneme on]", text))
	cmd.Dir = viper.GetString("dectalk")
	log.Trace().Strs("args", cmd.Args).Msg("Running dectalk")
	err = cmd.Run()
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate .wav file")
		return nil, err
	}

	wavFile, err := os.Open(file)
	if err != nil {
		log.Error().Err(err).Msg("Failed to open .wav file")
		return nil, err
	}

	return wavFile, nil
}
