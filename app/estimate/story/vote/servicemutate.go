// Content managed by Project Forge, see [projectforge.md] for details.
package vote

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

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Vote) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *Vote, _ int) {
		model.Created = time.Now()
		model.Updated = util.NowPointer()
	})
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), s.db.Placeholder())
	vals := lo.FlatMap(models, func(arg *Vote, _ int) []any {
		return arg.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Vote, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.StoryID, model.UserID, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original vote [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.NowPointer()
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"story_id\" = $6 and \"user_id\" = $7", s.db.Placeholder())
	data := model.ToData()
	data = append(data, model.StoryID, model.UserID)
	_, err = s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Vote) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *Vote, _ int) {
		model.Created = time.Now()
		model.Updated = util.NowPointer()
	})
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"story_id", "user_id"}, columnsQuoted, s.db.Placeholder())
	data := lo.FlatMap(models, func(model *Vote, _ int) []any {
		return model.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, storyID uuid.UUID, userID uuid.UUID, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0), s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, 1, logger, storyID, userID)
	return err
}

func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	q := database.SQLDelete(tableQuoted, wc, s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, expected, logger, values...)
	return err
}
