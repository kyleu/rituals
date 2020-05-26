package config

import (
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/database"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/invitation"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/socket"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/user"
	"logur.dev/logur"
)

type AppInfo struct {
	Debug      bool
	Version    string
	Commit     string
	Logger     logur.Logger
	User       *user.Service
	Auth       *auth.Service
	Action     *action.Service
	Invitation *invitation.Service
	Team       *team.Service
	Sprint     *sprint.Service
	Estimate   *estimate.Service
	Standup    *standup.Service
	Retro      *retro.Service
	Socket     *socket.Service
	Database   *database.Service
}

func (a *AppInfo) Valid() bool {
	return true
}
