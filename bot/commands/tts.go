package commands

import (
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lofustudio/VEGA/bot/voice"
	"github.com/lofustudio/VEGA/tts"
	"github.com/matthew-balzan/dca"
	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
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
	method := words[1]
	lang := words[2]

	find, err := tts.FindTTS(method)
	if err != nil {
		log.Error().Err(err).Msg("Failed to find TTS method")
		return
	}

	vc := voice.FindByGuild(m.GuildID)
	if vc == nil {
		log.Error().Msg("Failed to find VC")
		return
	}

	var encoded dca.OpusReader
	if slices.Contains(find.Voices(), lang) {
		encoded, err = find.Run(tts.TTSRequest{Text: strings.Join(words[3:], " "), Voice: lang})
	} else {
		encoded, err = find.Run(tts.TTSRequest{Text: strings.Join(words[2:], " ")})
	}
	if err != nil {
		log.Error().Err(err).Msg("Failed to play TTS")
		return
	}

	vc.AddToQueue(encoded)
}

var _ Command = (*TtsCommand)(nil)
