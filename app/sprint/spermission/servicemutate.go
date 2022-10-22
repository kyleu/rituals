// Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*SprintPermission) error {
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

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *SprintPermission, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.SprintID, model.K, model.V, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original permission [%s]", model.String())
	}
	model.Created = curr.Created
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"sprint_id\" = $6 and \"k\" = $7 and \"v\" = $8", "")
	data := model.ToData()
	data = append(data, model.SprintID, model.K, model.V)
	_, err = s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*SprintPermission) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"sprint_id", "k", "v"}, columnsQuoted, "")
	var data []any
	for _, model := range models {
		data = append(data, model.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, sprintID uuid.UUID, k string, v string, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0))
	_, err := s.db.Delete(ctx, q, tx, 1, logger, sprintID, k, v)
	return err
}

func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	q := database.SQLDelete(tableQuoted, wc)
	_, err := s.db.Delete(ctx, q, tx, expected, logger, values...)
	return err
}
