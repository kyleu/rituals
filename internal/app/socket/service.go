package socket

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"logur.dev/logur"
)

type Service struct {
	connections map[uuid.UUID]*connection
	channels    map[uuid.UUID][]uuid.UUID
	logger      logur.LoggerFacade
	estimates   *estimate.Service
}

func NewSocketService(logger logur.LoggerFacade, estimates *estimate.Service) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "socket"})
	return Service{
		connections: make(map[uuid.UUID]*connection),
		channels:    make(map[uuid.UUID][]uuid.UUID),
		logger:      logger,
		estimates:   estimates,
	}
}

var systemID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000")

func (s *Service) List() ([]*Status, error) {
	ret := make([]*Status, 0)
	ret = append(ret, &Status{ID: systemID, UserID: systemID})
	for _, conn := range s.connections {
		ret = append(ret, &Status{ID: conn.ID, UserID: conn.UserID})
	}
	return ret, nil
}

func (s *Service) GetById(id uuid.UUID) (*Status, error) {
	if id == systemID {
		return &Status{ID: systemID, UserID: systemID}, nil
	}
	conn, ok := s.connections[id]
	if !ok {
		return nil, nil
	}
	return &Status{ID: conn.ID, UserID: conn.UserID}, nil
}

func onMessage(s *Service, connID uuid.UUID, message Message) error {
	if connID == systemID {
		s.logger.Warn("--- admin message received ---")
		s.logger.Warn(fmt.Sprintf("%v", message))
		return nil
	}
	c, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("cannot load connection [" + connID.String() + "]"))
	}
	var err error = nil
	switch message.Svc {
	case "estimate":
		err = onEstimateMessage(s, connID, c.UserID, message.Cmd, message.Param)
	default:
		return errors.WithStack(errors.New("invalid service [" + message.Svc + "]"))
	}
	return errors.WithStack(errors.Wrap(err, "error handling message [" + message.String() + "]"))
}
