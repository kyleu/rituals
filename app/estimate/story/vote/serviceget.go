// Content managed by Project Forge, see [projectforge.md] for details.
package vote

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Votes, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get votes")
	}
	return ret.ToVotes(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple(columnsString, tableQuoted, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of votes")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, storyID uuid.UUID, userID uuid.UUID, logger util.Logger) (*Vote, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, storyID, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get vote by storyID [%v], userID [%v]", storyID, userID)
	}
	return ret.ToVote(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, pks ...*PK) (Votes, error) {
	if len(pks) == 0 {
		return Votes{}, nil
	}
	wc := "("
	for idx := range pks {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(story_id = $%d and user_id = $%d)", (idx*2)+1, (idx*2)+2)
	}
	wc += ")"
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(pks)*2)
	for _, x := range pks {
		vals = append(vals, x.StoryID, x.UserID)
	}
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Votes for [%d] pks", len(pks))
	}
	return ret.ToVotes(), nil
}

func (s *Service) GetByStoryID(ctx context.Context, tx *sqlx.Tx, storyID uuid.UUID, params *filter.Params, logger util.Logger) (Votes, error) {
	params = filters(params)
	wc := "\"story_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, storyID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get votes by storyID [%v]", storyID)
	}
	return ret.ToVotes(), nil
}

func (s *Service) GetByUserID(ctx context.Context, tx *sqlx.Tx, userID uuid.UUID, params *filter.Params, logger util.Logger) (Votes, error) {
	params = filters(params)
	wc := "\"user_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get votes by userID [%v]", userID)
	}
	return ret.ToVotes(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger) (Votes, error) {
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get votes using custom SQL")
	}
	return ret.ToVotes(), nil
}