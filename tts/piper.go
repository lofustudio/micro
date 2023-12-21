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

type PiperTTS struct{}

// Name implements TTS.
func (*PiperTTS) Name() string {
	return "piper"
}

// Voices implements TTS.
func (*PiperTTS) Voices() []string {
	return []string{"glados", "ar_JO-kareem-medium", "ca_ES-upc_ona-medium", "ca_ES-upc_ona-x_low", "ca_ES-upc_pau-x_low", "cs_CZ-jirka-low", "cs_CZ-jirka-medium", "da_DK-talesyntese-medium", "de_DE-eva_k-x_low", "de_DE-karlsson-low", "de_DE-kerstin-low", "de_DE-pavoque-low", "de_DE-ramona-low", "de_DE-thorsten-high", "de_DE-thorsten-low", "de_DE-thorsten-medium", "de_DE-thorsten_emotional-medium", "el_GR-rapunzelina-low", "en_GB-alan-low", "en_GB-alan-medium", "en_GB-alba-medium", "en_GB-aru-medium", "en_GB-jenny_dioco-medium", "en_GB-northern_english_male-medium", "en_GB-semaine-medium", "en_GB-southern_english_female-low", "en_GB-vctk-medium", "en_US-amy-low", "en_US-amy-medium", "en_US-arctic-medium", "en_US-danny-low", "en_US-hfc_male-medium", "en_US-joe-medium", "en_US-kathleen-low", "en_US-kusal-medium", "en_US-l2arctic-medium", "en_US-lessac-high", "en_US-lessac-low", "en_US-lessac-medium", "en_US-libritts-high", "en_US-libritts_r-medium", "en_US-ryan-high", "en_US-ryan-low", "en_US-ryan-medium", "es_ES-carlfm-x_low", "es_ES-davefx-medium", "es_ES-mls_10246-low", "es_ES-mls_9972-low", "es_ES-sharvard-medium", "es_MX-ald-medium", "fi_FI-harri-low", "fi_FI-harri-medium", "fr_FR-gilles-low", "fr_FR-mls_1840-low", "fr_FR-siwis-low", "fr_FR-siwis-medium", "fr_FR-upmc-medium", "hu_HU-anna-medium", "hu_HU-berta-medium", "hu_HU-imre-medium", "is_IS-bui-medium", "is_IS-salka-medium", "is_IS-steinn-medium", "is_IS-ugla-medium", "it_IT-riccardo-x_low", "ka_GE-natia-medium", "kk_KZ-iseke-x_low", "kk_KZ-issai-high", "kk_KZ-raya-x_low", "lb_LU-marylux-medium", "ne_NP-google-medium", "ne_NP-google-x_low", "nl_BE-nathalie-medium", "nl_BE-nathalie-x_low", "nl_BE-rdh-medium", "nl_BE-rdh-x_low", "nl_NL-mls_5809-low", "nl_NL-mls_7432-low", "no_NO-talesyntese-medium", "pl_PL-darkman-medium", "pl_PL-gosia-medium", "pl_PL-mc_speech-medium", "pl_PL-mls_6892-low", "pt_BR-edresson-low", "pt_BR-faber-medium", "pt_PT-tugão-medium", "ro_RO-mihai-medium", "ru_RU-denis-medium", "ru_RU-dmitri-medium", "ru_RU-irina-medium", "ru_RU-ruslan-medium", "sk_SK-lili-medium", "sr_RS-serbski_institut-medium", "sv_SE-nst-medium", "sw_CD-lanfrica-medium", "tr_TR-dfki-medium", "tr_TR-fahrettin-medium", "uk_UA-lada-x_low", "uk_UA-ukrainian_tts-medium", "vi_VN-25hours_single-low", "vi_VN-vais1000-medium", "vi_VN-vivos-x_low", "zh_CN-huayan-medium", "zh_CN-huayan-x_low"}
}

// Run implements TTS.
// TODO: delete the wav file
func (*PiperTTS) Run(request TTSRequest) (dca.OpusReader, error) {
	if !viper.IsSet("piper") {
		return nil, errors.New("piper not found")
	}

	if request.Voice == "" {
		request.Voice = "en_US-lessac-medium"
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
		// Windows build cannot download models on its own
		cmd = exec.Command("cmd", "/C", fmt.Sprint("echo ", request.Text, " | ", "piper.exe ", "--model ", request.Voice+".onnx ", "--output_file ", file))
		cmd.Dir = viper.GetString("piper")
	case "linux":
		cmd = exec.Command("/usr/bin/echo", "'"+request.Text+"'", "|", "./piper", "--model", request.Voice, "--output_file", file, "--data-dir", "./data", "--download-dir", "./download")
		cmd.Dir = viper.GetString("piper")
	default:
		return nil, errors.ErrUnsupported
	}

	log.Trace().Strs("args", cmd.Args).Msg("Running piper")
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

var _ TTS = (*PiperTTS)(nil)
