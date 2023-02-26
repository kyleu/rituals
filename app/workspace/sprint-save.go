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
	ctx context.Context, id uuid.UUID, title string, user uuid.UUID, name string, picture string, teamID *uuid.UUID, logger util.Logger,
) (*sprint.Sprint, *smember.SprintMember, error) {
	slug := s.s.Slugify(ctx, id, title, "", s.sh, nil, logger)
	model := &sprint.Sprint{ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, TeamID: teamID, Created: time.Now()}
	err := s.s.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save sprint")
	}
	member, err := s.sm.Register(ctx, model.ID, user, name, picture, enum.MemberStatusOwner, nil, s.a, s.send, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save sprint owner")
	}
	if teamID != nil {
		msg := map[string]any{"type": enum.ModelServiceSprint, "id": model.ID, "title": model.Title, "path": model.PublicWebPath(), "icon": model.IconSafe()}
		err = s.send(enum.ModelServiceTeam, *teamID, action.ActChildAdd, msg, &user, logger)
		if err != nil {
			return nil, nil, errors.Wrap(err, "unable to notify team")
		}
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
	}

	err = s.s.Save(ctx, tx, logger, spr)
	if err != nil {
		return nil, err
	}

	act := action.NewAction(enum.ModelServiceSprint, spr.ID, user, "update", util.ValueMap{}, "")
	err = s.a.Save(ctx, tx, logger, act)
	if err != nil {
		return nil, err
	}

	return spr, nil
}

func (s *Service) DeleteSprint(ctx context.Context, fs *FullSprint, logger util.Logger) error {
	tx, err := s.db.StartTransaction(logger)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	for _, e := range fs.Estimates {
		e.SprintID = nil
		err = s.e.Update(ctx, tx, e, logger)
		if err != nil {
			return err
		}
	}

	for _, u := range fs.Standups {
		u.SprintID = nil
		err = s.u.Update(ctx, tx, u, logger)
		if err != nil {
			return err
		}
	}

	for _, r := range fs.Retros {
		r.SprintID = nil
		err = s.r.Update(ctx, tx, r, logger)
		if err != nil {
			return err
		}
	}

	for _, h := range fs.Histories {
		err = s.sh.Delete(ctx, tx, h.Slug, logger)
		if err != nil {
			return err
		}
	}

	for _, c := range fs.Comments.GetByModel(enum.ModelServiceSprint, fs.Sprint.ID) {
		err = s.c.Delete(ctx, tx, c.ID, logger)
		if err != nil {
			return err
		}
	}

	for _, p := range fs.Permissions {
		err = s.sp.Delete(ctx, tx, p.SprintID, p.Key, p.Value, logger)
		if err != nil {
			return err
		}
	}

	for _, m := range fs.Members {
		err = s.sm.Delete(ctx, tx, m.SprintID, m.UserID, logger)
		if err != nil {
			return err
		}
	}

	err = s.s.Delete(ctx, tx, fs.Sprint.ID, logger)
	if err != nil {
		return err
	}

	err = s.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActReset, nil, nil, logger)
	if err != nil {
		return err
	}

	return tx.Commit()
}
