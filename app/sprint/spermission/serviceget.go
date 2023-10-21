// Package spermission - Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (SprintPermissions, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get permissions")
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Placeholder(), whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of permissions")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, sprintID uuid.UUID, key string, value string, logger util.Logger) (*SprintPermission, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
	err := s.db.Get(ctx, ret, q, tx, logger, sprintID, key, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get sprintPermission by sprintID [%v], key [%v], value [%v]", sprintID, key, value)
	}
	return ret.ToSprintPermission(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, pks ...*PK) (SprintPermissions, error) {
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
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
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
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, key)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get sprint_permissions by key [%v]", key)
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) GetBySprintID(ctx context.Context, tx *sqlx.Tx, sprintID uuid.UUID, params *filter.Params, logger util.Logger) (SprintPermissions, error) {
	params = filters(params)
	wc := "\"sprint_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, sprintID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get sprint_permissions by sprintID [%v]", sprintID)
	}
	return ret.ToSprintPermissions(), nil
}

//nolint:lll
func (s *Service) GetBySprintIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, sprintIDs ...uuid.UUID) (SprintPermissions, error) {
	if len(sprintIDs) == 0 {
		return SprintPermissions{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("sprint_id", len(sprintIDs), 0, s.db.Placeholder())
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
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
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get sprint_permissions by value [%v]", value)
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (SprintPermissions, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get permissions using custom SQL")
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*SprintPermission, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Placeholder())
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random permissions")
	}
	return ret.ToSprintPermission(), nil
}
