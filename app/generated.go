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

func initGeneratedServices(ctx context.Context, st *State, logger util.Logger) GeneratedServices {
	return GeneratedServices{
		Action:             action.NewService(st.DB),
		Comment:            comment.NewService(st.DB),
		Email:              email.NewService(st.DB),
		Estimate:           estimate.NewService(st.DB),
		EstimateHistory:    ehistory.NewService(st.DB),
		EstimateMember:     emember.NewService(st.DB),
		EstimatePermission: epermission.NewService(st.DB),
		Feedback:           feedback.NewService(st.DB),
		Report:             report.NewService(st.DB),
		Retro:              retro.NewService(st.DB),
		RetroHistory:       rhistory.NewService(st.DB),
		RetroMember:        rmember.NewService(st.DB),
		RetroPermission:    rpermission.NewService(st.DB),
		Sprint:             sprint.NewService(st.DB),
		SprintHistory:      shistory.NewService(st.DB),
		SprintMember:       smember.NewService(st.DB),
		SprintPermission:   spermission.NewService(st.DB),
		Standup:            standup.NewService(st.DB),
		StandupHistory:     uhistory.NewService(st.DB),
		StandupMember:      umember.NewService(st.DB),
		StandupPermission:  upermission.NewService(st.DB),
		Story:              story.NewService(st.DB),
		Team:               team.NewService(st.DB),
		TeamHistory:        thistory.NewService(st.DB),
		TeamMember:         tmember.NewService(st.DB),
		TeamPermission:     tpermission.NewService(st.DB),
		User:               user.NewService(st.DB),
		Vote:               vote.NewService(st.DB),
	}
}
