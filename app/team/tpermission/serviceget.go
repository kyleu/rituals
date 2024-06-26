package tpermission

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

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, teamID uuid.UUID, key string, value string, logger util.Logger) (*TeamPermission, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, teamID, key, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get teamPermission by teamID [%v], key [%v], value [%v]", teamID, key, value)
	}
	return ret.ToTeamPermission(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, pks ...*PK) (TeamPermissions, error) {
	if len(pks) == 0 {
		return TeamPermissions{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(team_id = $%d and key = $%d and value = $%d)", (idx*3)+1, (idx*3)+2, (idx*3)+3)
	})
	wc += ")"
	ret := rows{}
	params = filters(params)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.TeamID, x.Key, x.Value}
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get TeamPermissions for [%d] pks", len(pks))
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) GetByKey(ctx context.Context, tx *sqlx.Tx, key string, params *filter.Params, logger util.Logger) (TeamPermissions, error) {
	params = filters(params)
	wc := "\"key\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, key)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by key [%v]", key)
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) GetByKeys(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, keys ...string) (TeamPermissions, error) {
	if len(keys) == 0 {
		return TeamPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("key", len(keys), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(keys)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get TeamPermissions for [%d] keys", len(keys))
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) GetByTeamID(ctx context.Context, tx *sqlx.Tx, teamID uuid.UUID, params *filter.Params, logger util.Logger) (TeamPermissions, error) {
	params = filters(params)
	wc := "\"team_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, teamID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by teamID [%v]", teamID)
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) GetByTeamIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, teamIDs ...uuid.UUID) (TeamPermissions, error) {
	if len(teamIDs) == 0 {
		return TeamPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("team_id", len(teamIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(teamIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get TeamPermissions for [%d] teamIDs", len(teamIDs))
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) GetByValue(ctx context.Context, tx *sqlx.Tx, value string, params *filter.Params, logger util.Logger) (TeamPermissions, error) {
	params = filters(params)
	wc := "\"value\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Permissions by value [%v]", value)
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) GetByValues(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, values ...string) (TeamPermissions, error) {
	if len(values) == 0 {
		return TeamPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("value", len(values), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(values)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get TeamPermissions for [%d] values", len(values))
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*TeamPermission, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random permissions")
	}
	return ret.ToTeamPermission(), nil
}
