// Content managed by Project Forge, see [projectforge.md] for details.
package tpermission

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

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (TeamPermissions, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get permissions")
	}
	return ret.ToTeamPermissions(), nil
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

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, teamID uuid.UUID, k string, v string, logger util.Logger) (*TeamPermission, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, teamID, k, v)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get teamPermission by teamID [%v], k [%v], v [%v]", teamID, k, v)
	}
	return ret.ToTeamPermission(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, pks ...*PK) (TeamPermissions, error) {
	if len(pks) == 0 {
		return TeamPermissions{}, nil
	}
	wc := "("
	for idx := range pks {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(team_id = $%d and k = $%d and v = $%d)", (idx*3)+1, (idx*3)+2, (idx*3)+3)
	}
	wc += ")"
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(pks)*3)
	for _, x := range pks {
		vals = append(vals, x.TeamID, x.K, x.V)
	}
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get TeamPermissions for [%d] pks", len(pks))
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) GetByK(ctx context.Context, tx *sqlx.Tx, k string, params *filter.Params, logger util.Logger) (TeamPermissions, error) {
	params = filters(params)
	wc := "\"k\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, k)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get team_permissions by k [%v]", k)
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) GetByTeamID(ctx context.Context, tx *sqlx.Tx, teamID uuid.UUID, params *filter.Params, logger util.Logger) (TeamPermissions, error) {
	params = filters(params)
	wc := "\"team_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, teamID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get team_permissions by teamID [%v]", teamID)
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) GetByV(ctx context.Context, tx *sqlx.Tx, v string, params *filter.Params, logger util.Logger) (TeamPermissions, error) {
	params = filters(params)
	wc := "\"v\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, v)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get team_permissions by v [%v]", v)
	}
	return ret.ToTeamPermissions(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger) (TeamPermissions, error) {
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get permissions using custom SQL")
	}
	return ret.ToTeamPermissions(), nil
}
