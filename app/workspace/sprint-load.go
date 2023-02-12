package workspace

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
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
	UtilMembers util.Members                  `json:"-"`
	Self        *smember.SprintMember         `json:"self,omitempty"`
	Permissions spermission.SprintPermissions `json:"permissions,omitempty"`
	Team        *team.Team                    `json:"team,omitempty"`
	Estimates   estimate.Estimates            `json:"estimates,omitempty"`
	Standups    standup.Standups              `json:"standups,omitempty"`
	Retros      retro.Retros                  `json:"retros,omitempty"`
	Comments    comment.Comments              `json:"comments,omitempty"`
	Actions     action.Actions                `json:"actions,omitempty"`
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
	sf := func() (sprint.Sprints, error) { return nil, nil }
	return s.loadFullSprint(p, spr, tf, sf)
}

func (s *Service) loadFullSprint(p *LoadParams, spr *sprint.Sprint, tf func() (team.Teams, error), sf func() (sprint.Sprints, error)) (*FullSprint, error) {
	ret := &FullSprint{Sprint: spr}

	var er error
	ret.Permissions, er = s.sp.GetBySprintID(p.Ctx, p.Tx, spr.ID, p.Params.Get("spermission", nil, p.Logger), p.Logger)
	if er != nil {
		return nil, er
	}
	if ok, msg := CheckPermissions(util.KeySprint, ret.Permissions.ToPermissions(), p.Accounts, tf, sf); !ok {
		return nil, errors.New(msg)
	}

	funcs := []func() error{
		func() error {
			var err error
			ret.Histories, err = s.sh.GetBySprintID(p.Ctx, p.Tx, spr.ID, p.Params.Get("shistory", nil, p.Logger), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Members, ret.Self, err = s.membersSprint(p, spr.ID)
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
			ret.Estimates, err = s.e.GetBySprintID(p.Ctx, p.Tx, &spr.ID, p.Params.Get(util.KeyEstimate, nil, p.Logger), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Standups, err = s.u.GetBySprintID(p.Ctx, p.Tx, &spr.ID, p.Params.Get(util.KeyStandup, nil, p.Logger), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Retros, err = s.r.GetBySprintID(p.Ctx, p.Tx, &spr.ID, p.Params.Get(util.KeyRetro, nil, p.Logger), p.Logger)
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

func (s *Service) membersSprint(p *LoadParams, sprintID uuid.UUID) (smember.SprintMembers, *smember.SprintMember, error) {
	params := p.Params.Get("smember", nil, p.Logger)
	members, err := s.sm.GetBySprintID(p.Ctx, p.Tx, sprintID, params, p.Logger)
	if err != nil {
		return nil, nil, err
	}
	self := members.Get(sprintID, p.Profile.ID)
	if self == nil && p.Profile.Name != "" {
		err = s.us.CreateIfNeeded(p.Ctx, p.Profile.ID, p.Profile.Name, p.Tx, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		_, err = s.sm.Register(p.Ctx, sprintID, p.Profile.ID, p.Profile.Name, enum.MemberStatusMember, p.Tx, s.a, s.send, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		members, err = s.sm.GetBySprintID(p.Ctx, p.Tx, sprintID, params, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		self = members.Get(sprintID, p.Profile.ID)
	}
	return members, self, nil
}
