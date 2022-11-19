package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) CreateSprint(
	ctx context.Context, id uuid.UUID, title string, user uuid.UUID, name string, teamID *uuid.UUID, logger util.Logger,
) (*sprint.Sprint, *smember.SprintMember, error) {
	slug := s.s.Slugify(ctx, id, title, "", s.sh, nil, logger)
	model := &sprint.Sprint{ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, Owner: user, TeamID: teamID, Created: time.Now()}
	err := s.s.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save sprint")
	}

	err = s.a.Post(ctx, util.KeySprint, model.ID, user, action.ActCreate, nil, nil, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save sprint activity")
	}

	member, err := s.sm.Register(ctx, model.ID, user, name, enum.MemberStatusOwner, nil, s.a, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save sprint owner")
	}

	return model, member, nil
}

func (s *Service) SaveSprint(ctx context.Context, spr *sprint.Sprint, user uuid.UUID, tx *sqlx.Tx, logger util.Logger) (*sprint.Sprint, error) {
	curr, err := s.s.Get(ctx, tx, spr.ID, logger)
	if err != nil {
		return nil, err
	}
	if curr == nil {
		return nil, errors.Errorf("no existing sprint found with id [%s]", spr.ID.String())
	}

	if curr.Slug != spr.Slug {
		_ = s.sh.DeleteWhere(ctx, tx, "sprint_id = $1 and slug = $2", -1, logger, spr.ID, curr.Slug)
		hist := &shistory.SprintHistory{Slug: curr.Slug, SprintID: spr.ID, SprintName: spr.TitleString(), Created: time.Now()}
		err = s.sh.Create(ctx, tx, logger, hist)
		if err != nil {
			return nil, err
		}
		curr.Slug = spr.Slug
	}

	err = s.s.Update(ctx, tx, spr, logger)
	if err != nil {
		return nil, err
	}

	act := action.NewAction(enum.ModelServiceSprint, spr.ID, user, "update", util.ValueMap{}, "")
	err = s.a.Save(ctx, tx, logger, act)
	if err != nil {
		return nil, err
	}

	return curr, nil
}
