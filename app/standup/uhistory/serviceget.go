// Content managed by Project Forge, see [projectforge.md] for details.
package uhistory

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (StandupHistories, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get histories")
	}
	return ret.ToStandupHistories(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple(columnsString, tableQuoted, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of histories")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, slug string, logger util.Logger) (*StandupHistory, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, slug)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standupHistory by slug [%v]", slug)
	}
	return ret.ToStandupHistory(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, slugs ...string) (StandupHistories, error) {
	if len(slugs) == 0 {
		return StandupHistories{}, nil
	}
	wc := database.SQLInClause("slug", len(slugs), 0)
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(slugs))
	for _, x := range slugs {
		vals = append(vals, x)
	}
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get StandupHistories for [%d] slugs", len(slugs))
	}
	return ret.ToStandupHistories(), nil
}

func (s *Service) GetByStandupID(ctx context.Context, tx *sqlx.Tx, standupID uuid.UUID, params *filter.Params, logger util.Logger) (StandupHistories, error) {
	params = filters(params)
	wc := "\"standup_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, standupID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standup_histories by standupID [%v]", standupID)
	}
	return ret.ToStandupHistories(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger) (StandupHistories, error) {
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get histories using custom SQL")
	}
	return ret.ToStandupHistories(), nil
}