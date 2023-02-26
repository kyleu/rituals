package workspace

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

var RetroDefaultCategories = []string{"good", "bad"}

type FullRetro struct {
	Retro       *retro.Retro                 `json:"retro"`
	Histories   rhistory.RetroHistories      `json:"histories,omitempty"`
	Members     rmember.RetroMembers         `json:"members,omitempty"`
	UtilMembers util.Members                 `json:"-"`
	Self        *rmember.RetroMember         `json:"self,omitempty"`
	Permissions rpermission.RetroPermissions `json:"permissions,omitempty"`
	Team        *team.Team                   `json:"team,omitempty"`
	Sprint      *sprint.Sprint               `json:"sprint,omitempty"`
	Feedbacks   feedback.Feedbacks           `json:"feedbacks,omitempty"`
	Comments    comment.Comments             `json:"comments,omitempty"`
	Actions     action.Actions               `json:"actions,omitempty"`
}

func (f *FullRetro) Admin() bool {
	return f.Self.Role == enum.MemberStatusOwner
}

func (s *Service) LoadRetro(p *LoadParams, tf func() (team.Teams, error), sf func() (sprint.Sprints, error)) (*FullRetro, error) {
	r, err := s.r.GetBySlug(p.Ctx, p.Tx, p.Slug, p.Logger)
	if err != nil {
		if hist, _ := s.rh.Get(p.Ctx, p.Tx, p.Slug, p.Logger); hist != nil {
			r, err = s.r.Get(p.Ctx, p.Tx, hist.RetroID, p.Logger)
			if err != nil {
				return nil, errors.Errorf("no retro found with slug [%s]", p.Slug)
			}
		}
	}
	if r == nil {
		id := util.UUIDFromString(p.Slug)
		if id == nil {
			return nil, errors.Errorf("no retro found with slug [%s]", p.Slug)
		}
		r, err = s.r.Get(p.Ctx, p.Tx, *id, p.Logger)
		if err != nil {
			return nil, errors.Errorf("no retro found with id [%s]", p.Slug)
		}
	}
	return s.loadFullRetro(p, r, tf, sf)
}

func (s *Service) loadFullRetro(p *LoadParams, r *retro.Retro, tf func() (team.Teams, error), sf func() (sprint.Sprints, error)) (*FullRetro, error) {
	ret := &FullRetro{Retro: r}

	var er error
	ret.Permissions, er = s.rp.GetByRetroID(p.Ctx, p.Tx, r.ID, p.Params.Get("rpermission", nil, p.Logger).Sanitize("rpermission"), p.Logger)
	if er != nil {
		return nil, er
	}
	if ok, msg := CheckPermissions(util.KeyRetro, ret.Permissions.ToPermissions(), p.Accounts, r.TeamID, tf, r.SprintID, sf); !ok {
		return nil, errors.New(msg)
	}

	funcs := []func() error{
		func() error {
			var err error
			ret.Histories, err = s.rh.GetByRetroID(p.Ctx, p.Tx, r.ID, p.Params.Get("rhistory", nil, p.Logger).Sanitize("rhistory"), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Members, ret.Self, err = s.membersRetro(p, r.ID)
			online := s.online(util.KeyRetro + ":" + r.ID.String())
			ret.UtilMembers = ret.Members.ToMembers(online)
			return err
		},
		func() error {
			var err error
			if r.TeamID != nil {
				ret.Team, err = s.t.Get(p.Ctx, p.Tx, *r.TeamID, p.Logger)
			}
			return err
		},
		func() error {
			var err error
			if r.SprintID != nil {
				ret.Sprint, err = s.s.Get(p.Ctx, p.Tx, *r.SprintID, p.Logger)
			}
			return err
		},
		func() error {
			var err error
			ret.Feedbacks, err = s.f.GetByRetroID(p.Ctx, p.Tx, r.ID, p.Params.Get("feedback", nil, p.Logger).Sanitize("feedback"), p.Logger)
			if err != nil {
				return err
			}
			args := make([]any, 0, (len(ret.Feedbacks)*2)+2)
			args = append(args, util.KeyRetro, r.ID)
			for _, f := range ret.Feedbacks {
				args = append(args, util.KeyFeedback, f.ID)
			}
			ret.Comments, err = s.c.GetByModels(p.Ctx, p.Tx, p.Logger, args...)
			return err
		},
		func() error {
			var err error
			ret.Actions, err = s.a.GetByModels(p.Ctx, p.Tx, p.Logger, enum.ModelServiceRetro, ret.Retro.ID)
			return err
		},
		func() error {
			var err error

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

func (s *Service) membersRetro(p *LoadParams, retroID uuid.UUID) (rmember.RetroMembers, *rmember.RetroMember, error) {
	params := p.Params.Get("rmember", nil, p.Logger).Sanitize("rmember")
	members, err := s.rm.GetByRetroID(p.Ctx, p.Tx, retroID, params, p.Logger)
	if err != nil {
		return nil, nil, err
	}
	self := members.Get(retroID, p.Profile.ID)
	if self == nil && p.Profile.Name != "" {
		err = s.us.CreateIfNeeded(p.Ctx, p.Profile.ID, p.Profile.Name, p.Tx, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		role := enum.MemberStatusMember
		if len(members) == 0 {
			role = enum.MemberStatusOwner
		}
		_, err = s.rm.Register(p.Ctx, retroID, p.Profile.ID, p.Profile.Name, p.Accounts.Image(), role, p.Tx, s.a, s.send, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		members, err = s.rm.GetByRetroID(p.Ctx, p.Tx, retroID, params, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		self = members.Get(retroID, p.Profile.ID)
	}
	return members, self, nil
}
