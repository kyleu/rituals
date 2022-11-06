package vote

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) GetByStoryIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, storyIDs ...string) (Votes, error) {
	params = filters(params)
	var placeholders []string
	var values []any
	for idx, sid := range storyIDs {
		placeholders = append(placeholders, fmt.Sprintf("$%d", idx+1))
		values = append(values, sid)
	}
	wc := "\"story_id\" in (%s)"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get votes by storyIDs")
	}
	return ret.ToVotes(), nil
}
