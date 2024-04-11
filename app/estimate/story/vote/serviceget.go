// Package vote - Content managed by Project Forge, see [projectforge.md] for details.
package vote

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, storyID uuid.UUID, userID uuid.UUID, logger util.Logger) (*Vote, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, storyID, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get vote by storyID [%v], userID [%v]", storyID, userID)
	}
	return ret.ToVote(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, pks ...*PK) (Votes, error) {
	if len(pks) == 0 {
		return Votes{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(story_id = $%d and user_id = $%d)", (idx*2)+1, (idx*2)+2)
	})
	wc += ")"
	ret := rows{}
	params = filters(params)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.StoryID, x.UserID}
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Votes for [%d] pks", len(pks))
	}
	return ret.ToVotes(), nil
}

func (s *Service) GetByStoryID(ctx context.Context, tx *sqlx.Tx, storyID uuid.UUID, params *filter.Params, logger util.Logger) (Votes, error) {
	params = filters(params)
	wc := "\"story_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, storyID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Votes by storyID [%v]", storyID)
	}
	return ret.ToVotes(), nil
}

func (s *Service) GetByStoryIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, storyIDs ...uuid.UUID) (Votes, error) {
	if len(storyIDs) == 0 {
		return Votes{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("story_id", len(storyIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(storyIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Votes for [%d] storyIDs", len(storyIDs))
	}
	return ret.ToVotes(), nil
}

func (s *Service) GetByUserID(ctx context.Context, tx *sqlx.Tx, userID uuid.UUID, params *filter.Params, logger util.Logger) (Votes, error) {
	params = filters(params)
	wc := "\"user_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Votes by userID [%v]", userID)
	}
	return ret.ToVotes(), nil
}

func (s *Service) GetByUserIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, userIDs ...uuid.UUID) (Votes, error) {
	if len(userIDs) == 0 {
		return Votes{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("user_id", len(userIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(userIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Votes for [%d] userIDs", len(userIDs))
	}
	return ret.ToVotes(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Vote, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random votes")
	}
	return ret.ToVote(), nil
}
