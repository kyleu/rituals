package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"sync"
)

type channel struct {
	Svc string
	ID  uuid.UUID
}

func (c channel) String() string {
	return fmt.Sprintf("%s:%s", c.Svc, c.ID)
}

func (s *Service) Register(userID uuid.UUID, c *websocket.Conn) (uuid.UUID, error) {
	conn := connection{
		ID:      util.UUID(),
		UserID:  userID,
		Svc:     "",
		ModelID: nil,
		Channel: nil,
		socket:  c,
		mu:      sync.Mutex{},
	}
	s.connectionsMu.Lock()
	defer s.connectionsMu.Unlock()

	s.connections[conn.ID] = &conn
	return conn.ID, nil
}

func (s *Service) Join(connID uuid.UUID, ch channel) error {
	conn, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("invalid connection [" + connID.String() + "]"))
	}
	if conn.Channel != &ch {
		conn.Channel = &ch
	}

	s.channelsMu.Lock()
	defer s.channelsMu.Unlock()

	curr, ok := s.channels[ch]
	if !ok {
		curr = make([]uuid.UUID, 0)
	}
	if !contains(curr, connID) {
		s.channels[ch] = append(curr, connID)
	}
	return nil
}

func (s *Service) Disconnect(connID uuid.UUID) (bool, error) {
	conn, ok := s.connections[connID]
	if !ok {
		return false, errors.WithStack(errors.New("invalid connection [" + connID.String() + "]"))
	}
	left := false

	if conn.Channel != nil {
		left = true
		err := s.Leave(connID, *conn.Channel)
		if err != nil {
			return left, errors.WithStack(errors.Wrap(err, "error leaving channel ["+conn.Channel.String()+"]"))
		}
	}

	s.connectionsMu.Lock()
	defer s.connectionsMu.Unlock()

	delete(s.connections, connID)
	return left, nil
}

func (s *Service) Leave(connID uuid.UUID, ch channel) error {
	conn, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("invalid connection [" + connID.String() + "]"))
	}
	conn.Channel = nil

	s.channelsMu.Lock()
	defer s.channelsMu.Unlock()

	curr, ok := s.channels[ch]
	if !ok {
		curr = make([]uuid.UUID, 0)
	}
	filtered := make([]uuid.UUID, 0)
	for _, i := range curr {
		if i != connID {
			filtered = append(filtered, i)
		}
	}

	if len(filtered) == 0 {
		delete(s.channels, ch)
		return nil
	} else {
		s.channels[ch] = filtered
		return s.SendOnline(ch)
	}
}

func contains(s []uuid.UUID, e uuid.UUID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
