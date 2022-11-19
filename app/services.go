// Package app - $PF_IGNORE$
package app

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/gql"
	"github.com/kyleu/rituals/app/lib/database/migrate"
	"github.com/kyleu/rituals/app/lib/websocket"
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
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/queries/migrations"
)

type Services struct {
	Team           *team.Service
	TeamMember     *tmember.Service
	TeamHistory    *thistory.Service
	TeamPermission *tpermission.Service

	Sprint           *sprint.Service
	SprintMember     *smember.Service
	SprintHistory    *shistory.Service
	SprintPermission *spermission.Service

	Estimate           *estimate.Service
	EstimateMember     *emember.Service
	EstimateHistory    *ehistory.Service
	EstimatePermission *epermission.Service
	Story              *story.Service
	Vote               *vote.Service

	Standup           *standup.Service
	StandupMember     *umember.Service
	StandupHistory    *uhistory.Service
	StandupPermission *upermission.Service
	Report            *report.Service

	Retro           *retro.Service
	RetroMember     *rmember.Service
	RetroHistory    *rhistory.Service
	RetroPermission *rpermission.Service
	Feedback        *feedback.Service

	User    *user.Service
	Action  *action.Service
	Comment *comment.Service
	Email   *email.Service

	Workspace *workspace.Service
	GQL       *gql.Schema
	Socket    *websocket.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to run database migrations")
	}

	t := team.NewService(st.DB)
	tm := tmember.NewService(st.DB)
	th := thistory.NewService(st.DB)
	tp := tpermission.NewService(st.DB)

	s := sprint.NewService(st.DB)
	sm := smember.NewService(st.DB)
	sh := shistory.NewService(st.DB)
	sp := spermission.NewService(st.DB)

	e := estimate.NewService(st.DB)
	em := emember.NewService(st.DB)
	eh := ehistory.NewService(st.DB)
	ep := epermission.NewService(st.DB)
	sy := story.NewService(st.DB)
	v := vote.NewService(st.DB)

	u := standup.NewService(st.DB)
	um := umember.NewService(st.DB)
	uh := uhistory.NewService(st.DB)
	up := upermission.NewService(st.DB)
	rt := report.NewService(st.DB)

	r := retro.NewService(st.DB)
	rm := rmember.NewService(st.DB)
	rh := rhistory.NewService(st.DB)
	rp := rpermission.NewService(st.DB)
	f := feedback.NewService(st.DB)

	us := user.NewService(st.DB)
	a := action.NewService(st.DB)
	c := comment.NewService(st.DB)
	el := email.NewService(st.DB)

	g := gql.NewSchema(st.GraphQL)
	w := workspace.NewService(t, th, tm, tp, s, sh, sm, sp, e, eh, em, ep, sy, v, u, uh, um, up, rt, r, rh, rm, rp, f, us, a, c, el)
	ws := websocket.NewService(w.SocketOpen, w.SocketHandler, w.SocketClose, nil)

	return &Services{
		Team: t, TeamMember: tm, TeamHistory: th, TeamPermission: tp,
		Sprint: s, SprintMember: sm, SprintHistory: sh, SprintPermission: sp,
		Estimate: e, EstimateMember: em, EstimateHistory: eh, EstimatePermission: ep, Story: sy, Vote: v,
		Standup: u, StandupMember: um, StandupHistory: uh, StandupPermission: up, Report: rt,
		Retro: r, RetroMember: rm, RetroHistory: rh, RetroPermission: rp, Feedback: f,
		User: us, Action: a, Comment: c, Email: el, Workspace: w, GQL: g, Socket: ws,
	}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
