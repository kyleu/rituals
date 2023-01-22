package standup

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) GetByMember(ctx context.Context, tx *sqlx.Tx, u uuid.UUID, params *filter.Params, logger util.Logger) (Standups, error) {
	params = filters(params)
	wc := "\"owner\" = $1 or id in (select standup_id from standup_member where user_id = $1)"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, u)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standups for member [%v]", u)
	}
	return ret.ToStandups(), nil
}
