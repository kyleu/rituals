// Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (SprintPermissions, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
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
	q := database.SQLSelectSimple(columnsString, tableQuoted, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of permissions")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, sprintID uuid.UUID, k string, v string, logger util.Logger) (*SprintPermission, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, sprintID, k, v)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get sprintPermission by sprintID [%v], k [%v], v [%v]", sprintID, k, v)
	}
	return ret.ToSprintPermission(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, pks ...*PK) (SprintPermissions, error) {
	if len(pks) == 0 {
		return SprintPermissions{}, nil
	}
	wc := "("
	for idx := range pks {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(sprint_id = $%d and k = $%d and v = $%d)", (idx*3)+1, (idx*3)+2, (idx*3)+3)
	}
	wc += ")"
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(pks)*3)
	for _, x := range pks {
		vals = append(vals, x.SprintID, x.K, x.V)
	}
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get SprintPermissions for [%d] pks", len(pks))
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) GetByK(ctx context.Context, tx *sqlx.Tx, k string, params *filter.Params, logger util.Logger) (SprintPermissions, error) {
	params = filters(params)
	wc := "\"k\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, k)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get sprint_permissions by k [%v]", k)
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) GetBySprintID(ctx context.Context, tx *sqlx.Tx, sprintID uuid.UUID, params *filter.Params, logger util.Logger) (SprintPermissions, error) {
	params = filters(params)
	wc := "\"sprint_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, sprintID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get sprint_permissions by sprintID [%v]", sprintID)
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) GetByV(ctx context.Context, tx *sqlx.Tx, v string, params *filter.Params, logger util.Logger) (SprintPermissions, error) {
	params = filters(params)
	wc := "\"v\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, v)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get sprint_permissions by v [%v]", v)
	}
	return ret.ToSprintPermissions(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger) (SprintPermissions, error) {
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get permissions using custom SQL")
	}
	return ret.ToSprintPermissions(), nil
}
