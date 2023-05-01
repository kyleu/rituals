package workspace

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type FullStandup struct {
	Standup     *standup.Standup               `json:"standup"`
	Histories   uhistory.StandupHistories      `json:"histories,omitempty"`
	Members     umember.StandupMembers         `json:"members,omitempty"`
	UtilMembers util.Members                   `json:"-"`
	Self        *umember.StandupMember         `json:"self,omitempty"`
	Permissions upermission.StandupPermissions `json:"permissions,omitempty"`
	Team        *team.Team                     `json:"team,omitempty"`
	Sprint      *sprint.Sprint                 `json:"sprint,omitempty"`
	Reports     report.Reports                 `json:"reports,omitempty"`
	Comments    comment.Comments               `json:"comments,omitempty"`
	Actions     action.Actions                 `json:"actions,omitempty"`
	Registered  bool                           `json:"registered,omitempty"`
}

func (f *FullStandup) Admin() bool {
	return f.Self.Role == enum.MemberStatusOwner
}

func (s *Service) LoadStandup(p *LoadParams, tf func() (team.Teams, error), sf func() (sprint.Sprints, error)) (*FullStandup, error) {
	u, err := s.u.GetBySlug(p.Ctx, p.Tx, p.Slug, p.Logger)
	if err != nil {
		if hist, _ := s.uh.Get(p.Ctx, p.Tx, p.Slug, p.Logger); hist != nil {
			u, err = s.u.Get(p.Ctx, p.Tx, hist.StandupID, p.Logger)
			if err != nil {
				return nil, errors.Errorf("no standup found with slug [%s]", p.Slug)
			}
		}
	}
	if u == nil {
		id := util.UUIDFromString(p.Slug)
		if id == nil {
			return nil, errors.Errorf("no standup found with slug [%s]", p.Slug)
		}
		u, err = s.u.Get(p.Ctx, p.Tx, *id, p.Logger)
		if err != nil {
			return nil, errors.Errorf("no standup found with id [%s]", p.Slug)
		}
	}
	return s.loadFullStandup(p, u, tf, sf)
}

func (s *Service) loadFullStandup(p *LoadParams, u *standup.Standup, tf func() (team.Teams, error), sf func() (sprint.Sprints, error)) (*FullStandup, error) {
	ret := &FullStandup{Standup: u}

	var er error
	ret.Permissions, er = s.up.GetByStandupID(p.Ctx, p.Tx, u.ID, p.Params.Get("upermission", nil, p.Logger).Sanitize("upermission"), p.Logger)
	if er != nil {
		return nil, er
	}
	if ok, msg := CheckPermissions(util.KeyStandup, ret.Permissions.ToPermissions(), p.Accounts, u.TeamID, tf, u.SprintID, sf); !ok {
		return nil, errors.New(msg)
	}

	funcs := []func() error{
		func() error {
			var err error
			ret.Histories, err = s.uh.GetByStandupID(p.Ctx, p.Tx, u.ID, p.Params.Get("uhistory", nil, p.Logger).Sanitize("uhistory"), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Members, ret.Self, ret.Registered, err = s.membersStandup(p, u.ID)
			online := s.online(util.KeyStandup + ":" + u.ID.String())
			ret.UtilMembers = ret.Members.ToMembers(online)
			return err
		},
		func() error {
			var err error
			if u.TeamID != nil {
				ret.Team, err = s.t.Get(p.Ctx, p.Tx, *u.TeamID, p.Logger)
			}
			return err
		},
		func() error {
			var err error
			if u.SprintID != nil {
				ret.Sprint, err = s.s.Get(p.Ctx, p.Tx, *u.SprintID, p.Logger)
			}
			return err
		},
		func() error {
			var err error
			ret.Reports, err = s.rt.GetByStandupID(p.Ctx, p.Tx, u.ID, p.Params.Get(util.KeyReport, nil, p.Logger).Sanitize(util.KeyReport), p.Logger)
			if err != nil {
				return err
			}
			args := make([]any, 0, (len(ret.Reports)*2)+2)
			args = append(args, util.KeyStandup, u.ID)
			for _, rpt := range ret.Reports {
				args = append(args, util.KeyReport, rpt.ID)
			}
			ret.Comments, err = s.c.GetByModels(p.Ctx, p.Tx, p.Logger, args...)
			return err
		},
		func() error {
			var err error
			ret.Actions, err = s.a.GetByModels(p.Ctx, p.Tx, p.Logger, enum.ModelServiceStandup, ret.Standup.ID)
			return err
		},
	}
	_, errs := util.AsyncCollect(funcs, func(f func() error) (any, error) {
		return nil, f()
	})
	if len(errs) > 0 {
		return nil, util.ErrorMerge(errs...)
	}

	return ret, nil
}

func (s *Service) membersStandup(p *LoadParams, standupID uuid.UUID) (umember.StandupMembers, *umember.StandupMember, bool, error) {
	params := p.Params.Get("umember", nil, p.Logger).Sanitize("umember")
	members, err := s.um.GetByStandupID(p.Ctx, p.Tx, standupID, params, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	if self := members.Get(standupID, p.Profile.ID); self != nil {
		return members, self, false, nil
	}
	err = s.us.CreateIfNeeded(p.Ctx, p.Profile.ID, p.Profile.Name, p.Tx, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	role := enum.MemberStatusMember
	if len(members) == 0 {
		role = enum.MemberStatusOwner
	}
	_, err = s.um.Register(p.Ctx, standupID, p.Profile.ID, p.Profile.Name, p.Accounts.Image(), role, nil, s.a, s.send, s.us, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	members, err = s.um.GetByStandupID(p.Ctx, p.Tx, standupID, params, p.Logger)
	if err != nil {
		return nil, nil, false, err
	}
	return members, members.Get(standupID, p.Profile.ID), true, nil
}
