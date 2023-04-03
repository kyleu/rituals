package retro

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) GetByMember(ctx context.Context, tx *sqlx.Tx, u uuid.UUID, params *filter.Params, logger util.Logger) (Retros, error) {
	params = filters(params)
	wc := "id in (select retro_id from retro_member where user_id = $1)"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, u)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retros for member [%v]", u)
	}
	return ret.ToRetros(), nil
}
