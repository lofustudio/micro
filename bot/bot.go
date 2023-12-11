package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lofustudio/VEGA/bot/commands"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.etcd.io/bbolt"
)

func Start(db *bbolt.DB) func() error {
	// Create Discord session
	dg, err := discordgo.New("Bot " + viper.GetString("token"))
	if err != nil {
		log.Panic().Err(err).Msg("Error creating Discord session")
	}

	// Create command handler
	_ = commands.Start(db)

	// Set intents and handlers
	dg.Identify.Intents = discordgo.Intent(viper.GetInt("intents"))
	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)
	dg.AddHandler(commands.MessageCommandHandler)

	// Open the Discord Session
	err = dg.Open()
	if err != nil {
		log.Panic().Err(err).Msg("Error opening Discord session")
	}

	return dg.Close
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Info().Msg("The bot is ready!")
}
