package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lofustudio/VEGA/bot/voice"
	"github.com/lofustudio/VEGA/tts"
	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
	"strings"
)

type TtsCommand struct{}

func (p TtsCommand) Name() string {
	return "tts"
}

func (p TtsCommand) Description() string {
	return "Plays the message in TTS!"
}

func (p TtsCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, _ *bbolt.DB) {
	words := strings.Split(m.Content, " ")
	content := strings.Join(words[1:], " ")

	vc := voice.FindByGuild(m.GuildID)
	if vc == nil {
		log.Error().Msg("Failed to find VC")
		return
	}

	speech, err := tts.NewGTSpeech(content)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate speech")
		return
	}

	encoded, err := tts.EncodeToOpus(speech)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode speech")
		return
	}

	vc.AddToQueue(encoded)
}

var _ Command = (*TtsCommand)(nil)
