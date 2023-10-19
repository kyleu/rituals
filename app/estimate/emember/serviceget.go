// Package emember - Content managed by Project Forge, see [projectforge.md] for details.
package emember

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

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (EstimateMembers, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get members")
	}
	return ret.ToEstimateMembers(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Placeholder(), whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of members")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, estimateID uuid.UUID, userID uuid.UUID, logger util.Logger) (*EstimateMember, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
	err := s.db.Get(ctx, ret, q, tx, logger, estimateID, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimateMember by estimateID [%v], userID [%v]", estimateID, userID)
	}
	return ret.ToEstimateMember(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, pks ...*PK) (EstimateMembers, error) {
	if len(pks) == 0 {
		return EstimateMembers{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(estimate_id = $%d and user_id = $%d)", (idx*2)+1, (idx*2)+2)
	})
	wc += ")"
	ret := rows{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.EstimateID, x.UserID}
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get EstimateMembers for [%d] pks", len(pks))
	}
	return ret.ToEstimateMembers(), nil
}

func (s *Service) GetByEstimateID(ctx context.Context, tx *sqlx.Tx, estimateID uuid.UUID, params *filter.Params, logger util.Logger) (EstimateMembers, error) {
	params = filters(params)
	wc := "\"estimate_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, estimateID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimate_members by estimateID [%v]", estimateID)
	}
	return ret.ToEstimateMembers(), nil
}

func (s *Service) GetByUserID(ctx context.Context, tx *sqlx.Tx, userID uuid.UUID, params *filter.Params, logger util.Logger) (EstimateMembers, error) {
	params = filters(params)
	wc := "\"user_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimate_members by userID [%v]", userID)
	}
	return ret.ToEstimateMembers(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (EstimateMembers, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get members using custom SQL")
	}
	return ret.ToEstimateMembers(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*EstimateMember, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Placeholder())
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random members")
	}
	return ret.ToEstimateMember(), nil
}
