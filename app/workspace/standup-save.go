package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) CreateStandup(
	ctx context.Context, id uuid.UUID, title string, user uuid.UUID, name string, picture string, teamID *uuid.UUID, sprintID *uuid.UUID, logger util.Logger,
) (*standup.Standup, *umember.StandupMember, error) {
	slug := s.u.Slugify(ctx, id, title, "", s.uh, nil, logger)
	model := &standup.Standup{
		ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, TeamID: teamID, SprintID: sprintID, Created: time.Now(),
	}
	err := s.u.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save standup")
	}

	err = s.a.Post(ctx, util.KeyStandup, model.ID, user, action.ActCreate, util.ValueMap{"payload": model}, nil, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save standup activity")
	}

	member, err := s.um.Register(ctx, model.ID, user, name, picture, enum.MemberStatusOwner, nil, s.a, s.send, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save standup owner")
	}

	return model, member, nil
}

func (s *Service) SaveStandup(ctx context.Context, u *standup.Standup, user uuid.UUID, tx *sqlx.Tx, logger util.Logger) (*standup.Standup, error) {
	curr, err := s.u.Get(ctx, tx, u.ID, logger)
	if err != nil {
		return nil, err
	}
	if curr == nil {
		return nil, errors.Errorf("no existing standup found with id [%s]", u.ID.String())
	}

	if curr.Slug != u.Slug {
		_ = s.uh.DeleteWhere(ctx, tx, "standup_id = $1 and slug = $2", -1, logger, u.ID, curr.Slug)
		hist := &uhistory.StandupHistory{Slug: curr.Slug, StandupID: u.ID, StandupName: u.TitleString(), Created: time.Now()}
		err = s.uh.Create(ctx, tx, logger, hist)
		if err != nil {
			return nil, err
		}
	}

	err = s.u.Save(ctx, tx, logger, u)
	if err != nil {
		return nil, err
	}

	act := action.NewAction(enum.ModelServiceStandup, u.ID, user, "update", util.ValueMap{}, "")
	err = s.a.Save(ctx, tx, logger, act)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) DeleteStandup(ctx context.Context, fu *FullStandup, logger util.Logger) error {
	tx, err := s.db.StartTransaction(logger)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	for _, r := range fu.Reports {
		err = s.rt.Delete(ctx, tx, r.ID, logger)
		if err != nil {
			return err
		}
	}

	for _, h := range fu.Histories {
		err = s.uh.Delete(ctx, tx, h.Slug, logger)
		if err != nil {
			return err
		}
	}

	for _, c := range fu.Comments {
		err = s.c.Delete(ctx, tx, c.ID, logger)
		if err != nil {
			return err
		}
	}

	for _, p := range fu.Permissions {
		err = s.up.Delete(ctx, tx, p.StandupID, p.Key, p.Value, logger)
		if err != nil {
			return err
		}
	}

	for _, m := range fu.Members {
		err = s.um.Delete(ctx, tx, m.StandupID, m.UserID, logger)
		if err != nil {
			return err
		}
	}

	err = s.u.Delete(ctx, tx, fu.Standup.ID, logger)
	if err != nil {
		return err
	}

	err = s.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActReset, nil, nil, logger)
	if err != nil {
		return err
	}

	return tx.Commit()
}
