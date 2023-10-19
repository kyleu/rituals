// Package rmember - Content managed by Project Forge, see [projectforge.md] for details.
package rmember

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (RetroMembers, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get members")
	}
	return ret.ToRetroMembers(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Placeholder(), whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of members")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, retroID uuid.UUID, userID uuid.UUID, logger util.Logger) (*RetroMember, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
	err := s.db.Get(ctx, ret, q, tx, logger, retroID, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retroMember by retroID [%v], userID [%v]", retroID, userID)
	}
	return ret.ToRetroMember(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, pks ...*PK) (RetroMembers, error) {
	if len(pks) == 0 {
		return RetroMembers{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(retro_id = $%d and user_id = $%d)", (idx*2)+1, (idx*2)+2)
	})
	wc += ")"
	ret := rows{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.RetroID, x.UserID}
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get RetroMembers for [%d] pks", len(pks))
	}
	return ret.ToRetroMembers(), nil
}

func (s *Service) GetByRetroID(ctx context.Context, tx *sqlx.Tx, retroID uuid.UUID, params *filter.Params, logger util.Logger) (RetroMembers, error) {
	params = filters(params)
	wc := "\"retro_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, retroID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retro_members by retroID [%v]", retroID)
	}
	return ret.ToRetroMembers(), nil
}

func (s *Service) GetByUserID(ctx context.Context, tx *sqlx.Tx, userID uuid.UUID, params *filter.Params, logger util.Logger) (RetroMembers, error) {
	params = filters(params)
	wc := "\"user_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retro_members by userID [%v]", userID)
	}
	return ret.ToRetroMembers(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (RetroMembers, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get members using custom SQL")
	}
	return ret.ToRetroMembers(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*RetroMember, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Placeholder())
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random members")
	}
	return ret.ToRetroMember(), nil
}
