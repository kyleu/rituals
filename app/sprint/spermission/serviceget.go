// Package spermission - Content managed by Project Forge, see [projectforge.md] for details.
package spermission

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

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, sprintID uuid.UUID, key string, value string, logger util.Logger) (*SprintPermission, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, sprintID, key, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get sprintPermission by sprintID [%v], key [%v], value [%v]", sprintID, key, value)
	}
	return ret.ToSprintPermission(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, pks ...*PK) (SprintPermissions, error) {
	if len(pks) == 0 {
		return SprintPermissions{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(sprint_id = $%d and key = $%d and value = $%d)", (idx*3)+1, (idx*3)+2, (idx*3)+3)
	})
	wc += ")"
	ret := rows{}
	params = filters(params)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.SprintID, x.Key, x.Value}
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get SprintPermissions for [%d] pks", len(pks))
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) GetByKey(ctx context.Context, tx *sqlx.Tx, key string, params *filter.Params, logger util.Logger) (SprintPermissions, error) {
	params = filters(params)
	wc := "\"key\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, key)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by key [%v]", key)
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) GetByKeys(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, keys ...string) (SprintPermissions, error) {
	if len(keys) == 0 {
		return SprintPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("key", len(keys), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(keys)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get SprintPermissions for [%d] keys", len(keys))
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) GetBySprintID(ctx context.Context, tx *sqlx.Tx, sprintID uuid.UUID, params *filter.Params, logger util.Logger) (SprintPermissions, error) {
	params = filters(params)
	wc := "\"sprint_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, sprintID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by sprintID [%v]", sprintID)
	}
	return ret.ToSprintPermissions(), nil
}

//nolint:lll
func (s *Service) GetBySprintIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, sprintIDs ...uuid.UUID) (SprintPermissions, error) {
	if len(sprintIDs) == 0 {
		return SprintPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("sprint_id", len(sprintIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(sprintIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get SprintPermissions for [%d] sprintIDs", len(sprintIDs))
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) GetByValue(ctx context.Context, tx *sqlx.Tx, value string, params *filter.Params, logger util.Logger) (SprintPermissions, error) {
	params = filters(params)
	wc := "\"value\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by value [%v]", value)
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) GetByValues(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, values ...string) (SprintPermissions, error) {
	if len(values) == 0 {
		return SprintPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("value", len(values), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(values)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get SprintPermissions for [%d] values", len(values))
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*SprintPermission, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random permissions")
	}
	return ret.ToSprintPermission(), nil
}
