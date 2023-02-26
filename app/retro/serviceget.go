// Content managed by Project Forge, see [projectforge.md] for details.
package retro

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

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Retros, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get retros")
	}
	return ret.ToRetros(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of retros")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Retro, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retro by id [%v]", id)
	}
	return ret.ToRetro(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, ids ...uuid.UUID) (Retros, error) {
	if len(ids) == 0 {
		return Retros{}, nil
	}
	wc := database.SQLInClause("id", len(ids), 0)
	ret := rows{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(ids))
	for _, x := range ids {
		vals = append(vals, x)
	}
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Retros for [%d] ids", len(ids))
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetBySlug(ctx context.Context, tx *sqlx.Tx, slug string, logger util.Logger) (*Retro, error) {
	wc := "\"slug\" = $1"
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, slug)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retro by slug [%v]", slug)
	}
	return ret.ToRetro(), nil
}

func (s *Service) GetBySprintID(ctx context.Context, tx *sqlx.Tx, sprintID *uuid.UUID, params *filter.Params, logger util.Logger) (Retros, error) {
	params = filters(params)
	wc := "\"sprint_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, sprintID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retros by sprintID [%v]", sprintID)
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetByStatus(ctx context.Context, tx *sqlx.Tx, status enum.SessionStatus, params *filter.Params, logger util.Logger) (Retros, error) {
	params = filters(params)
	wc := "\"status\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, status)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retros by status [%v]", status)
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetByTeamID(ctx context.Context, tx *sqlx.Tx, teamID *uuid.UUID, params *filter.Params, logger util.Logger) (Retros, error) {
	params = filters(params)
	wc := "\"team_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, teamID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retros by teamID [%v]", teamID)
	}
	return ret.ToRetros(), nil
}

const searchClause = "(lower(id::text) like $1 or lower(slug) like $1 or lower(title) like $1 or lower(categories::text) like $1)"

func (s *Service) Search(ctx context.Context, query string, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Retros, error) {
	params = filters(params)
	wc := searchClause
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	return ret.ToRetros(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Retros, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get retros using custom SQL")
	}
	return ret.ToRetros(), nil
}
