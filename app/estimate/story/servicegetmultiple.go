package story

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) GetByEstimateIDs(ctx context.Context, tx *sqlx.Tx, logger util.Logger, estimateIDs ...uuid.UUID) (Stories, error) {
	if len(estimateIDs) == 0 {
		return Stories{}, nil
	}
	wc := database.SQLInClause("estimate_id", len(estimateIDs), 0)
	ret := rows{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(estimateIDs))
	for _, x := range estimateIDs {
		vals = append(vals, x)
	}
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Stories for [%d] estimate IDs", len(estimateIDs))
	}
	return ret.ToStories(), nil
}
