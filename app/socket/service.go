package socket

import (
	"fmt"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"sync"

	"github.com/kyleu/rituals.dev/app/model/auth"

	"github.com/kyleu/rituals.dev/app/model/team"

	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/action"
	"github.com/kyleu/rituals.dev/app/model/sprint"

	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/user"
	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"logur.dev/logur"
)

type Service struct {
	connections   map[uuid.UUID]*connection
	connectionsMu sync.Mutex
	channels      map[Channel][]uuid.UUID
	channelsMu    sync.Mutex
	Logger        logur.Logger
	comments      *comment.Service
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
		logger logur.Logger, actions *action.Service, users *user.Service, comments *comment.Service,
		auths *auth.Service, teams *team.Service, sprints *sprint.Service,
		estimates *estimate.Service, standups *standup.Service, retros *retro.Service) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.KeySocket})
	return &Service{
		connections:   make(map[uuid.UUID]*connection),
		connectionsMu: sync.Mutex{},
		channels:      make(map[Channel][]uuid.UUID),
		channelsMu:    sync.Mutex{},
		Logger:        logger,
		comments:      comments,
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

var systemID = uuid.FromStringOrNil("FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF")
var systemStatus = Status{ID: systemID, UserID: systemID, Username: "System Broadcast", ChannelSvc: util.SvcSystem, ChannelID: &systemID}

func (s *Service) List(params *query.Params) Statuses {
	params = query.ParamsWithDefaultOrdering(util.KeyConnection, params)
	ret := make(Statuses, 0)
	ret = append(ret, &systemStatus)
	var idx = 0
	for _, conn := range s.connections {
		if idx >= params.Offset && (params.Limit == 0 || idx < params.Limit) {
			ret = append(ret, conn.ToStatus())
		}
		idx++
	}
	return ret
}

func (s *Service) GetByID(id uuid.UUID) *Status {
	if id == systemID {
		return &systemStatus
	}
	conn, ok := s.connections[id]
	if !ok {
		util.LogError(s.Logger, "error getting connection by id [%v]", id)
		return nil
	}
	return conn.ToStatus()
}

func (s *Service) Count() int {
	return len(s.connections)
}

func (s *Service) RemoveComment(commentID uuid.UUID, userID uuid.UUID) error {
	c := s.GetByID(commentID)
	if c.UserID != userID {
		return errors.New("This is not your comment")
	}
	return s.comments.RemoveComment(commentID)
}

func onMessage(s *Service, connID uuid.UUID, message Message) error {
	if connID == systemID {
		s.Logger.Warn("--- admin message received ---")
		s.Logger.Warn(fmt.Sprint(message))
		return nil
	}
	c, ok := s.connections[connID]
	if !ok {
		return invalidConnection(connID)
	}
	var err error

	switch message.Svc {
	case util.SvcSystem.Key:
		err = onSystemMessage(s, c, message.Cmd, message.Param)
	case util.SvcTeam.Key:
		err = onTeamMessage(s, c, message.Cmd, message.Param)
	case util.SvcSprint.Key:
		err = onSprintMessage(s, c, message.Cmd, message.Param)
	case util.SvcEstimate.Key:
		err = onEstimateMessage(s, c, message.Cmd, message.Param)
	case util.SvcStandup.Key:
		err = onStandupMessage(s, c, message.Cmd, message.Param)
	case util.SvcRetro.Key:
		err = onRetroMessage(s, c, message.Cmd, message.Param)
	default:
		err = errors.New(util.IDErrorString(util.KeyService, message.Svc))
	}
	return errors.Wrap(err, "error handling socket message ["+message.String()+"]")
}
