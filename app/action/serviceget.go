// Package action - Content managed by Project Forge, see [projectforge.md] for details.
package action

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Actions, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get actions")
	}
	return ret.ToActions(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Type, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of actions")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Action, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get action by id [%v]", id)
	}
	return ret.ToAction(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, ids ...uuid.UUID) (Actions, error) {
	if len(ids) == 0 {
		return Actions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("id", len(ids), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Actions for [%d] ids", len(ids))
	}
	return ret.ToActions(), nil
}

func (s *Service) GetByUserID(ctx context.Context, tx *sqlx.Tx, userID uuid.UUID, params *filter.Params, logger util.Logger) (Actions, error) {
	params = filters(params)
	wc := "\"user_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Actions by userID [%v]", userID)
	}
	return ret.ToActions(), nil
}

func (s *Service) GetByUserIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, userIDs ...uuid.UUID) (Actions, error) {
	if len(userIDs) == 0 {
		return Actions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("user_id", len(userIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(userIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Actions for [%d] userIDs", len(userIDs))
	}
	return ret.ToActions(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Actions, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get actions using custom SQL")
	}
	return ret.ToActions(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Action, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random actions")
	}
	return ret.ToAction(), nil
}
