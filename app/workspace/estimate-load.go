package workspace

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type FullEstimate struct {
	Estimate    *estimate.Estimate              `json:"estimate"`
	Histories   ehistory.EstimateHistories      `json:"histories,omitempty"`
	Members     emember.EstimateMembers         `json:"members,omitempty"`
	UtilMembers util.Members                    `json:"-"`
	Self        *emember.EstimateMember         `json:"self,omitempty"`
	Permissions epermission.EstimatePermissions `json:"permissions,omitempty"`
	Team        *team.Team                      `json:"team,omitempty"`
	Sprint      *sprint.Sprint                  `json:"sprint,omitempty"`
	Stories     story.Stories                   `json:"stories,omitempty"`
	Votes       vote.Votes                      `json:"votes,omitempty"`
	Comments    comment.Comments                `json:"comments,omitempty"`
	Actions     action.Actions                  `json:"actions,omitempty"`
}

func (s *Service) LoadEstimate(p *LoadParams) (*FullEstimate, error) {
	e, err := s.e.GetBySlug(p.Ctx, p.Tx, p.Slug, p.Logger)
	if err != nil {
		if hist, _ := s.eh.Get(p.Ctx, p.Tx, p.Slug, p.Logger); hist != nil {
			e, err = s.e.Get(p.Ctx, p.Tx, hist.EstimateID, p.Logger)
			if err != nil {
				return nil, errors.Errorf("no estimate found with slug [%s]", p.Slug)
			}
		}
	}
	if e == nil {
		id := util.UUIDFromString(p.Slug)
		if id == nil {
			return nil, errors.Errorf("no estimate found with slug [%s]", p.Slug)
		}
		e, err = s.e.Get(p.Ctx, p.Tx, *id, p.Logger)
		if err != nil {
			return nil, errors.Errorf("no estimate found with id [%s]", p.Slug)
		}
	}
	return s.loadFullEstimate(p, e)
}

func (s *Service) loadFullEstimate(p *LoadParams, e *estimate.Estimate) (*FullEstimate, error) {
	ret := &FullEstimate{Estimate: e}
	var er error
	ret.Stories, er = s.st.GetByEstimateID(p.Ctx, p.Tx, e.ID, p.Params.Get(util.KeyStory, nil, p.Logger), p.Logger)
	if er != nil {
		return nil, er
	}
	ret.Stories.Sort()
	funcs := []func() error{
		func() error {
			var err error
			ret.Histories, err = s.eh.GetByEstimateID(p.Ctx, p.Tx, e.ID, p.Params.Get("ehistory", nil, p.Logger), p.Logger)
			return err
		},
		func() error {
			var err error
			ret.Members, ret.Self, err = s.membersEstimate(p, e.ID)
			online := s.online(util.KeyEstimate + ":" + e.ID.String())
			ret.UtilMembers = ret.Members.ToMembers(online)
			return err
		},
		func() error {
			var err error
			ret.Permissions, err = s.ep.GetByEstimateID(p.Ctx, p.Tx, e.ID, p.Params.Get("epermission", nil, p.Logger), p.Logger)
			return err
		},
		func() error {
			var err error
			if e.TeamID != nil {
				ret.Team, err = s.t.Get(p.Ctx, p.Tx, *e.TeamID, p.Logger)
			}
			return err
		},
		func() error {
			var err error
			if e.SprintID != nil {
				ret.Sprint, err = s.s.Get(p.Ctx, p.Tx, *e.SprintID, p.Logger)
			}
			return err
		},
		func() error {
			var err error
			args := make([]any, 0, (len(ret.Stories)*2)+2)
			args = append(args, util.KeyEstimate, e.ID)
			for _, str := range ret.Stories {
				args = append(args, util.KeyStory, str.ID)
			}
			ret.Comments, err = s.c.GetByModels(p.Ctx, p.Tx, p.Logger, args...)
			return err
		},
		func() error {
			var err error
			ret.Votes, err = s.v.GetByStoryIDs(p.Ctx, p.Tx, p.Params.Get(util.KeyVote, nil, p.Logger), p.Logger, ret.Stories.IDStrings(false)...)
			return err
		},
		func() error {
			var err error
			ret.Actions, err = s.a.GetByModels(p.Ctx, p.Tx, p.Logger, enum.ModelServiceEstimate, ret.Estimate.ID)
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

func (s *Service) membersEstimate(p *LoadParams, estimateID uuid.UUID) (emember.EstimateMembers, *emember.EstimateMember, error) {
	params := p.Params.Get("emember", nil, p.Logger)
	members, err := s.em.GetByEstimateID(p.Ctx, p.Tx, estimateID, params, p.Logger)
	if err != nil {
		return nil, nil, err
	}
	self := members.Get(estimateID, p.UserID)
	if self == nil && p.Username != "" {
		err = s.us.CreateIfNeeded(p.Ctx, p.UserID, p.Username, p.Tx, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		_, err = s.em.Register(p.Ctx, estimateID, p.UserID, p.Username, enum.MemberStatusMember, p.Tx, s.a, s.send, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		members, err = s.em.GetByEstimateID(p.Ctx, p.Tx, estimateID, params, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		self = members.Get(estimateID, p.UserID)
	}
	return members, self, nil
}
