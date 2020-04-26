package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"logur.dev/logur"
)

type Service struct {
	connections map[uuid.UUID]connection
	logger      logur.LoggerFacade
}

func NewSocketService(logger logur.LoggerFacade) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "socket"})
	return Service{
		connections: make(map[uuid.UUID]connection),
		logger: logger,
	}
}

func (s *Service) Register(userID uuid.UUID, c *websocket.Conn) (uuid.UUID, error) {
	conn := connection{
		ID:     util.UUID(),
		UserID: userID,
		socket: c,
	}
	s.connections[conn.ID] = conn
	return conn.ID, nil
}

func (s *Service) Write(connID uuid.UUID, message string) error {
	c, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("cannot load connection [" + connID.String() + "]"))
	}
	err := c.socket.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "unable to write to websocket"))
	}
	return nil
}

func (s *Service) WriteMessage(connID uuid.UUID, message Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error marshalling websocket message"))
	}
	return s.Write(connID, string(data))
}

func (s *Service) ReadLoop(connID uuid.UUID) error {
	c, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("cannot load connection [" + connID.String() + "]"))
	}
	defer func() {
		s.logger.Debug("closing websocket [" + connID.String() + "]")
		_ = c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			break
		}

		var m Message
		err = json.Unmarshal(message, &m)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error decoding websocket message"))
		}

		err = s.OnMessage(connID, m)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error handling websocket message"))
		}
	}
	return nil
}

func (s *Service) OnMessage(connID uuid.UUID, message Message) error {
	switch message.T {
	case "connect":
		err := s.onConnect(connID, message.K, message.V)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error connecting session"))
		}
	default:
		s.logger.Warn("unhandled message of type [" + message.T + "]")
	}

	err := s.WriteMessage(connID, message)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing websocket message"))
	}

	return nil
}

func (s *Service) onConnect(connID uuid.UUID, k string, v string) error {
	s.logger.Warn("Connect TODO")
	switch k {
	case "estimate":
		err := s.Write(connID, "{}")
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error connecting websocket"))
		}
	}
	return nil
}
