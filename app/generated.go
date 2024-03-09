// Package app - Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"context"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
)

type GeneratedServices struct {
	Action             *action.Service
	Comment            *comment.Service
	Email              *email.Service
	Estimate           *estimate.Service
	EstimateHistory    *ehistory.Service
	EstimateMember     *emember.Service
	EstimatePermission *epermission.Service
	Feedback           *feedback.Service
	Report             *report.Service
	Retro              *retro.Service
	RetroHistory       *rhistory.Service
	RetroMember        *rmember.Service
	RetroPermission    *rpermission.Service
	Sprint             *sprint.Service
	SprintHistory      *shistory.Service
	SprintMember       *smember.Service
	SprintPermission   *spermission.Service
	Standup            *standup.Service
	StandupHistory     *uhistory.Service
	StandupMember      *umember.Service
	StandupPermission  *upermission.Service
	Story              *story.Service
	Team               *team.Service
	TeamHistory        *thistory.Service
	TeamMember         *tmember.Service
	TeamPermission     *tpermission.Service
	User               *user.Service
	Vote               *vote.Service
}

func initGeneratedServices(ctx context.Context, dbSvc *database.Service, logger util.Logger) GeneratedServices {
	return GeneratedServices{
		Action:             action.NewService(dbSvc),
		Comment:            comment.NewService(dbSvc),
		Email:              email.NewService(dbSvc),
		Estimate:           estimate.NewService(dbSvc),
		EstimateHistory:    ehistory.NewService(dbSvc),
		EstimateMember:     emember.NewService(dbSvc),
		EstimatePermission: epermission.NewService(dbSvc),
		Feedback:           feedback.NewService(dbSvc),
		Report:             report.NewService(dbSvc),
		Retro:              retro.NewService(dbSvc),
		RetroHistory:       rhistory.NewService(dbSvc),
		RetroMember:        rmember.NewService(dbSvc),
		RetroPermission:    rpermission.NewService(dbSvc),
		Sprint:             sprint.NewService(dbSvc),
		SprintHistory:      shistory.NewService(dbSvc),
		SprintMember:       smember.NewService(dbSvc),
		SprintPermission:   spermission.NewService(dbSvc),
		Standup:            standup.NewService(dbSvc),
		StandupHistory:     uhistory.NewService(dbSvc),
		StandupMember:      umember.NewService(dbSvc),
		StandupPermission:  upermission.NewService(dbSvc),
		Story:              story.NewService(dbSvc),
		Team:               team.NewService(dbSvc),
		TeamHistory:        thistory.NewService(dbSvc),
		TeamMember:         tmember.NewService(dbSvc),
		TeamPermission:     tpermission.NewService(dbSvc),
		User:               user.NewService(dbSvc),
		Vote:               vote.NewService(dbSvc),
	}
}
