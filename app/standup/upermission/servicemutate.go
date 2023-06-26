// Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*StandupPermission) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *StandupPermission, _ int) {
		model.Created = time.Now()
	})
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), s.db.Placeholder())
	vals := lo.FlatMap(models, func(arg *StandupPermission, _ int) []any {
		return arg.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *StandupPermission, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.StandupID, model.Key, model.Value, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original permission [%s]", model.String())
	}
	model.Created = curr.Created
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"standup_id\" = $6 and \"key\" = $7 and \"value\" = $8", s.db.Placeholder())
	data := model.ToData()
	data = append(data, model.StandupID, model.Key, model.Value)
	_, err = s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*StandupPermission) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *StandupPermission, _ int) {
		model.Created = time.Now()
	})
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"standup_id", "key", "value"}, columnsQuoted, s.db.Placeholder())
	data := lo.FlatMap(models, func(model *StandupPermission, _ int) []any {
		return model.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, standupID uuid.UUID, key string, value string, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0), s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, 1, logger, standupID, key, value)
	return err
}

func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	q := database.SQLDelete(tableQuoted, wc, s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, expected, logger, values...)
	return err
}
