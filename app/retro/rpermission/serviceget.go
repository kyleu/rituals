// Package rpermission - Content managed by Project Forge, see [projectforge.md] for details.
package rpermission

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

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, retroID uuid.UUID, key string, value string, logger util.Logger) (*RetroPermission, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, retroID, key, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retroPermission by retroID [%v], key [%v], value [%v]", retroID, key, value)
	}
	return ret.ToRetroPermission(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, pks ...*PK) (RetroPermissions, error) {
	if len(pks) == 0 {
		return RetroPermissions{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(retro_id = $%d and key = $%d and value = $%d)", (idx*3)+1, (idx*3)+2, (idx*3)+3)
	})
	wc += ")"
	ret := rows{}
	params = filters(params)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.RetroID, x.Key, x.Value}
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get RetroPermissions for [%d] pks", len(pks))
	}
	return ret.ToRetroPermissions(), nil
}

func (s *Service) GetByKey(ctx context.Context, tx *sqlx.Tx, key string, params *filter.Params, logger util.Logger) (RetroPermissions, error) {
	params = filters(params)
	wc := "\"key\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, key)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by key [%v]", key)
	}
	return ret.ToRetroPermissions(), nil
}

func (s *Service) GetByKeys(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, keys ...string) (RetroPermissions, error) {
	if len(keys) == 0 {
		return RetroPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("key", len(keys), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(keys)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get RetroPermissions for [%d] keys", len(keys))
	}
	return ret.ToRetroPermissions(), nil
}

func (s *Service) GetByRetroID(ctx context.Context, tx *sqlx.Tx, retroID uuid.UUID, params *filter.Params, logger util.Logger) (RetroPermissions, error) {
	params = filters(params)
	wc := "\"retro_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, retroID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by retroID [%v]", retroID)
	}
	return ret.ToRetroPermissions(), nil
}

func (s *Service) GetByRetroIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, retroIDs ...uuid.UUID) (RetroPermissions, error) {
	if len(retroIDs) == 0 {
		return RetroPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("retro_id", len(retroIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(retroIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get RetroPermissions for [%d] retroIDs", len(retroIDs))
	}
	return ret.ToRetroPermissions(), nil
}

func (s *Service) GetByValue(ctx context.Context, tx *sqlx.Tx, value string, params *filter.Params, logger util.Logger) (RetroPermissions, error) {
	params = filters(params)
	wc := "\"value\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by value [%v]", value)
	}
	return ret.ToRetroPermissions(), nil
}

func (s *Service) GetByValues(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, values ...string) (RetroPermissions, error) {
	if len(values) == 0 {
		return RetroPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("value", len(values), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(values)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get RetroPermissions for [%d] values", len(values))
	}
	return ret.ToRetroPermissions(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*RetroPermission, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random permissions")
	}
	return ret.ToRetroPermission(), nil
}
