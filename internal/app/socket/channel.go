package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/kyleu/rituals.dev/internal/app/util"
)

func (s *Service) Register(userID uuid.UUID, c *websocket.Conn) (uuid.UUID, error) {
	conn := connection{
		ID:       util.UUID(),
		UserID:   userID,
		Channels: make([]uuid.UUID, 0),
		socket:   c,
	}
	s.connections[conn.ID] = &conn
	return conn.ID, nil
}

func (s *Service) Join(connID uuid.UUID, channelID uuid.UUID) error {
	conn, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("invalid connection [" + connID.String() + "]"))
	}
	if !contains(conn.Channels, channelID) {
		conn.Channels = append(conn.Channels, channelID)
	}

	curr, ok := s.channels[channelID]
	if !ok {
		curr = make([]uuid.UUID, 0)
	}
	if !contains(curr, connID) {
		s.channels[channelID] = append(curr, connID)
	}
	return nil
}

func (s *Service) Disconnect(connID uuid.UUID) (int, error) {
	conn, ok := s.connections[connID]
	if !ok {
		return 0, errors.WithStack(errors.New("invalid connection [" + connID.String() + "]"))
	}
	size := len(conn.Channels)
	for _, c := range conn.Channels {
		err := s.Leave(connID, c)
		if err != nil {
			return size, errors.WithStack(errors.Wrap(err, "error leaving channel [" + c.String() + "]"))
		}
	}
	delete(s.connections, connID)
	return size, nil
}

func (s *Service) Leave(connID uuid.UUID, channelID uuid.UUID) error {
	conn, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("invalid connection [" + connID.String() + "]"))
	}
	if contains(conn.Channels, channelID) {
		chans := make([]uuid.UUID, 0)
		for _, c := range conn.Channels {
			if c != channelID {
				chans = append(chans, c)
			}
		}
		conn.Channels = chans
	}

	curr, ok := s.channels[channelID]
	if !ok {
		curr = make([]uuid.UUID, 0)
	}
	filtered := make([]uuid.UUID, 0)
	for _, i := range curr {
		if i != connID {
			filtered = append(filtered, i)
		}
	}
	s.channels[channelID] = filtered
	return nil
}

func contains(s []uuid.UUID, e uuid.UUID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
