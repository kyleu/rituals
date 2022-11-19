package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
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
	Histories   thistory.TeamHistories      `json:"histories,omitempty"`
	Members     tmember.TeamMembers         `json:"members,omitempty"`
	Self        *tmember.TeamMember         `json:"self,omitempty"`
	Permissions tpermission.TeamPermissions `json:"permissions,omitempty"`
	Sprints     sprint.Sprints              `json:"sprints,omitempty"`
	Estimates   estimate.Estimates          `json:"estimates,omitempty"`
	Standups    standup.Standups            `json:"standups,omitempty"`
	Retros      retro.Retros                `json:"retros,omitempty"`
	Comments    comment.Comments            `json:"comments,omitempty"`
	Actions     action.Actions              `json:"actions,omitempty"`
}

func (s *Service) LoadTeam(ctx context.Context, slug string, user uuid.UUID, tx *sqlx.Tx, params filter.ParamSet, logger util.Logger) (*FullTeam, error) {
	t, err := s.t.GetBySlug(ctx, tx, slug, logger)
	if err != nil {
		if hist, _ := s.th.Get(ctx, tx, slug, logger); hist != nil {
			t, err = s.t.Get(ctx, tx, hist.TeamID, logger)
			if err != nil {
				return nil, errors.Errorf("no team found with slug [%s]", slug)
			}
		}
	}
	if t == nil {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no team found with slug [%s]", slug)
		}
		t, err = s.t.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no team found with id [%s]", slug)
		}
	}
	ret := &FullTeam{Team: t}

	ret.Histories, err = s.th.GetByTeamID(ctx, tx, t.ID, params.Get("thistory", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Members, err = s.tm.GetByTeamID(ctx, tx, t.ID, params.Get("tmember", nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Self = ret.Members.Get(t.ID, user)
	ret.Permissions, err = s.tp.GetByTeamID(ctx, tx, t.ID, params.Get("tpermission", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	ret.Sprints, err = s.s.GetByTeamID(ctx, tx, &t.ID, params.Get(util.KeySprint, nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Estimates, err = s.e.GetByTeamID(ctx, tx, &t.ID, params.Get(util.KeyEstimate, nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Standups, err = s.u.GetByTeamID(ctx, tx, &t.ID, params.Get(util.KeyStandup, nil, logger), logger)
	if err != nil {
		return nil, err
	}
	ret.Retros, err = s.r.GetByTeamID(ctx, tx, &t.ID, params.Get(util.KeyRetro, nil, logger), logger)
	if err != nil {
		return nil, err
	}

	args := make([]any, 0, (len(ret.Sprints)*2)+(len(ret.Estimates)*2)+(len(ret.Standups)*2)+(len(ret.Retros)*2)+2)
	args = append(args, util.KeyTeam, t.ID)
	for _, x := range ret.Sprints {
		args = append(args, util.KeySprint, x.ID)
	}
	for _, x := range ret.Estimates {
		args = append(args, util.KeyEstimate, x.ID)
	}
	for _, x := range ret.Standups {
		args = append(args, util.KeyStandup, x.ID)
	}
	for _, x := range ret.Retros {
		args = append(args, util.KeyRetro, x.ID)
	}

	ret.Comments, err = s.c.GetByModels(ctx, tx, logger, args...)
	if err != nil {
		return nil, err
	}
	ret.Actions, err = s.a.GetByModels(ctx, tx, logger, enum.ModelServiceTeam, ret.Team.ID)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
