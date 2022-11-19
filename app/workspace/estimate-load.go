package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type FullEstimate struct {
	Estimate    *estimate.Estimate              `json:"estimate"`
	Histories   ehistory.EstimateHistories      `json:"histories,omitempty"`
	Members     emember.EstimateMembers         `json:"members,omitempty"`
	Self        *emember.EstimateMember         `json:"self,omitempty"`
	Permissions epermission.EstimatePermissions `json:"permissions,omitempty"`
	Team        *team.Team                      `json:"team,omitempty"`
	Sprint      *sprint.Sprint                  `json:"sprint,omitempty"`
	Stories     story.Stories                   `json:"stories,omitempty"`
	Votes       vote.Votes                      `json:"votes,omitempty"`
	Comments    comment.Comments                `json:"comments,omitempty"`
	Actions     action.Actions                  `json:"actions,omitempty"`
}

func (s *Service) LoadEstimate(
	ctx context.Context, slug string, user uuid.UUID, tx *sqlx.Tx, params filter.ParamSet, logger util.Logger,
) (*FullEstimate, error) {
	e, err := s.e.GetBySlug(ctx, tx, slug, logger)
	if err != nil {
		if hist, _ := s.eh.Get(ctx, tx, slug, logger); hist != nil {
			e, err = s.e.Get(ctx, tx, hist.EstimateID, logger)
			if err != nil {
				return nil, errors.Errorf("no estimate found with slug [%s]", slug)
			}
		}
	}
	if e == nil {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no estimate found with slug [%s]", slug)
		}
		e, err = s.e.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no estimate found with id [%s]", slug)
		}
	}
	ret := &FullEstimate{Estimate: e}

	ret.Histories, err = s.eh.GetByEstimateID(ctx, tx, e.ID, params.Get("ehistory", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Members, err = s.em.GetByEstimateID(ctx, tx, e.ID, params.Get("emember", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Self = ret.Members.Get(e.ID, user)
	ret.Permissions, err = s.ep.GetByEstimateID(ctx, tx, e.ID, params.Get("epermission", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	if e.TeamID != nil {
		ret.Team, err = s.t.Get(ctx, tx, *e.TeamID, logger)
		if err != nil {
			return nil, err
		}
	}
	if e.SprintID != nil {
		ret.Sprint, err = s.s.Get(ctx, tx, *e.SprintID, logger)
		if err != nil {
			return nil, err
		}
	}

	ret.Stories, err = s.st.GetByEstimateID(ctx, tx, e.ID, params.Get(util.KeyStory, nil, logger), logger)
	if err != nil {
		return nil, err
	}

	ret.Votes, err = s.v.GetByStoryIDs(ctx, tx, params.Get(util.KeyVote, nil, logger), logger, ret.Stories.IDStrings(false)...)
	if err != nil {
		return nil, err
	}

	args := make([]any, 0, (len(ret.Stories)*2)+2)
	args = append(args, util.KeyEstimate, e.ID)
	for _, str := range ret.Stories {
		args = append(args, util.KeyStory, str.ID)
	}

	ret.Comments, err = s.c.GetByModels(ctx, tx, logger, args...)
	if err != nil {
		return nil, err
	}
	ret.Actions, err = s.a.GetByModels(ctx, tx, logger, enum.ModelServiceEstimate, ret.Estimate.ID)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
