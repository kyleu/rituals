package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) CreateRetro(
	ctx context.Context, id uuid.UUID, title string, user uuid.UUID, name string, picture string, teamID *uuid.UUID, sprintID *uuid.UUID, logger util.Logger,
) (*retro.Retro, *rmember.RetroMember, error) {
	slug := s.r.Slugify(ctx, id, title, "", s.rh, nil, logger)
	model := &retro.Retro{
		ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, TeamID: teamID, SprintID: sprintID, Created: time.Now(),
	}
	err := s.r.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save retro")
	}

	err = s.a.Post(ctx, util.KeyRetro, model.ID, user, action.ActCreate, util.ValueMap{"payload": model}, nil, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save retro activity")
	}

	member, err := s.rm.Register(ctx, model.ID, user, name, picture, enum.MemberStatusOwner, nil, s.a, s.send, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save retro owner")
	}

	return model, member, nil
}

func (s *Service) SaveRetro(ctx context.Context, r *retro.Retro, user uuid.UUID, tx *sqlx.Tx, logger util.Logger) (*retro.Retro, error) {
	curr, err := s.r.Get(ctx, tx, r.ID, logger)
	if err != nil {
		return nil, err
	}
	if curr == nil {
		return nil, errors.Errorf("no existing retro found with id [%s]", r.ID.String())
	}

	if curr.Slug != r.Slug {
		_ = s.rh.DeleteWhere(ctx, tx, "retro_id = $1 and slug = $2", -1, logger, r.ID, curr.Slug)
		hist := &rhistory.RetroHistory{Slug: curr.Slug, RetroID: r.ID, RetroName: r.TitleString(), Created: time.Now()}
		err = s.rh.Create(ctx, tx, logger, hist)
		if err != nil {
			return nil, err
		}
	}

	if len(r.Categories) == 0 {
		r.Categories = RetroDefaultCategories
	}

	err = s.r.Save(ctx, tx, logger, r)
	if err != nil {
		return nil, err
	}

	act := action.NewAction(enum.ModelServiceRetro, r.ID, user, "update", util.ValueMap{}, "")
	err = s.a.Save(ctx, tx, logger, act)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Service) DeleteRetro(ctx context.Context, fr *FullRetro, logger util.Logger) error {
	tx, err := s.db.StartTransaction(logger)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	for _, f := range fr.Feedbacks {
		err = s.f.Delete(ctx, tx, f.ID, logger)
		if err != nil {
			return err
		}
	}

	for _, h := range fr.Histories {
		err = s.rh.Delete(ctx, tx, h.Slug, logger)
		if err != nil {
			return err
		}
	}

	for _, c := range fr.Comments {
		err = s.c.Delete(ctx, tx, c.ID, logger)
		if err != nil {
			return err
		}
	}

	for _, p := range fr.Permissions {
		err = s.rp.Delete(ctx, tx, p.RetroID, p.Key, p.Value, logger)
		if err != nil {
			return err
		}
	}

	for _, m := range fr.Members {
		err = s.rm.Delete(ctx, tx, m.RetroID, m.UserID, logger)
		if err != nil {
			return err
		}
	}

	err = s.r.Delete(ctx, tx, fr.Retro.ID, logger)
	if err != nil {
		return err
	}

	err = s.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActReset, nil, nil, logger)
	if err != nil {
		return err
	}

	return tx.Commit()
}
