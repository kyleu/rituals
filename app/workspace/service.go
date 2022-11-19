package workspace

import (
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
)

type Service struct {
	t  *team.Service
	th *thistory.Service
	tm *tmember.Service
	tp *tpermission.Service

	s  *sprint.Service
	sh *shistory.Service
	sm *smember.Service
	sp *spermission.Service

	e  *estimate.Service
	eh *ehistory.Service
	em *emember.Service
	ep *epermission.Service
	st *story.Service
	v  *vote.Service

	u  *standup.Service
	uh *uhistory.Service
	um *umember.Service
	up *upermission.Service
	rt *report.Service

	r  *retro.Service
	rh *rhistory.Service
	rm *rmember.Service
	rp *rpermission.Service
	f  *feedback.Service

	us *user.Service
	a  *action.Service
	c  *comment.Service
	el *email.Service
}

func NewService(
	t *team.Service, th *thistory.Service, tm *tmember.Service, tp *tpermission.Service,
	s *sprint.Service, sh *shistory.Service, sm *smember.Service, sp *spermission.Service,
	e *estimate.Service, eh *ehistory.Service, em *emember.Service, ep *epermission.Service, st *story.Service, v *vote.Service,
	u *standup.Service, uh *uhistory.Service, um *umember.Service, up *upermission.Service, rt *report.Service,
	r *retro.Service, rh *rhistory.Service, rm *rmember.Service, rp *rpermission.Service, f *feedback.Service,
	us *user.Service, a *action.Service, c *comment.Service, el *email.Service,
) *Service {
	return &Service{
		t: t, th: th, tm: tm, tp: tp,
		s: s, sh: sh, sm: sm, sp: sp,
		e: e, eh: eh, em: em, ep: ep, st: st, v: v,
		u: u, uh: uh, um: um, up: up, rt: rt,
		r: r, rh: rh, rm: rm, rp: rp, f: f,
		us: us, a: a, c: c, el: el,
	}
}
