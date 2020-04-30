package socket

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"logur.dev/logur"
)

type Service struct {
	connections map[uuid.UUID]connection
	channels    map[uuid.UUID][]uuid.UUID
	logger      logur.LoggerFacade
	estimates   *estimate.Service
}

func NewSocketService(logger logur.LoggerFacade, estimates *estimate.Service) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "socket"})
	return Service{
		connections: make(map[uuid.UUID]connection),
		channels:    make(map[uuid.UUID][]uuid.UUID),
		logger:      logger,
		estimates:   estimates,
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

func contains(s []uuid.UUID, e uuid.UUID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s *Service) Join(connID uuid.UUID, channelID uuid.UUID) error {
	curr, ok := s.channels[channelID]
	if !ok {
		curr = make([]uuid.UUID, 0)
	}
	if !contains(curr, connID) {
		s.channels[channelID] = append(curr, connID)
	}
	return nil
}

func (s *Service) Leave(connID uuid.UUID, channelID uuid.UUID) error {
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

func onMessage(s *Service, connID uuid.UUID, message Message) error {
	c, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("cannot load connection [" + connID.String() + "]"))
	}
	var err error = nil
	switch message.Svc {
	case "estimate":
		err = onEstimateMessage(s, connID, c.UserID, message.Cmd, message.Param)
	default:
		s.logger.Warn("unhandled message of type [" + message.Svc + "]")
	}
	return errors.WithStack(errors.Wrap(err, fmt.Sprintf("error handling message [%s] %s / %s", message.Svc, message.Cmd, message.Param)))
}
