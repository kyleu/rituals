package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type FullSprint struct {
	Sprint      *sprint.Sprint                `json:"sprint"`
	Histories   shistory.SprintHistories      `json:"histories"`
	Members     smember.SprintMembers         `json:"members"`
	Permissions spermission.SprintPermissions `json:"permissions"`
	Team        *team.Team                    `json:"team"`
	Estimates   estimate.Estimates            `json:"estimates"`
	Standups    standup.Standups              `json:"standups"`
	Retros      retro.Retros                  `json:"retro"`
}

func (s *Service) CreateSprint(
	ctx context.Context, id uuid.UUID, slug string, title string, user uuid.UUID, name string, teamID *uuid.UUID, logger util.Logger,
) (*sprint.Sprint, error) {
	model := &sprint.Sprint{ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, Owner: user, TeamID: teamID, Created: time.Now()}
	err := s.s.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save sprint")
	}
	member := &smember.SprintMember{SprintID: model.ID, UserID: user, Name: name, Role: enum.MemberStatusOwner, Created: time.Now()}
	err = s.sm.Create(ctx, nil, logger, member)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save sprint member")
	}
	return model, nil
}

func (s *Service) LoadSprint(ctx context.Context, slug string, user uuid.UUID, tx *sqlx.Tx, logger util.Logger) (*FullSprint, error) {
	bySlug, err := s.s.GetBySlug(ctx, tx, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no sprint found with slug [%s]", slug)
		}
		s, err := s.s.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no sprint found with slug [%s]", slug)
		}
		bySlug = sprint.Sprints{s}
	}
	spr := bySlug[0]
	ret := &FullSprint{Sprint: spr}

	ret.Histories, err = s.sh.GetBySprintID(ctx, tx, spr.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	ret.Members, err = s.sm.GetBySprintID(ctx, tx, spr.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	ret.Permissions, err = s.sp.GetBySprintID(ctx, tx, spr.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	if spr.TeamID != nil {
		ret.Team, err = s.t.Get(ctx, tx, *spr.TeamID, logger)
		if err != nil {
			return nil, err
		}
	}

	ret.Estimates, err = s.e.GetBySprintID(ctx, tx, &spr.ID, nil, logger)
	if err != nil {
		return nil, err
	}
	ret.Standups, err = s.u.GetBySprintID(ctx, tx, &spr.ID, nil, logger)
	if err != nil {
		return nil, err
	}
	ret.Retros, err = s.r.GetBySprintID(ctx, tx, &spr.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
