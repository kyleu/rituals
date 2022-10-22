// Content managed by Project Forge, see [projectforge.md] for details.
package shistory

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*SprintHistory) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
	}
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), "")
	vals := make([]any, 0, len(models)*len(columnsQuoted))
	for _, arg := range models {
		vals = append(vals, arg.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *SprintHistory, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.Slug, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original history [%s]", model.String())
	}
	model.Created = curr.Created
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"slug\" = $5", "")
	data := model.ToData()
	data = append(data, model.Slug)
	_, err = s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*SprintHistory) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"slug"}, columnsQuoted, "")
	var data []any
	for _, model := range models {
		data = append(data, model.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, slug string, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0))
	_, err := s.db.Delete(ctx, q, tx, 1, logger, slug)
	return err
}

func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	q := database.SQLDelete(tableQuoted, wc)
	_, err := s.db.Delete(ctx, q, tx, expected, logger, values...)
	return err
}
