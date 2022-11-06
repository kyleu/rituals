// Content managed by Project Forge, see [projectforge.md] for details.
package team

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Teams, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get teams")
	}
	return ret.ToTeams(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple(columnsString, tableQuoted, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of teams")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Team, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get team by id [%v]", id)
	}
	return ret.ToTeam(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, ids ...uuid.UUID) (Teams, error) {
	if len(ids) == 0 {
		return Teams{}, nil
	}
	wc := database.SQLInClause("id", len(ids), 0)
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(ids))
	for _, x := range ids {
		vals = append(vals, x)
	}
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Teams for [%d] ids", len(ids))
	}
	return ret.ToTeams(), nil
}

func (s *Service) GetByOwner(ctx context.Context, tx *sqlx.Tx, owner uuid.UUID, params *filter.Params, logger util.Logger) (Teams, error) {
	params = filters(params)
	wc := "\"owner\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, owner)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get teams by owner [%v]", owner)
	}
	return ret.ToTeams(), nil
}

func (s *Service) GetBySlug(ctx context.Context, tx *sqlx.Tx, slug string, params *filter.Params, logger util.Logger) (Teams, error) {
	params = filters(params)
	wc := "\"slug\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, slug)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get teams by slug [%v]", slug)
	}
	return ret.ToTeams(), nil
}

func (s *Service) GetByStatus(ctx context.Context, tx *sqlx.Tx, status enum.SessionStatus, params *filter.Params, logger util.Logger) (Teams, error) {
	params = filters(params)
	wc := "\"status\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, status)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get teams by status [%v]", status)
	}
	return ret.ToTeams(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger) (Teams, error) {
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get teams using custom SQL")
	}
	return ret.ToTeams(), nil
}