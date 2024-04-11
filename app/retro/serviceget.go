// Package retro - Content managed by Project Forge, see [projectforge.md] for details.
package retro

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Retro, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retro by id [%v]", id)
	}
	return ret.ToRetro(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, ids ...uuid.UUID) (Retros, error) {
	if len(ids) == 0 {
		return Retros{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("id", len(ids), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Retros for [%d] ids", len(ids))
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetBySlug(ctx context.Context, tx *sqlx.Tx, slug string, logger util.Logger) (*Retro, error) {
	wc := "\"slug\" = $1"
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, slug)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get retro by slug [%v]", slug)
	}
	return ret.ToRetro(), nil
}

func (s *Service) GetBySlugs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, slugs ...string) (Retros, error) {
	if len(slugs) == 0 {
		return Retros{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("slug", len(slugs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(slugs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Retros for [%d] slugs", len(slugs))
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetBySprintID(ctx context.Context, tx *sqlx.Tx, sprintID *uuid.UUID, params *filter.Params, logger util.Logger) (Retros, error) {
	params = filters(params)
	wc := "\"sprint_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, sprintID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Retros by sprintID [%v]", sprintID)
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetBySprintIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, sprintIDs ...*uuid.UUID) (Retros, error) {
	if len(sprintIDs) == 0 {
		return Retros{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("sprint_id", len(sprintIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(sprintIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Retros for [%d] sprintIDs", len(sprintIDs))
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetByStatus(ctx context.Context, tx *sqlx.Tx, status enum.SessionStatus, params *filter.Params, logger util.Logger) (Retros, error) {
	params = filters(params)
	wc := "\"status\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, status)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Retros by status [%v]", status)
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetByStatuses(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, statuses ...enum.SessionStatus) (Retros, error) {
	if len(statuses) == 0 {
		return Retros{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("status", len(statuses), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(statuses)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Retros for [%d] statuses", len(statuses))
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetByTeamID(ctx context.Context, tx *sqlx.Tx, teamID *uuid.UUID, params *filter.Params, logger util.Logger) (Retros, error) {
	params = filters(params)
	wc := "\"team_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, teamID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Retros by teamID [%v]", teamID)
	}
	return ret.ToRetros(), nil
}

func (s *Service) GetByTeamIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, teamIDs ...*uuid.UUID) (Retros, error) {
	if len(teamIDs) == 0 {
		return Retros{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("team_id", len(teamIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(teamIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Retros for [%d] teamIDs", len(teamIDs))
	}
	return ret.ToRetros(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Retro, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random retros")
	}
	return ret.ToRetro(), nil
}
