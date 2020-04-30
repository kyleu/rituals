package socket

import (
	"encoding/json"
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

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

func (s *Service) WriteChannel(channelID uuid.UUID, message Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error marshalling websocket message"))
	}

	conns, ok := s.channels[channelID]
	if !ok {
		s.logger.Warn("unable to load missing channel [" + channelID.String() + "]")
		return nil
	}

	s.logger.Debug(fmt.Sprintf("sending message [%v::%v] to [%v] connections", message.Svc, message.Cmd, len(conns)))
	for _, conn := range conns {
		connID := conn
		go func() {
			_ = s.Write(connID, string(data))
		}()
	}
	return nil
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

		err = onMessage(s, connID, m)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error handling websocket message"))
		}
	}
	return nil
}
