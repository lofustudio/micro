package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
)

type PingCommand struct{}

func (p PingCommand) Name() string {
	return "ping"
}

func (p PingCommand) Description() string {
	return "pong"
}

func (p PingCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, _ *bbolt.DB) {
	_, err := s.ChannelMessageSendReply(m.ChannelID, "pong!", m.Reference())
	if err != nil {
		log.Error().Err(err).Msg("Error while executing PingCommand")
	}
}

var _ Command = (*PingCommand)(nil)
