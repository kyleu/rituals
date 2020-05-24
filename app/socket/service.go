package socket

import (
	"fmt"
	"sync"

	"github.com/kyleu/rituals.dev/app/auth"

	"github.com/kyleu/rituals.dev/app/team"

	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/sprint"

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
	logger        logur.Logger
	actions       *action.Service
	users         *user.Service
	auths         *auth.Service
	teams         *team.Service
	sprints       *sprint.Service
	estimates     *estimate.Service
	standups      *standup.Service
	retros        *retro.Service
}

func NewService(
	logger logur.Logger, actions *action.Service, users *user.Service, auths *auth.Service,
	teams *team.Service, sprints *sprint.Service,
	estimates *estimate.Service, standups *standup.Service, retros *retro.Service) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.KeySocket})
	return Service{
		connections:   make(map[uuid.UUID]*connection),
		connectionsMu: sync.Mutex{},
		channels:      make(map[channel][]uuid.UUID),
		channelsMu:    sync.Mutex{},
		logger:        logger,
		actions:       actions,
		users:         users,
		auths:         auths,
		teams:         teams,
		sprints:       sprints,
		estimates:     estimates,
		standups:      standups,
		retros:        retros,
	}
}

var systemID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000")
var systemStatus = Status{ID: systemID, UserID: systemID, Username: "System Broadcast", ChannelSvc: util.SvcSystem.Key, ChannelID: &systemID}

func (s *Service) List(params *query.Params) ([]*Status, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyConnection, params)
	ret := make([]*Status, 0)
	ret = append(ret, &systemStatus)
	var idx = 0
	for _, conn := range s.connections {
		if idx >= params.Offset && (params.Limit == 0 || idx < params.Limit) {
			ret = append(ret, conn.ToStatus())
		}
		idx++
	}
	return ret, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Status, error) {
	if id == systemID {
		return &systemStatus, nil
	}
	conn, ok := s.connections[id]
	if !ok {
		return nil, invalidConnection(id)
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
		return invalidConnection(connID)
	}
	var err error

	switch message.Svc {
	case util.SvcSystem.Key:
		err = onSystemMessage(s, c, c.Profile.UserID, message.Cmd, message.Param)
	case util.SvcTeam.Key:
		err = onTeamMessage(s, c, c.Profile.UserID, message.Cmd, message.Param)
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
