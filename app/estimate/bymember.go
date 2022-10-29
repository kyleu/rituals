package estimate

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) GetByMember(ctx context.Context, tx *sqlx.Tx, owner uuid.UUID, params *filter.Params, logger util.Logger) (Estimates, error) {
	params = filters(params)
	wc := "\"owner\" = $1 or id in (select estimate_id from estimate_member where user_id = $1)"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, owner)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get estimates by owner [%v]", owner)
	}
	return ret.ToEstimates(), nil
}
