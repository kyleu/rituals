package team

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) GetByMember(ctx context.Context, tx *sqlx.Tx, u uuid.UUID, params *filter.Params, logger util.Logger) (Teams, error) {
	params = filters(params)
	wc := "id in (select team_id from team_member where user_id = $1)"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, u)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get teams for member [%v]", u)
	}
	return ret.ToTeams(), nil
}
