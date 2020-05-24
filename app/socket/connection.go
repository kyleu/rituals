package socket

import (
	"sync"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) Register(profile util.Profile, c *websocket.Conn) (uuid.UUID, error) {
	conn := &connection{
		ID:      util.UUID(),
		Profile: profile,
		Svc:     "",
		ModelID: nil,
		Channel: nil,
		socket:  c,
		mu:      sync.Mutex{},
	}

	s.connectionsMu.Lock()
	defer s.connectionsMu.Unlock()

	s.connections[conn.ID] = conn
	return conn.ID, nil
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

func invalidConnection(id uuid.UUID) error {
	return errors.WithStack(errors.New("invalid connection [" + id.String() + "]"))
}
