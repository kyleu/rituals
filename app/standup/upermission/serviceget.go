package upermission

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

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, standupID uuid.UUID, key string, value string, logger util.Logger) (*StandupPermission, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, standupID, key, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standupPermission by standupID [%v], key [%v], value [%v]", standupID, key, value)
	}
	return ret.ToStandupPermission(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, pks ...*PK) (StandupPermissions, error) {
	if len(pks) == 0 {
		return StandupPermissions{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(standup_id = $%d and key = $%d and value = $%d)", (idx*3)+1, (idx*3)+2, (idx*3)+3)
	})
	wc += ")"
	ret := rows{}
	params = filters(params)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.StandupID, x.Key, x.Value}
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get StandupPermissions for [%d] pks", len(pks))
	}
	return ret.ToStandupPermissions(), nil
}

func (s *Service) GetByKey(ctx context.Context, tx *sqlx.Tx, key string, params *filter.Params, logger util.Logger) (StandupPermissions, error) {
	params = filters(params)
	wc := "\"key\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, key)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by key [%v]", key)
	}
	return ret.ToStandupPermissions(), nil
}

func (s *Service) GetByKeys(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, keys ...string) (StandupPermissions, error) {
	if len(keys) == 0 {
		return StandupPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("key", len(keys), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(keys)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get StandupPermissions for [%d] keys", len(keys))
	}
	return ret.ToStandupPermissions(), nil
}

func (s *Service) GetByStandupID(ctx context.Context, tx *sqlx.Tx, standupID uuid.UUID, params *filter.Params, logger util.Logger) (StandupPermissions, error) {
	params = filters(params)
	wc := "\"standup_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, standupID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by standupID [%v]", standupID)
	}
	return ret.ToStandupPermissions(), nil
}

//nolint:lll
func (s *Service) GetByStandupIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, standupIDs ...uuid.UUID) (StandupPermissions, error) {
	if len(standupIDs) == 0 {
		return StandupPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("standup_id", len(standupIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(standupIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get StandupPermissions for [%d] standupIDs", len(standupIDs))
	}
	return ret.ToStandupPermissions(), nil
}

func (s *Service) GetByValue(ctx context.Context, tx *sqlx.Tx, value string, params *filter.Params, logger util.Logger) (StandupPermissions, error) {
	params = filters(params)
	wc := "\"value\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by value [%v]", value)
	}
	return ret.ToStandupPermissions(), nil
}

func (s *Service) GetByValues(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, values ...string) (StandupPermissions, error) {
	if len(values) == 0 {
		return StandupPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("value", len(values), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(values)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get StandupPermissions for [%d] values", len(values))
	}
	return ret.ToStandupPermissions(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*StandupPermission, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random permissions")
	}
	return ret.ToStandupPermission(), nil
}
