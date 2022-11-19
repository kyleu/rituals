package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/util"
)

var DefaultEstimateChoices = []string{"0", "1", "2", "3", "5", "8", "13", "100"}

func (s *Service) CreateEstimate(
	ctx context.Context, id uuid.UUID, title string, user uuid.UUID, name string, teamID *uuid.UUID, sprintID *uuid.UUID, logger util.Logger,
) (*estimate.Estimate, *emember.EstimateMember, error) {
	slug := s.t.Slugify(ctx, id, title, "", s.th, nil, logger)
	model := &estimate.Estimate{
		ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, Owner: user, TeamID: teamID, SprintID: sprintID, Created: time.Now(),
	}
	err := s.e.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save estimate")
	}

	err = s.a.Post(ctx, util.KeyEstimate, model.ID, user, action.ActCreate, nil, nil, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save estimate activity")
	}

	member, err := s.em.Register(ctx, model.ID, user, name, enum.MemberStatusOwner, nil, s.a, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save estimate owner")
	}

	return model, member, nil
}

func (s *Service) SaveEstimate(ctx context.Context, e *estimate.Estimate, user uuid.UUID, tx *sqlx.Tx, logger util.Logger) (*estimate.Estimate, error) {
	curr, err := s.e.Get(ctx, tx, e.ID, logger)
	if err != nil {
		return nil, err
	}
	if curr == nil {
		return nil, errors.Errorf("no existing estimate found with id [%s]", e.ID.String())
	}

	if curr.Slug != e.Slug {
		_ = s.eh.DeleteWhere(ctx, tx, "estimate_id = $1 and slug = $2", -1, logger, e.ID, curr.Slug)
		hist := &ehistory.EstimateHistory{Slug: curr.Slug, EstimateID: e.ID, EstimateName: e.TitleString(), Created: time.Now()}
		err = s.eh.Create(ctx, tx, logger, hist)
		if err != nil {
			return nil, err
		}
		curr.Slug = e.Slug
	}

	if len(e.Choices) == 0 {
		e.Choices = DefaultEstimateChoices
	}

	err = s.e.Update(ctx, tx, e, logger)
	if err != nil {
		return nil, err
	}

	act := action.NewAction(enum.ModelServiceEstimate, e.ID, user, "update", util.ValueMap{}, "")
	err = s.a.Save(ctx, tx, logger, act)
	if err != nil {
		return nil, err
	}

	return curr, nil
}
