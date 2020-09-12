package socket

import (
	"encoding/json"

	"github.com/kyleu/npn/npnconnection"

	"github.com/kyleu/npn/npncore"

	"github.com/kyleu/rituals.dev/app/comment"

	"github.com/kyleu/npn/npnservice/auth"

	"github.com/kyleu/rituals.dev/app/team"

	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/sprint"

	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/estimate"
	"logur.dev/logur"
)

type services struct {
	comments  *comment.Service
	actions   *action.Service
	teams     *team.Service
	sprints   *sprint.Service
	estimates *estimate.Service
	standups  *standup.Service
	retros    *retro.Service
	users     *user.Service
	auths     *auth.Service
}

func NewService(
		logger logur.Logger, actions *action.Service, users *user.Service, comments *comment.Service,
		auths *auth.Service, teams *team.Service, sprints *sprint.Service,
		estimates *estimate.Service, standups *standup.Service, retros *retro.Service) *npnconnection.Service {
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: npncore.KeySocket})

	ctx := &services{
		comments:  comments,
		actions:   actions,
		teams:     teams,
		sprints:   sprints,
		estimates: estimates,
		standups:  standups,
		retros:    retros,
		auths:     auths,
		users:     users,
	}

	return npnconnection.NewService(logger, handler, ctx)
}

func ctx(s *npnconnection.Service) *services {
	return s.Context.(*services)
}

func handler(s *npnconnection.Service, c *npnconnection.Connection, svc string, cmd string, param json.RawMessage) error {
	var err error
	switch svc {
	case util.SvcSystem.Key:
		err = onSystemMessage(s, ctx(s).users, c, cmd, param)
	case util.SvcTeam.Key:
		err = onTeamMessage(s, ctx(s).auths, c, cmd, param)
	case util.SvcSprint.Key:
		err = onSprintMessage(s, ctx(s).auths, c, cmd, param)
	case util.SvcEstimate.Key:
		err = onEstimateMessage(s, ctx(s).auths, c, cmd, param)
	case util.SvcStandup.Key:
		err = onStandupMessage(s, ctx(s).auths, c, cmd, param)
	case util.SvcRetro.Key:
		err = onRetroMessage(s, ctx(s).auths, c, cmd, param)
	default:
		err = errors.New(npncore.IDErrorString(npncore.KeyService, svc))
	}
	return errors.Wrap(err, "error handling socket message ["+cmd+"]")
}
