// Content managed by Project Forge, see [projectforge.md] for details.
package standup

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Standups, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get standups")
	}
	return ret.ToStandups(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Placeholder(), whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of standups")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Standup, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standup by id [%v]", id)
	}
	return ret.ToStandup(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, ids ...uuid.UUID) (Standups, error) {
	if len(ids) == 0 {
		return Standups{}, nil
	}
	wc := database.SQLInClause("id", len(ids), 0, s.db.Placeholder())
	ret := rows{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Standups for [%d] ids", len(ids))
	}
	return ret.ToStandups(), nil
}

func (s *Service) GetBySlug(ctx context.Context, tx *sqlx.Tx, slug string, logger util.Logger) (*Standup, error) {
	wc := "\"slug\" = $1"
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
	err := s.db.Get(ctx, ret, q, tx, logger, slug)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standup by slug [%v]", slug)
	}
	return ret.ToStandup(), nil
}

func (s *Service) GetBySprintID(ctx context.Context, tx *sqlx.Tx, sprintID *uuid.UUID, params *filter.Params, logger util.Logger) (Standups, error) {
	params = filters(params)
	wc := "\"sprint_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, sprintID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standups by sprintID [%v]", sprintID)
	}
	return ret.ToStandups(), nil
}

func (s *Service) GetByStatus(ctx context.Context, tx *sqlx.Tx, status enum.SessionStatus, params *filter.Params, logger util.Logger) (Standups, error) {
	params = filters(params)
	wc := "\"status\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, status)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standups by status [%v]", status)
	}
	return ret.ToStandups(), nil
}

func (s *Service) GetByTeamID(ctx context.Context, tx *sqlx.Tx, teamID *uuid.UUID, params *filter.Params, logger util.Logger) (Standups, error) {
	params = filters(params)
	wc := "\"team_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, teamID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get standups by teamID [%v]", teamID)
	}
	return ret.ToStandups(), nil
}

const searchClause = "(lower(id::text) like $1 or lower(slug) like $1 or lower(title) like $1)"

func (s *Service) Search(ctx context.Context, query string, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Standups, error) {
	params = filters(params)
	wc := searchClause
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	return ret.ToStandups(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Standups, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get standups using custom SQL")
	}
	return ret.ToStandups(), nil
}
