package story

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) GetByEstimateIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, estimateIDs ...uuid.UUID) (Stories, error) {
	if len(estimateIDs) == 0 {
		return Stories{}, nil
	}
	wc := database.SQLInClause("estimate_id", len(estimateIDs), 0, "")
	ret := rows{}
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	err := s.db.Select(ctx, &ret, q, tx, logger, util.InterfaceArrayFrom(estimateIDs...)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Stories for [%d] estimate IDs", len(estimateIDs))
	}
	return ret.ToStories(), nil
}
