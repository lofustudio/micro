package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lofustudio/VEGA/bot/voice"
	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
)

type LeaveCommand struct{}

func (p LeaveCommand) Name() string {
	return "leave"
}

func (p LeaveCommand) Description() string {
	return "Leaves the voice channel the bot is in!"
}

func (p LeaveCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, _ *bbolt.DB) {
	conn := voice.FindByGuild(m.GuildID)

	if conn == nil {
		log.Error().Msg("Failed to find the connection")
		return
	}

	conn.Disconnect()
}

var _ Command = (*LeaveCommand)(nil)
