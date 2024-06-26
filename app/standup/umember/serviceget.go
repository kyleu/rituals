package umember

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

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, standupID uuid.UUID, userID uuid.UUID, logger util.Logger) (*StandupMember, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, standupID, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standupMember by standupID [%v], userID [%v]", standupID, userID)
	}
	return ret.ToStandupMember(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, pks ...*PK) (StandupMembers, error) {
	if len(pks) == 0 {
		return StandupMembers{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(standup_id = $%d and user_id = $%d)", (idx*2)+1, (idx*2)+2)
	})
	wc += ")"
	ret := rows{}
	params = filters(params)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.StandupID, x.UserID}
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get StandupMembers for [%d] pks", len(pks))
	}
	return ret.ToStandupMembers(), nil
}

func (s *Service) GetByStandupID(ctx context.Context, tx *sqlx.Tx, standupID uuid.UUID, params *filter.Params, logger util.Logger) (StandupMembers, error) {
	params = filters(params)
	wc := "\"standup_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, standupID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Members by standupID [%v]", standupID)
	}
	return ret.ToStandupMembers(), nil
}

//nolint:lll
func (s *Service) GetByStandupIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, standupIDs ...uuid.UUID) (StandupMembers, error) {
	if len(standupIDs) == 0 {
		return StandupMembers{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("standup_id", len(standupIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(standupIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get StandupMembers for [%d] standupIDs", len(standupIDs))
	}
	return ret.ToStandupMembers(), nil
}

func (s *Service) GetByUserID(ctx context.Context, tx *sqlx.Tx, userID uuid.UUID, params *filter.Params, logger util.Logger) (StandupMembers, error) {
	params = filters(params)
	wc := "\"user_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Members by userID [%v]", userID)
	}
	return ret.ToStandupMembers(), nil
}

func (s *Service) GetByUserIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, userIDs ...uuid.UUID) (StandupMembers, error) {
	if len(userIDs) == 0 {
		return StandupMembers{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("user_id", len(userIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(userIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get StandupMembers for [%d] userIDs", len(userIDs))
	}
	return ret.ToStandupMembers(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*StandupMember, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random members")
	}
	return ret.ToStandupMember(), nil
}
