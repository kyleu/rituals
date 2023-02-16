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

func (s *Service) LoadEstimate(p *LoadParams, tf func() (team.Teams, error), sf func() (sprint.Sprints, error)) (*FullEstimate, error) {
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
	return s.loadFullEstimate(p, e, tf, sf)
}

func (s *Service) loadFullEstimate(
	p *LoadParams, e *estimate.Estimate, tf func() (team.Teams, error), sf func() (sprint.Sprints, error),
) (*FullEstimate, error) {
	ret := &FullEstimate{Estimate: e}

	var er error
	ret.Stories, er = s.st.GetByEstimateID(p.Ctx, p.Tx, e.ID, p.Params.Get(util.KeyStory, nil, p.Logger).Sanitize(util.KeyStory), p.Logger)
	if er != nil {
		return nil, er
	}
	ret.Stories.Sort()

	ret.Permissions, er = s.ep.GetByEstimateID(p.Ctx, p.Tx, e.ID, p.Params.Get("epermission", nil, p.Logger).Sanitize("epermission"), p.Logger)
	if er != nil {
		return nil, er
	}
	if ok, msg := CheckPermissions(util.KeyEstimate, ret.Permissions.ToPermissions(), p.Accounts, e.TeamID, tf, e.SprintID, sf); !ok {
		return nil, errors.New(msg)
	}

	funcs := []func() error{
		func() error {
			var err error
			ret.Histories, err = s.eh.GetByEstimateID(p.Ctx, p.Tx, e.ID, p.Params.Get("ehistory", nil, p.Logger).Sanitize("ehistory"), p.Logger)
			return err
		},
		func() error {
			ret.Members, ret.Self, er = s.membersEstimate(p, e.ID)
			online := s.online(util.KeyEstimate + ":" + e.ID.String())
			ret.UtilMembers = ret.Members.ToMembers(online)
			return er
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
			prm := p.Params.Get(util.KeyVote, nil, p.Logger).Sanitize(util.KeyVote)
			ret.Votes, err = s.v.GetByStoryIDs(p.Ctx, p.Tx, prm, p.Logger, ret.Stories.IDStrings(false)...)
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
	params := p.Params.Get("emember", nil, p.Logger).Sanitize("emember")
	members, err := s.em.GetByEstimateID(p.Ctx, p.Tx, estimateID, params, p.Logger)
	if err != nil {
		return nil, nil, err
	}
	self := members.Get(estimateID, p.Profile.ID)
	if self == nil && p.Profile.Name != "" {
		err = s.us.CreateIfNeeded(p.Ctx, p.Profile.ID, p.Profile.Name, p.Tx, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		_, err = s.em.Register(p.Ctx, estimateID, p.Profile.ID, p.Profile.Name, enum.MemberStatusMember, p.Tx, s.a, s.send, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		members, err = s.em.GetByEstimateID(p.Ctx, p.Tx, estimateID, params, p.Logger)
		if err != nil {
			return nil, nil, err
		}
		self = members.Get(estimateID, p.Profile.ID)
	}
	return members, self, nil
}
