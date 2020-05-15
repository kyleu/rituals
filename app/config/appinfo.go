package config

import (
	"emperror.dev/emperror"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/invite"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/socket"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/user"
	"logur.dev/logur"
)

type AppInfo struct {
	Debug    bool
	Version  string
	Commit   string
	Logger   logur.LoggerFacade
	Errors   emperror.ErrorHandlerFacade
	User     *user.Service
	Auth     *auth.Service
	Action   *action.Service
	Invite   *invite.Service
	Estimate *estimate.Service
	Standup  *standup.Service
	Retro    *retro.Service
	Sprint   *sprint.Service
	Socket   *socket.Service
}

func (a *AppInfo) Valid() bool {
	return true
}
