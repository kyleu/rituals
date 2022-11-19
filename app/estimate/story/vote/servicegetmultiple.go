package vote

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) GetByStoryIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, storyIDs ...string) (Votes, error) {
	if len(storyIDs) == 0 {
		return Votes{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("\"story_id\"", len(storyIDs), 0)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, util.InterfaceArrayFrom(storyIDs...)...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get votes by storyIDs")
	}
	return ret.ToVotes(), nil
}
