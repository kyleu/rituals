package workspace

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/member"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type FullSprint struct {
	Sprint      *sprint.Sprint                `json:"sprint"`
	Histories   shistory.SprintHistories      `json:"histories,omitempty"`
	Members     smember.SprintMembers         `json:"members,omitempty"`
	UtilMembers member.Members                `json:"-"`
	Self        *smember.SprintMember         `json:"self,omitempty"`
	Permissions spermission.SprintPermissions `json:"permissions,omitempty"`
	Team        *team.Team                    `json:"team,omitempty"`
	Estimates   estimate.Estimates            `json:"estimates,omitempty"`
	Standups    standup.Standups              `json:"standups,omitempty"`
	Retros      retro.Retros                  `json:"retros,omitempty"`
	Comments    comment.Comments              `json:"comments,omitempty"`
	Actions     action.Actions                `json:"actions,omitempty"`
	Registered  bool                          `json:"registered,omitempty"`
}

func (f *FullSprint) Admin() bool {
	return f.Self.Role == enum.MemberStatusOwner
}

func (s *Service) LoadSprint(p *LoadParams, tf func() (team.Teams, error)) (*FullSprint, error) {
	spr, err := s.s.GetBySlug(p.Ctx, p.Tx, p.Slug, p.Logger)
	if err != nil {
		if hist, _ := s.sh.Get(p.Ctx, p.Tx, p.Slug, p.Logger); hist != nil {
			spr, err = s.s.Get(p.Ctx, p.Tx, hist.SprintID, p.Logger)
			if err != nil {
				return nil, errors.Errorf("no sprint found with slug [%s]", p.Slug)
			}
		}
	}
	if spr == nil {
		id := util.UUIDFromString(p.Slug)
		if id == nil {
			return nil, errors.Errorf("no sprint found with slug [%s]", p.Slug)
		}
		spr, err = s.s.Get(p.Ctx, p.Tx, *id, p.Logger)
		if err != nil {
			return nil, errors.Errorf("no sprint found with id [%s]", p.Slug)
		}
	}
	return s.loadFullSprint(p, spr, tf)
}

func (s *Service) loadFullSprint(p *LoadParams, spr *sprint.Sprint, tf func() (team.Teams, error)) (*FullSprint, error) {
	ret := &FullSprint{Sprint: spr}

	var er error
	ret.Permissions, er = s.sp.GetBySprintID(p.Ctx, p.Tx, spr.ID, p.Params.Get("spermission", nil, p.Logger).Sanitize("spermission"), p.Logger)
	if er != nil {
		return nil, er
	}
	if ok, msg := CheckPermissions(util.KeySprint, ret.Permissions.ToPermissions(), p.Accounts, spr.TeamID, tf, nil, nil); !ok {
		return nil, errors.New(msg)
	}

	funcs := []func() error{
		func() error {
			var err error
			ret.Histories, err = s.sh.GetBySprintID(p.Ctx, p.Tx, spr.ID, p.Params.Get("shistory", nil, p.Logger).Sanitize("shistory"), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Members, ret.Self, ret.Registered, err = s.membersSprint(p, spr.ID)
			online := s.online(util.KeySprint + ":" + spr.ID.String())
			ret.UtilMembers = ret.Members.ToMembers(online)
			return err
		},
		func() error {
			var err error
			if spr.TeamID != nil {
				ret.Team, err = s.t.Get(p.Ctx, p.Tx, *spr.TeamID, p.Logger)
			}
			return err
		},
		func() error {
			var err error
			prm := p.Params.Get(util.KeyEstimate, nil, p.Logger).Sanitize(util.KeyEstimate)
			ret.Estimates, err = s.e.GetBySprintID(p.Ctx, p.Tx, &spr.ID, prm, p.Logger)
			return err
		},
		func() error {
			var err error
			prm := p.Params.Get(util.KeyStandup, nil, p.Logger).Sanitize(util.KeyStandup)
			ret.Standups, err = s.u.GetBySprintID(p.Ctx, p.Tx, &spr.ID, prm, p.Logger)
			return err
		},
		func() error {
			var err error
			prm := p.Params.Get(util.KeyRetro, nil, p.Logger).Sanitize(util.KeyRetro)
			ret.Retros, err = s.r.GetBySprintID(p.Ctx, p.Tx, &spr.ID, prm, p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Actions, err = s.a.GetByModels(p.Ctx, p.Tx, p.Logger, enum.ModelServiceSprint, ret.Sprint.ID)
			return err
		},
	}
	_, errs := util.AsyncCollect(funcs, func(f func() error) (any, error) {
		return nil, f()
	})
	if len(errs) > 0 {
		return nil, util.ErrorMerge(errs...)
	}

	args := make([]any, 0, (len(ret.Estimates)*2)+(len(ret.Standups)*2)+(len(ret.Retros)*2)+2)
	args = append(args, util.KeySprint, spr.ID)
	for _, x := range ret.Estimates {
		args = append(args, util.KeyEstimate, x.ID)
	}
	for _, x := range ret.Standups {
		args = append(args, util.KeyStandup, x.ID)
	}
	for _, x := range ret.Retros {
		args = append(args, util.KeyRetro, x.ID)
	}

	var err error
	ret.Comments, err = s.c.GetByModels(p.Ctx, p.Tx, p.Logger, args...)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Service) membersSprint(p *LoadParams, sprintID uuid.UUID) (smember.SprintMembers, *smember.SprintMember, bool, error) {
	params := p.Params.Get("smember", nil, p.Logger).Sanitize("smember")
	members, err := s.sm.GetBySprintID(p.Ctx, p.Tx, sprintID, params, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	err = s.us.CreateIfNeeded(p.Ctx, p.Profile.ID, p.Profile.Name, p.Tx, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	if self := members.Get(sprintID, p.Profile.ID); self != nil {
		return members, self, false, nil
	}
	role := enum.MemberStatusMember
	if len(members) == 0 {
		role = enum.MemberStatusOwner
	}
	_, err = s.sm.Register(p.Ctx, sprintID, p.Profile.ID, p.Profile.Name, p.Accounts.Image(), role, p.Tx, s.a, s.send, s.us, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	members, err = s.sm.GetBySprintID(p.Ctx, p.Tx, sprintID, params, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	return members, members.Get(sprintID, p.Profile.ID), true, nil
}
