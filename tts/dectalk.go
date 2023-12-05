package tts

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// NewDecSpeech uses DECtalk to generate a speech (wav)
func NewDecSpeech(text string) (io.ReadCloser, error) {
	tempdir, err := detectTempDir()
	if err != nil {
		log.Error().Err(err).Msg("Failed to find temporary directory")
		return nil, err
	}

	file := path.Join(tempdir, fmt.Sprint("temp_", rand.Intn(1000), ".wav"))
	log.Trace().Str("file", file).Msg("Temporary wav file")

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("./say.exe", "-w", file, fmt.Sprint("[:phoneme on]", text))
		cmd.Dir = viper.GetString("dectalk")
	case "linux":
		cmd = exec.Command("./say", "-fo", file, "-a", fmt.Sprint('"', "[:phoneme on]", text, '"'))
		cmd.Dir = viper.GetString("dectalk")
	default:
		return nil, errors.ErrUnsupported
	}

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
