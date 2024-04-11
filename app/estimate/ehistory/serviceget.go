// Package ehistory - Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, slug string, logger util.Logger) (*EstimateHistory, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, slug)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimateHistory by slug [%v]", slug)
	}
	return ret.ToEstimateHistory(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, slugs ...string) (EstimateHistories, error) {
	if len(slugs) == 0 {
		return EstimateHistories{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("slug", len(slugs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(slugs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get EstimateHistories for [%d] slugs", len(slugs))
	}
	return ret.ToEstimateHistories(), nil
}

//nolint:lll
func (s *Service) GetByEstimateID(ctx context.Context, tx *sqlx.Tx, estimateID uuid.UUID, params *filter.Params, logger util.Logger) (EstimateHistories, error) {
	params = filters(params)
	wc := "\"estimate_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, estimateID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Histories by estimateID [%v]", estimateID)
	}
	return ret.ToEstimateHistories(), nil
}

//nolint:lll
func (s *Service) GetByEstimateIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, estimateIDs ...uuid.UUID) (EstimateHistories, error) {
	if len(estimateIDs) == 0 {
		return EstimateHistories{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("estimate_id", len(estimateIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(estimateIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get EstimateHistories for [%d] estimateIDs", len(estimateIDs))
	}
	return ret.ToEstimateHistories(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*EstimateHistory, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random histories")
	}
	return ret.ToEstimateHistory(), nil
}
