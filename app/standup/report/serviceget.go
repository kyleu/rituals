package report

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

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Report, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get report by id [%v]", id)
	}
	return ret.ToReport(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, ids ...uuid.UUID) (Reports, error) {
	if len(ids) == 0 {
		return Reports{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("id", len(ids), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Reports for [%d] ids", len(ids))
	}
	return ret.ToReports(), nil
}

func (s *Service) GetByStandupID(ctx context.Context, tx *sqlx.Tx, standupID uuid.UUID, params *filter.Params, logger util.Logger) (Reports, error) {
	params = filters(params)
	wc := "\"standup_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, standupID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Reports by standupID [%v]", standupID)
	}
	return ret.ToReports(), nil
}

func (s *Service) GetByStandupIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, standupIDs ...uuid.UUID) (Reports, error) {
	if len(standupIDs) == 0 {
		return Reports{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("standup_id", len(standupIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(standupIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Reports for [%d] standupIDs", len(standupIDs))
	}
	return ret.ToReports(), nil
}

func (s *Service) GetByUserID(ctx context.Context, tx *sqlx.Tx, userID uuid.UUID, params *filter.Params, logger util.Logger) (Reports, error) {
	params = filters(params)
	wc := "\"user_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Reports by userID [%v]", userID)
	}
	return ret.ToReports(), nil
}

func (s *Service) GetByUserIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, userIDs ...uuid.UUID) (Reports, error) {
	if len(userIDs) == 0 {
		return Reports{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("user_id", len(userIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(userIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Reports for [%d] userIDs", len(userIDs))
	}
	return ret.ToReports(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Report, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random reports")
	}
	return ret.ToReport(), nil
}
