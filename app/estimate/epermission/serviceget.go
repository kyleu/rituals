package epermission

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, estimateID uuid.UUID, key string, value string, logger util.Logger) (*EstimatePermission, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, estimateID, key, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimatePermission by estimateID [%v], key [%v], value [%v]", estimateID, key, value)
	}
	return ret.ToEstimatePermission(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, pks ...*PK) (EstimatePermissions, error) {
	if len(pks) == 0 {
		return EstimatePermissions{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(estimate_id = $%d and key = $%d and value = $%d)", (idx*3)+1, (idx*3)+2, (idx*3)+3)
	})
	wc += ")"
	ret := rows{}
	params = filters(params)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.EstimateID, x.Key, x.Value}
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get EstimatePermissions for [%d] pks", len(pks))
	}
	return ret.ToEstimatePermissions(), nil
}

//nolint:lll
func (s *Service) GetByEstimateID(ctx context.Context, tx *sqlx.Tx, estimateID uuid.UUID, params *filter.Params, logger util.Logger) (EstimatePermissions, error) {
	params = filters(params)
	wc := "\"estimate_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, estimateID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by estimateID [%v]", estimateID)
	}
	return ret.ToEstimatePermissions(), nil
}

//nolint:lll
func (s *Service) GetByEstimateIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, estimateIDs ...uuid.UUID) (EstimatePermissions, error) {
	if len(estimateIDs) == 0 {
		return EstimatePermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("estimate_id", len(estimateIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(estimateIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get EstimatePermissions for [%d] estimateIDs", len(estimateIDs))
	}
	return ret.ToEstimatePermissions(), nil
}

func (s *Service) GetByKey(ctx context.Context, tx *sqlx.Tx, key string, params *filter.Params, logger util.Logger) (EstimatePermissions, error) {
	params = filters(params)
	wc := "\"key\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, key)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by key [%v]", key)
	}
	return ret.ToEstimatePermissions(), nil
}

func (s *Service) GetByKeys(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, keys ...string) (EstimatePermissions, error) {
	if len(keys) == 0 {
		return EstimatePermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("key", len(keys), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(keys)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get EstimatePermissions for [%d] keys", len(keys))
	}
	return ret.ToEstimatePermissions(), nil
}

func (s *Service) GetByValue(ctx context.Context, tx *sqlx.Tx, value string, params *filter.Params, logger util.Logger) (EstimatePermissions, error) {
	params = filters(params)
	wc := "\"value\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by value [%v]", value)
	}
	return ret.ToEstimatePermissions(), nil
}

func (s *Service) GetByValues(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, values ...string) (EstimatePermissions, error) {
	if len(values) == 0 {
		return EstimatePermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("value", len(values), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(values)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get EstimatePermissions for [%d] values", len(values))
	}
	return ret.ToEstimatePermissions(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*EstimatePermission, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random permissions")
	}
	return ret.ToEstimatePermission(), nil
}
