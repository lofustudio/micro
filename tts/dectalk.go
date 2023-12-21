package tts

import (
	"errors"
	"fmt"
	"math/rand"
	"os/exec"
	"path"
	"runtime"

	"github.com/matthew-balzan/dca"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type DectalkTTS struct{}

// Name implements TTS.
func (*DectalkTTS) Name() string {
	return "dec"
}

// Voices implements TTS.
func (*DectalkTTS) Voices() []string {
	return nil
}

// Run implements TTS.
// TODO: delete the wav file
func (*DectalkTTS) Run(request TTSRequest) (dca.OpusReader, error) {
	if !viper.IsSet("dectalk") {
		return nil, errors.New("DECtalk not found")
	}

	tempdir, err := getTempDir()
	if err != nil {
		log.Error().Err(err).Msg("Failed to find temporary directory")
		return nil, err
	}

	file := path.Join(tempdir, fmt.Sprint("temp_", rand.Intn(1000), ".wav"))
	log.Trace().Str("file", file).Msg("Temporary wav file")

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("./say.exe", "-w", file, fmt.Sprint("[:phoneme on]", request.Text))
		cmd.Dir = viper.GetString("dectalk")
	case "linux":
		cmd = exec.Command("./say", "-fo", file, "-a", fmt.Sprint("[:phoneme on]", request.Text))
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

	encoded, err := EncodeFileToOpus(file)
	if err != nil {
		return nil, err
	}

	return encoded, nil
}

var _ TTS = (*DectalkTTS)(nil)
