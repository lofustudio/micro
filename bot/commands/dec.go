package commands

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lofustudio/VEGA/bot/voice"
	"github.com/lofustudio/VEGA/tts"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.etcd.io/bbolt"
)

type DecCommand struct{}

func (p DecCommand) Name() string {
	return "dec"
}

func (p DecCommand) Description() string {
	return "Plays the message in TTS!"
}

func (p DecCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, _ *bbolt.DB) {
	if !viper.IsSet("dectalk") {
		log.Trace().Msg("DECtalk not detected")
		return
	}

	words := strings.Split(m.Content, " ")
	content := strings.Join(words[1:], " ")

	vc := voice.FindByGuild(m.GuildID)
	if vc == nil {
		log.Error().Msg("Failed to find VC")
		return
	}

	speech, err := tts.NewDecSpeech(content)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate speech")
		return
	}

	encoded, err := tts.EncodeToOpus(speech)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode speech")
		return
	}

	// Delete the tempoary .wav file
	file, ok := speech.(*os.File)
	if ok {
		err = os.Remove(file.Name())
		if err != nil {
			log.Error().Err(err).Msg("Failed to delete temporary .wav file")
			return
		}
	}

	vc.AddToQueue(encoded)
}

var _ Command = (*DecCommand)(nil)
