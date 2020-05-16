package socket

import (
	"fmt"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/sprint"
	"sync"

	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/user"
	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"logur.dev/logur"
)

type Service struct {
	connections   map[uuid.UUID]*connection
	connectionsMu sync.Mutex
	channels      map[channel][]uuid.UUID
	channelsMu    sync.Mutex
	logger        logur.LoggerFacade
	actions       *action.Service
	users         *user.Service
	sprints       *sprint.Service
	estimates     *estimate.Service
	standups      *standup.Service
	retros        *retro.Service
}

func NewService(actions *action.Service, logger logur.LoggerFacade, users *user.Service, sprints *sprint.Service, estimates *estimate.Service, standups *standup.Service, retros *retro.Service) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "socket"})
	return Service{
		connections:   make(map[uuid.UUID]*connection),
		connectionsMu: sync.Mutex{},
		channels:      make(map[channel][]uuid.UUID),
		channelsMu:    sync.Mutex{},
		logger:        logger,
		actions:       actions,
		sprints:       sprints,
		estimates:     estimates,
		standups:      standups,
		retros:        retros,
		users:         users,
	}
}

var systemID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000")
var systemStatus = Status{ID: systemID, UserID: systemID, Username: "System Broadcast", ChannelSvc: util.SvcSystem.Key, ChannelID: &systemID}

func (s *Service) List() ([]*Status, error) {
	ret := make([]*Status, 0)
	ret = append(ret, &systemStatus)
	for _, conn := range s.connections {
		ret = append(ret, conn.ToStatus())
	}
	return ret, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Status, error) {
	if id == systemID {
		return &systemStatus, nil
	}
	conn, ok := s.connections[id]
	if !ok {
		return nil, errors.New("invalid connection [" + id.String() + "]")
	}
	return conn.ToStatus(), nil
}

func onMessage(s *Service, connID uuid.UUID, message Message) error {
	if connID == systemID {
		s.logger.Warn("--- admin message received ---")
		s.logger.Warn(fmt.Sprint(message))
		return nil
	}
	c, ok := s.connections[connID]
	if !ok {
		return errors.WithStack(errors.New("cannot load connection [" + connID.String() + "]"))
	}
	var err error
	switch message.Svc {
	case util.SvcSystem.Key:
		err = onSystemMessage(s, c, c.Profile.UserID, message.Cmd, message.Param)
	case util.SvcSprint.Key:
		err = onSprintMessage(s, c, c.Profile.UserID, message.Cmd, message.Param)
	case util.SvcEstimate.Key:
		err = onEstimateMessage(s, c, c.Profile.UserID, message.Cmd, message.Param)
	case util.SvcStandup.Key:
		err = onStandupMessage(s, c, c.Profile.UserID, message.Cmd, message.Param)
	case util.SvcRetro.Key:
		err = onRetroMessage(s, c, c.Profile.UserID, message.Cmd, message.Param)
	default:
		err = errors.New("invalid service [" + message.Svc + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling socket message ["+message.String()+"]"))
}
