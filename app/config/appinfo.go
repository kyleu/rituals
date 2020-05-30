package config

import (
	"github.com/kyleu/rituals.dev/app/database"
	"github.com/kyleu/rituals.dev/app/model/action"
	"github.com/kyleu/rituals.dev/app/model/auth"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/invitation"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/model/user"
	"github.com/kyleu/rituals.dev/app/socket"
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
