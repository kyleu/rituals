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
		return errors.New("cannot load connection [" + connID.String() + "]")
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.socket.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return errors.Wrap(err, "unable to write to websocket")
	}
	return nil
}

func (s *Service) WriteMessage(connID uuid.UUID, message *Message) error {
	return s.Write(connID, util.ToJSON(message))
}

func (s *Service) WriteChannel(channel channel, message *Message, except ...uuid.UUID) error {
	conns, ok := s.channels[channel]
	if !ok {
		return nil
	}

	// s.logger.Debug(fmt.Sprintf("sending message [%v::%v] to [%v] connections", message.Svc, message.Cmd, len(conns)))
	for _, conn := range conns {
		if !contains(except, conn) {
			connID := conn

			go func() {
				_ = s.Write(connID, util.ToJSON(message))
			}()
		}
	}
	return nil
}

func (s *Service) ReadLoop(connID uuid.UUID) error {
	c, ok := s.connections[connID]
	if !ok {
		return errors.New("cannot load connection [" + connID.String() + "]")
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
			return errors.Wrap(err, "error decoding websocket message")
		}

		err = onMessage(s, connID, m)
		if err != nil {
			_ = s.WriteMessage(c.ID, &Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdError, Param: err.Error()})
			return errors.Wrap(err, "error handling websocket message")
		}
	}
	return nil
}
