package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.etcd.io/bbolt"
	"strings"
)

type Command interface {
	Name() string
	Description() string
	Run(s *discordgo.Session, m *discordgo.MessageCreate, db *bbolt.DB)
}

var commands []Command
var db *bbolt.DB

func Start(database *bbolt.DB) []Command {
	db = database
	commands = nil
	commands = append(commands, new(PingCommand))
	commands = append(commands, new(JoinCommand))
	commands = append(commands, new(TtsCommand))
	commands = append(commands, new(DecCommand))

	return commands
}

func MessageCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID != "1180707123710144552" {
		return
	}

	words := strings.Split(m.Content, " ")

	if !strings.HasPrefix(words[0], viper.GetString("prefix")) {
		return
	}

	cmd := strings.ToLower(words[0][len(viper.GetString("prefix")):])

	for _, command := range commands {
		if cmd == command.Name() {
			log.Trace().Str("command", cmd).Msg("Running command")
			command.Run(s, m, db)
		}
	}
}
