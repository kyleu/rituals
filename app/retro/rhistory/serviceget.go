// Package rhistory - Content managed by Project Forge, see [projectforge.md] for details.
package rhistory

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (RetroHistories, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get histories")
	}
	return ret.ToRetroHistories(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Type, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of histories")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, slug string, logger util.Logger) (*RetroHistory, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, slug)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retroHistory by slug [%v]", slug)
	}
	return ret.ToRetroHistory(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, slugs ...string) (RetroHistories, error) {
	if len(slugs) == 0 {
		return RetroHistories{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("slug", len(slugs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(slugs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get RetroHistories for [%d] slugs", len(slugs))
	}
	return ret.ToRetroHistories(), nil
}

func (s *Service) GetByRetroID(ctx context.Context, tx *sqlx.Tx, retroID uuid.UUID, params *filter.Params, logger util.Logger) (RetroHistories, error) {
	params = filters(params)
	wc := "\"retro_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, retroID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Histories by retroID [%v]", retroID)
	}
	return ret.ToRetroHistories(), nil
}

func (s *Service) GetByRetroIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, retroIDs ...uuid.UUID) (RetroHistories, error) {
	if len(retroIDs) == 0 {
		return RetroHistories{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("retro_id", len(retroIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(retroIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get RetroHistories for [%d] retroIDs", len(retroIDs))
	}
	return ret.ToRetroHistories(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (RetroHistories, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get histories using custom SQL")
	}
	return ret.ToRetroHistories(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*RetroHistory, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random histories")
	}
	return ret.ToRetroHistory(), nil
}
