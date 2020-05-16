package socket

import (
	"encoding/json"
	"fmt"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

func (s *Service) Write(connID uuid.UUID, message string) error {
	if connID == systemID {
		s.logger.Warn("--- admin message sent ---")
		s.logger.Warn(fmt.Sprint(message))
		return nil
	}

	c, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("cannot load connection [" + connID.String() + "]"))
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.socket.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "unable to write to websocket"))
	}
	return nil
}

func (s *Service) WriteMessage(connID uuid.UUID, message *Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error marshalling websocket message"))
	}
	return s.Write(connID, string(data))
}

func (s *Service) WriteChannel(channel channel, message *Message, except ...uuid.UUID) error {
	data, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error marshalling websocket message"))
	}

	conns, ok := s.channels[channel]
	if !ok {
		s.logger.Warn("unable to load missing channel [" + channel.String() + "]")
		return nil
	}

	// s.logger.Debug(fmt.Sprintf("sending message [%v::%v] to [%v] connections", message.Svc, message.Cmd, len(conns)))
	for _, conn := range conns {
		if !contains(except, conn) {
			connID := conn
			go func() {
				_ = s.Write(connID, string(data))
			}()
		}
	}
	return nil
}

func (s *Service) ReadLoop(connID uuid.UUID) error {
	c, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("cannot load connection [" + connID.String() + "]"))
	}
	defer func() {
		_ = c.socket.Close()
		_, _ = s.Disconnect(connID)
		s.logger.Debug(fmt.Sprintf("closed websocket [%v]", connID.String()))
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
			_ = s.WriteMessage(c.ID, &Message{Svc: util.SvcSystem.Key, Cmd: util.ServerCmdError, Param: err.Error()})
			return errors.WithStack(errors.Wrap(err, "error handling websocket message"))
		}
	}
	return nil
}
