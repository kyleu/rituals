package workspace

import (
	"context"
	"github.com/kyleu/rituals/app/lib/filter"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

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
	Histories   ehistory.EstimateHistories      `json:"histories"`
	Members     emember.EstimateMembers         `json:"members"`
	Permissions epermission.EstimatePermissions `json:"permissions"`
	Team        *team.Team                      `json:"team"`
	Sprint      *sprint.Sprint                  `json:"sprint"`
	Stories     story.Stories                   `json:"stories"`
	Votes       vote.Votes                      `json:"votes"`
}

func (s *Service) CreateEstimate(
	ctx context.Context, id uuid.UUID, slug string, title string, user uuid.UUID, name string, teamID *uuid.UUID, sprintID *uuid.UUID, logger util.Logger,
) (*estimate.Estimate, error) {
	model := &estimate.Estimate{
		ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, Owner: user, TeamID: teamID, SprintID: sprintID, Created: time.Now(),
	}
	err := s.e.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save estimate")
	}
	member := &emember.EstimateMember{EstimateID: model.ID, UserID: user, Name: name, Role: enum.MemberStatusOwner, Created: time.Now()}
	err = s.em.Create(ctx, nil, logger, member)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save estimate member")
	}
	return model, nil
}

func (s *Service) LoadEstimate(ctx context.Context, slug string, user uuid.UUID, tx *sqlx.Tx, params filter.ParamSet, logger util.Logger) (*FullEstimate, error) {
	bySlug, err := s.e.GetBySlug(ctx, tx, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no estimate found with slug [%s]", slug)
		}
		e, err := s.e.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no estimate found with slug [%s]", slug)
		}
		bySlug = estimate.Estimates{e}
	}
	e := bySlug[0]
	ret := &FullEstimate{Estimate: e}

	ret.Histories, err = s.eh.GetByEstimateID(ctx, tx, e.ID, params.Get("ehistory", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	ret.Members, err = s.em.GetByEstimateID(ctx, tx, e.ID, params.Get("emember", nil, logger), logger)
	if err != nil {
		return nil, err
	}

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

	ret.Stories, err = s.st.GetByEstimateID(ctx, tx, e.ID, params.Get("story", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	ret.Votes, err = s.v.GetByStoryIDs(ctx, tx, params.Get("vote", nil, logger), logger, ret.Stories.IDStrings(false)...)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
