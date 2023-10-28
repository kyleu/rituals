// Package smember - Content managed by Project Forge, see [projectforge.md] for details.
package smember

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*SprintMember) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *SprintMember, _ int) {
		model.Created = util.TimeCurrent()
		model.Updated = util.TimeCurrentP()
	})
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), s.db.Placeholder())
	vals := lo.FlatMap(models, func(arg *SprintMember, _ int) []any {
		return arg.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *SprintMember, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.SprintID, model.UserID, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original member [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.TimeCurrentP()
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"sprint_id\" = $8 and \"user_id\" = $9", s.db.Placeholder())
	data := model.ToData()
	data = append(data, model.SprintID, model.UserID)
	_, err = s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*SprintMember) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *SprintMember, _ int) {
		model.Created = util.TimeCurrent()
		model.Updated = util.TimeCurrentP()
	})
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"sprint_id", "user_id"}, columnsQuoted, s.db.Placeholder())
	data := lo.FlatMap(models, func(model *SprintMember, _ int) []any {
		return model.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) SaveChunked(ctx context.Context, tx *sqlx.Tx, chunkSize int, logger util.Logger, models ...*SprintMember) error {
	for idx, chunk := range lo.Chunk(models, chunkSize) {
		if logger != nil {
			logger.Infof("saving members [%d-%d]", idx*chunkSize, ((idx+1)*chunkSize)-1)
		}
		if err := s.Save(ctx, tx, logger, chunk...); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, sprintID uuid.UUID, userID uuid.UUID, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0), s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, 1, logger, sprintID, userID)
	return err
}

func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	q := database.SQLDelete(tableQuoted, wc, s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, expected, logger, values...)
	return err
}
