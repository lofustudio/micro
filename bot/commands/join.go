package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lofustudio/VEGA/bot/voice"
	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
)

type JoinCommand struct{}

func (p JoinCommand) Name() string {
	return "join"
}

func (p JoinCommand) Description() string {
	return "Join the voice channel you are in!"
}

func (p JoinCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, _ *bbolt.DB) {
	authorVCID, err := voice.DetectVoiceChannel(s, m)
	if err != nil {
		log.Error().Err(err).Msg("Could not find the author's vc")
		_, err := s.ChannelMessageSendReply(m.ChannelID, "Couldn't find the voice channel!", m.Reference())
		if err != nil {
			log.Error().Err(err).Msg("Failed to send reply")
			return
		}
		return
	}

	_, err = voice.JoinVoice(s, m.GuildID, authorVCID)
	if err != nil {
		log.Error().Err(err).Msg("Could not join the author's vc")
		_, err := s.ChannelMessageSendReply(m.ChannelID, "Couldn't join the voice channel!", m.Reference())
		if err != nil {
			log.Error().Err(err).Msg("Failed to send reply")
			return
		}
		return
	}
}

var _ Command = (*JoinCommand)(nil)
