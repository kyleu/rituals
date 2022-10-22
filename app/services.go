// Package app - $PF_IGNORE$
package app

import (
	"context"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
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
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/standup/upermission"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/lib/database/migrate"
	"github.com/kyleu/rituals/app/lib/websocket"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/queries/migrations"
)

type Services struct {
	User               *user.Service
	Team               *team.Service
	TeamMember         *tmember.Service
	TeamHistory        *thistory.Service
	TeamPermission     *tpermission.Service
	Sprint             *sprint.Service
	SprintMember       *smember.Service
	SprintHistory      *shistory.Service
	SprintPermission   *spermission.Service
	Estimate           *estimate.Service
	EstimateMember     *emember.Service
	EstimateHistory    *ehistory.Service
	EstimatePermission *epermission.Service
	Story              *story.Service
	Vote               *vote.Service
	Standup            *standup.Service
	StandupMember      *umember.Service
	StandupHistory     *uhistory.Service
	StandupPermission  *upermission.Service
	Report             *report.Service
	Retro              *retro.Service
	RetroMember        *rmember.Service
	RetroHistory       *rhistory.Service
	RetroPermission    *rpermission.Service
	Feedback           *feedback.Service
	Action             *action.Service
	Comment            *comment.Service
	Email              *email.Service
	Socket             *websocket.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to run database migrations")
	}
	return &Services{
		User:               user.NewService(st.DB),
		Team:               team.NewService(st.DB),
		TeamMember:         tmember.NewService(st.DB),
		TeamHistory:        thistory.NewService(st.DB),
		TeamPermission:     tpermission.NewService(st.DB),
		Sprint:             sprint.NewService(st.DB),
		SprintMember:       smember.NewService(st.DB),
		SprintHistory:      shistory.NewService(st.DB),
		SprintPermission:   spermission.NewService(st.DB),
		Estimate:           estimate.NewService(st.DB),
		EstimateMember:     emember.NewService(st.DB),
		EstimateHistory:    ehistory.NewService(st.DB),
		EstimatePermission: epermission.NewService(st.DB),
		Story:              story.NewService(st.DB),
		Vote:               vote.NewService(st.DB),
		Standup:            standup.NewService(st.DB),
		StandupMember:      umember.NewService(st.DB),
		StandupHistory:     uhistory.NewService(st.DB),
		StandupPermission:  upermission.NewService(st.DB),
		Report:             report.NewService(st.DB),
		Retro:              retro.NewService(st.DB),
		RetroMember:        rmember.NewService(st.DB),
		RetroHistory:       rhistory.NewService(st.DB),
		RetroPermission:    rpermission.NewService(st.DB),
		Feedback:           feedback.NewService(st.DB),
		Action:             action.NewService(st.DB),
		Comment:            comment.NewService(st.DB),
		Email:              email.NewService(st.DB),
		Socket:             websocket.NewService(logger, nil, nil, nil, nil),
	}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
