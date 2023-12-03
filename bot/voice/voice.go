package voice

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/matthew-balzan/dca"
	"github.com/rs/zerolog/log"
	"io"
	"sync"
)

// TODO: convert into a map?
var voiceConnections []Connection
var vcMutex = &sync.Mutex{}

type Connection struct {
	conn  *discordgo.VoiceConnection
	queue chan dca.OpusReader
	stop  chan struct{}
}

func JoinVoice(s *discordgo.Session, gID string, vcID string) (*Connection, error) {
	vcMutex.Lock()
	defer vcMutex.Unlock()

	for _, connection := range voiceConnections {
		if connection.conn.GuildID == gID {
			if connection.conn.ChannelID == vcID {
				// Already connected to the requested channel
				return &connection, errors.New("inCurrentChannel")
			}
			// In the guild, but in a different channel
			err := connection.conn.ChangeChannel(vcID, false, true)
			if err != nil {
				return nil, err
			}
			return &connection, errors.New("inAnotherChannel")
		}
	}

	log.Trace().Str("gID", gID).Str("vcID", vcID).Msg("Joining voice channel")
	vc, err := s.ChannelVoiceJoin(gID, vcID, false, true)
	if err != nil {
		return nil, err
	}

	queue := make(chan dca.OpusReader, 10)
	stop := make(chan struct{})
	vConn := Connection{conn: vc, queue: queue, stop: stop}
	go vConn.handleQueue()
	voiceConnections = append(voiceConnections, vConn)
	return &voiceConnections[len(voiceConnections)-1], nil
}

func (vc *Connection) handleQueue() {
	for {
		var play dca.OpusReader
		select {
		case play = <-vc.queue:
			log.Trace().Str("gID", vc.conn.GuildID).Str("vcID", vc.conn.ChannelID).Msg("Playing audio")
			done := make(chan error)
			dca.NewStream(play, vc.conn, done)
			err := <-done

			if err != nil && err != io.EOF {
				log.Error().Err(err).Msg("Failed to play Opus data")
			}

			session, ok := play.(*dca.EncodeSession)
			if ok {
				session.Cleanup()
			}
		case <-vc.stop:
			break
		}
	}
}

func (vc *Connection) AddToQueue(r dca.OpusReader) {
	vc.queue <- r
}

func FindByGuild(gID string) *Connection {
	vcMutex.Lock()
	defer vcMutex.Unlock()

	for _, connection := range voiceConnections {
		if connection.conn.GuildID == gID {
			return &connection
		}
	}

	return nil
}

func (vc *Connection) Disconnect() {
	log.Trace().Str("gID", vc.conn.GuildID).Str("vcID", vc.conn.ChannelID).Msg("Disconnecting with VC")
	vc.stop <- struct{}{}
	err := vc.conn.Disconnect()
	if err != nil {
		log.Error().Err(err).Str("gID", vc.conn.GuildID).Str("vcID", vc.conn.ChannelID).Msg("Failed to disconnect")
	}

	vcMutex.Lock()
	defer vcMutex.Unlock()
	for i, connection := range voiceConnections {
		if connection.conn.ChannelID == vc.conn.ChannelID {
			voiceConnections[i] = voiceConnections[len(voiceConnections)-1]
			voiceConnections = voiceConnections[:len(voiceConnections)-1]
		}
	}
}

func DisconnectAll() {
	vcMutex.Lock()
	defer vcMutex.Unlock()
	for _, connection := range voiceConnections {
		connection.stop <- struct{}{}
		err := connection.conn.Disconnect()
		if err != nil {
			log.Error().Err(err).Str("gID", connection.conn.GuildID).Str("vcID", connection.conn.ChannelID).Msg("Failed to disconnect")
		}
	}
	voiceConnections = nil
}

// DetectVoiceChannel finds the voice channel the author is in
func DetectVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate) (string, error) {
	g, err := s.State.Guild(m.GuildID)
	if err != nil {
		return "", err
	}

	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			return vs.ChannelID, nil
		}
	}

	return "", errors.New("user voice channel not found")
}
