package socket

import (
	"fmt"
	"github.com/kyleu/rituals.dev/internal/app/member"
	"github.com/kyleu/rituals.dev/internal/app/retro"
	"github.com/kyleu/rituals.dev/internal/app/standup"
	"github.com/kyleu/rituals.dev/internal/app/user"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"sync"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"logur.dev/logur"
)

type Service struct {
	connections   map[uuid.UUID]*connection
	connectionsMu sync.Mutex
	channels      map[channel][]uuid.UUID
	channelsMu    sync.Mutex
	logger        logur.LoggerFacade
	users         *user.Service
	estimates     *estimate.Service
	standups      *standup.Service
	retros        *retro.Service
}

func NewSocketService(logger logur.LoggerFacade, users *user.Service, estimates *estimate.Service) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "socket"})
	return Service{
		connections:   make(map[uuid.UUID]*connection),
		connectionsMu: sync.Mutex{},
		channels:      make(map[channel][]uuid.UUID),
		channelsMu:    sync.Mutex{},
		logger:        logger,
		estimates:     estimates,
		users:         users,
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
		return nil, errors.New("invalid connection [" + id.String() + "]")
	}
	return &Status{ID: conn.ID, UserID: conn.UserID}, nil
}

func (s *Service) UpdateName(id uuid.UUID, name string) error {
	return s.users.UpdateUserName(id, name)
}

func (s *Service) SendMembers(memberSvc *member.Service, ch channel, connID *uuid.UUID) error {
	members, err := memberSvc.GetByModelID(ch.ID)
	if err != nil {
		return err
	}
	membersMsg := Message{Svc: util.SvcSystem, Cmd: "members", Param: members}
	if connID == nil {
		err = s.WriteChannel(ch, &membersMsg)
	} else {
		err = s.WriteMessage(*connID, &membersMsg)
	}
	if err != nil {
		return err
	}

	return s.SendOnline(ch)
}

func (s *Service) SendOnline(ch channel) error {
	connections, ok := s.channels[ch]
	if !ok {
		connections = make([]uuid.UUID, 0)
	}
	online := make([]uuid.UUID, 0)
	for _, cID := range connections {
		c, ok := s.connections[cID]
		if ok && c != nil && (!contains(online, c.UserID)) {
			online = append(online, c.UserID)
		}
	}

	onlineMsg := Message{Svc: util.SvcSystem, Cmd: "online", Param: online}
	return s.WriteChannel(ch, &onlineMsg)
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
	case "system":
		err = onSystemMessage(s, c, c.UserID, message.Cmd, message.Param)
	case util.SvcEstimate:
		err = onEstimateMessage(s, c, c.UserID, message.Cmd, message.Param)
	case util.SvcStandup:
		err = onStandupMessage(s, c, c.UserID, message.Cmd, message.Param)
	case util.SvcRetro:
		err = onRetroMessage(s, c, c.UserID, message.Cmd, message.Param)
	default:
		return errors.WithStack(errors.New("invalid service [" + message.Svc + "]"))
	}
	return errors.WithStack(errors.Wrap(err, "error handling message ["+message.String()+"]"))
}
