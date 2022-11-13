package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/util"
)

type FullTeam struct {
	Team        *team.Team                  `json:"team"`
	Histories   thistory.TeamHistories      `json:"histories"`
	Members     tmember.TeamMembers         `json:"members"`
	Permissions tpermission.TeamPermissions `json:"permissions"`
	Sprints     sprint.Sprints              `json:"sprints"`
	Estimates   estimate.Estimates          `json:"estimates"`
	Standups    standup.Standups            `json:"standups"`
	Retros      retro.Retros                `json:"retro"`
}

func (s *Service) CreateTeam(
	ctx context.Context, id uuid.UUID, slug string, title string, user uuid.UUID, name string, logger util.Logger,
) (*team.Team, error) {
	model := &team.Team{ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, Owner: user, Created: time.Now()}
	err := s.t.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save team")
	}
	member := &tmember.TeamMember{TeamID: model.ID, UserID: user, Name: name, Role: enum.MemberStatusOwner, Created: time.Now()}
	err = s.tm.Create(ctx, nil, logger, member)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save team member")
	}
	return model, nil
}

func (s *Service) LoadTeam(ctx context.Context, slug string, user uuid.UUID, tx *sqlx.Tx, params filter.ParamSet, logger util.Logger) (*FullTeam, error) {
	bySlug, err := s.t.GetBySlug(ctx, tx, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no team Teams with slug [%s]", slug)
		}
		t, err := s.t.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no team Teams with slug [%s]", slug)
		}
		bySlug = team.Teams{t}
	}
	t := bySlug[0]
	ret := &FullTeam{Team: t}

	ret.Histories, err = s.th.GetByTeamID(ctx, tx, t.ID, params.Get("thistory", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	ret.Members, err = s.tm.GetByTeamID(ctx, tx, t.ID, params.Get("tmember", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	ret.Permissions, err = s.tp.GetByTeamID(ctx, tx, t.ID, params.Get("tpermission", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	ret.Sprints, err = s.s.GetByTeamID(ctx, tx, &t.ID, params.Get("sprint", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Estimates, err = s.e.GetByTeamID(ctx, tx, &t.ID, params.Get("estimate", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Standups, err = s.u.GetByTeamID(ctx, tx, &t.ID, params.Get("standup", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Retros, err = s.r.GetByTeamID(ctx, tx, &t.ID, params.Get("retro", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
