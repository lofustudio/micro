package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() func() error {
	// Create a Discord session
	dg, err := discordgo.New("Bot " + viper.GetString("token"))
	if err != nil {
		log.Panic().Err(err).Msg("Error creating Discord session")
	}

	// Set intents and handlers
	dg.Identify.Intents = discordgo.Intent(viper.GetInt("intents"))
	dg.AddHandler(ready)

	// Open the Discord Session
	err = dg.Open()
	if err != nil {
		log.Panic().Err(err).Msg("Error opening Discord session")
	}

	return dg.Close
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Info().Msg("Ready event received")
}
