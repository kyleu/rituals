package config

import (
	"emperror.dev/emperror"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"github.com/kyleu/rituals.dev/internal/app/invite"
	"github.com/kyleu/rituals.dev/internal/app/retro"
	"github.com/kyleu/rituals.dev/internal/app/socket"
	"github.com/kyleu/rituals.dev/internal/app/standup"
	"github.com/kyleu/rituals.dev/internal/app/user"
	"logur.dev/logur"
)

type AppInfo struct {
	Debug    bool
	Version  string
	Commit   string
	Logger   logur.LoggerFacade
	Errors   emperror.ErrorHandlerFacade
	User     user.Service
	Invite   invite.Service
	Estimate estimate.Service
	Standup  standup.Service
	Retro    retro.Service
	Socket   socket.Service
}

func (a *AppInfo) Valid() bool {
	return true
}
