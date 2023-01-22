// Content managed by Project Forge, see [projectforge.md] for details.
package epermission

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

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (EstimatePermissions, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get permissions")
	}
	return ret.ToEstimatePermissions(), nil
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

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, estimateID uuid.UUID, key string, value string, logger util.Logger) (*EstimatePermission, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, estimateID, key, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimatePermission by estimateID [%v], key [%v], value [%v]", estimateID, key, value)
	}
	return ret.ToEstimatePermission(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, pks ...*PK) (EstimatePermissions, error) {
	if len(pks) == 0 {
		return EstimatePermissions{}, nil
	}
	wc := "("
	for idx := range pks {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(estimate_id = $%d and key = $%d and value = $%d)", (idx*3)+1, (idx*3)+2, (idx*3)+3)
	}
	wc += ")"
	ret := rows{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(pks)*3)
	for _, x := range pks {
		vals = append(vals, x.EstimateID, x.Key, x.Value)
	}
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get EstimatePermissions for [%d] pks", len(pks))
	}
	return ret.ToEstimatePermissions(), nil
}

func (s *Service) GetByEstimateID(ctx context.Context, tx *sqlx.Tx, estimateID uuid.UUID, params *filter.Params, logger util.Logger) (EstimatePermissions, error) { //nolint:lll
	params = filters(params)
	wc := "\"estimate_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, estimateID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimate_permissions by estimateID [%v]", estimateID)
	}
	return ret.ToEstimatePermissions(), nil
}

func (s *Service) GetByKey(ctx context.Context, tx *sqlx.Tx, key string, params *filter.Params, logger util.Logger) (EstimatePermissions, error) {
	params = filters(params)
	wc := "\"key\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, key)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimate_permissions by key [%v]", key)
	}
	return ret.ToEstimatePermissions(), nil
}

func (s *Service) GetByValue(ctx context.Context, tx *sqlx.Tx, value string, params *filter.Params, logger util.Logger) (EstimatePermissions, error) {
	params = filters(params)
	wc := "\"value\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, value)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimate_permissions by value [%v]", value)
	}
	return ret.ToEstimatePermissions(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (EstimatePermissions, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get permissions using custom SQL")
	}
	return ret.ToEstimatePermissions(), nil
}
