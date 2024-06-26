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
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/util"
)

type FullTeam struct {
	Team        *team.Team                  `json:"team"`
	Histories   thistory.TeamHistories      `json:"histories,omitempty"`
	Members     tmember.TeamMembers         `json:"members,omitempty"`
	UtilMembers member.Members              `json:"-"`
	Self        *tmember.TeamMember         `json:"self,omitempty"`
	Permissions tpermission.TeamPermissions `json:"permissions,omitempty"`
	Sprints     sprint.Sprints              `json:"sprints,omitempty"`
	Estimates   estimate.Estimates          `json:"estimates,omitempty"`
	Standups    standup.Standups            `json:"standups,omitempty"`
	Retros      retro.Retros                `json:"retros,omitempty"`
	Comments    comment.Comments            `json:"comments,omitempty"`
	Actions     action.Actions              `json:"actions,omitempty"`
	Registered  bool                        `json:"registered,omitempty"`
}

func (f *FullTeam) Admin() bool {
	return f.Self.Role == enum.MemberStatusOwner
}

func (s *Service) LoadTeam(p *LoadParams) (*FullTeam, error) {
	t, err := s.t.GetBySlug(p.Ctx, p.Tx, p.Slug, p.Logger)
	if err != nil {
		if hist, _ := s.th.Get(p.Ctx, p.Tx, p.Slug, p.Logger); hist != nil {
			t, err = s.t.Get(p.Ctx, p.Tx, hist.TeamID, p.Logger)
			if err != nil {
				return nil, errors.Errorf("no team found with slug [%s]", p.Slug)
			}
		}
	}
	if t == nil {
		id := util.UUIDFromString(p.Slug)
		if id == nil {
			return nil, errors.Errorf("no team found with slug [%s]", p.Slug)
		}
		t, err = s.t.Get(p.Ctx, p.Tx, *id, p.Logger)
		if err != nil {
			return nil, errors.Errorf("no team found with id [%s]", p.Slug)
		}
	}

	return s.loadFullTeam(p, t)
}

func (s *Service) loadFullTeam(p *LoadParams, t *team.Team) (*FullTeam, error) {
	ret := &FullTeam{Team: t}

	var er error
	ret.Permissions, er = s.tp.GetByTeamID(p.Ctx, p.Tx, t.ID, p.Params.Get("tpermission", nil, p.Logger).Sanitize("tpermission"), p.Logger)
	if er != nil {
		return nil, er
	}
	if ok, msg := CheckPermissions(util.KeyTeam, ret.Permissions.ToPermissions(), p.Accounts, nil, nil, nil, nil); !ok {
		return nil, errors.New(msg)
	}
	funcs := []func() error{
		func() error {
			var err error
			ret.Histories, err = s.th.GetByTeamID(p.Ctx, p.Tx, t.ID, p.Params.Get("thistory", nil, p.Logger).Sanitize("thistory"), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Members, ret.Self, ret.Registered, err = s.membersTeam(p, t.ID)
			online := s.online(util.KeyTeam + ":" + t.ID.String())
			ret.UtilMembers = ret.Members.ToMembers(online)
			return err
		},
		func() error {
			var err error
			ret.Sprints, err = s.s.GetByTeamID(p.Ctx, p.Tx, &t.ID, p.Params.Get(util.KeySprint, nil, p.Logger).Sanitize(util.KeySprint), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Estimates, err = s.e.GetByTeamID(p.Ctx, p.Tx, &t.ID, p.Params.Get(util.KeyEstimate, nil, p.Logger).Sanitize(util.KeyEstimate), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Standups, err = s.u.GetByTeamID(p.Ctx, p.Tx, &t.ID, p.Params.Get(util.KeyStandup, nil, p.Logger).Sanitize(util.KeyStandup), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Retros, err = s.r.GetByTeamID(p.Ctx, p.Tx, &t.ID, p.Params.Get(util.KeyRetro, nil, p.Logger).Sanitize(util.KeyRetro), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Actions, err = s.a.GetByModels(p.Ctx, p.Tx, p.Logger, enum.ModelServiceTeam, ret.Team.ID)
			return err
		},
	}
	_, errs := util.AsyncCollect(funcs, func(f func() error) (any, error) {
		return nil, f()
	})
	if len(errs) > 0 {
		return nil, util.ErrorMerge(errs...)
	}

	args := make([]any, 0, (len(ret.Sprints)*2)+(len(ret.Estimates)*2)+(len(ret.Standups)*2)+(len(ret.Retros)*2)+2)
	args = append(args, util.KeyTeam, t.ID)
	for _, x := range ret.Sprints {
		args = append(args, util.KeySprint, x.ID)
	}
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

func (s *Service) membersTeam(p *LoadParams, teamID uuid.UUID) (tmember.TeamMembers, *tmember.TeamMember, bool, error) {
	params := p.Params.Get("tmember", nil, p.Logger).Sanitize("tmember")
	members, err := s.tm.GetByTeamID(p.Ctx, p.Tx, teamID, params, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	err = s.us.CreateIfNeeded(p.Ctx, p.Profile.ID, p.Profile.Name, p.Tx, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	if self := members.Get(teamID, p.Profile.ID); self != nil {
		return members, self, false, nil
	}
	role := enum.MemberStatusMember
	if len(members) == 0 {
		role = enum.MemberStatusOwner
	}
	_, err = s.tm.Register(p.Ctx, teamID, p.Profile.ID, p.Profile.Name, p.Accounts.Image(), role, nil, s.a, s.send, s.us, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	members, err = s.tm.GetByTeamID(p.Ctx, p.Tx, teamID, params, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	return members, members.Get(teamID, p.Profile.ID), true, nil
}
