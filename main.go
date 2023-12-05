package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lofustudio/VEGA/bot"
	"github.com/lofustudio/VEGA/dash"
	"github.com/lofustudio/VEGA/tts"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.etcd.io/bbolt"
)

func main() {
	// Zerolog setup
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Viper setup
	config()

	defer tts.ClearTempDir()

	// Create or open the database
	db, err := bbolt.Open(viper.GetString("database"), 0600, nil)
	if err != nil {
		log.Panic().Err(err).Msg("Error creating or opening database")
	}
	defer db.Close()

	// Start the bot
	botClose := bot.Start(db)
	defer botClose()

	// Start the dash
	if viper.GetBool("dash") {
		dashClose := dash.Start()
		defer dashClose()
	}

	// Wait until CTRL-C
	log.Info().Msg("Bot is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func config() {
	// Set defaults
	viper.SetDefault("dash", false)
	viper.SetDefault("prefix", "!")
	viper.SetDefault("database", "vega.db")
	viper.SetDefault("intents", 33409)
	// Read from config
	viper.SetConfigName("prod")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic().Err(err).Msg("Error reading config file")
	}

	// Panic if token is not set
	if !viper.IsSet("token") {
		log.Panic().Msg("Discord bot token not found")
	}
}
