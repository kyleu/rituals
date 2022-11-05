package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type FullStandup struct {
	Standup     *standup.Standup               `json:"standup"`
	Histories   uhistory.StandupHistories      `json:"histories"`
	Members     umember.StandupMembers         `json:"members"`
	Permissions upermission.StandupPermissions `json:"permissions"`
	Team        *team.Team                     `json:"team"`
	Sprint      *sprint.Sprint                 `json:"sprint"`
	Reports     report.Reports                 `json:"reports"`
}

func (s *Service) CreateStandup(
	ctx context.Context, id uuid.UUID, slug string, title string, user uuid.UUID, name string, teamID *uuid.UUID, sprintID *uuid.UUID, logger util.Logger,
) (*standup.Standup, error) {
	model := &standup.Standup{
		ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, Owner: user, TeamID: teamID, SprintID: sprintID, Created: time.Now(),
	}
	err := s.u.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save standup")
	}
	member := &umember.StandupMember{StandupID: model.ID, UserID: user, Name: name, Role: enum.MemberStatusOwner, Created: time.Now()}
	err = s.um.Create(ctx, nil, logger, member)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save standup member")
	}
	return model, nil
}

func (s *Service) LoadStandup(ctx context.Context, slug string, user uuid.UUID, tx *sqlx.Tx, logger util.Logger) (*FullStandup, error) {
	bySlug, err := s.u.GetBySlug(ctx, tx, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no standup found with slug [%s]", slug)
		}
		u, err := s.u.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no standup found with slug [%s]", slug)
		}
		bySlug = standup.Standups{u}
	}
	u := bySlug[0]
	ret := &FullStandup{Standup: u}

	ret.Histories, err = s.uh.GetByStandupID(ctx, tx, u.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	ret.Members, err = s.um.GetByStandupID(ctx, tx, u.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	ret.Permissions, err = s.up.GetByStandupID(ctx, tx, u.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	if u.TeamID != nil {
		ret.Team, err = s.t.Get(ctx, tx, *u.TeamID, logger)
		if err != nil {
			return nil, err
		}
	}
	if u.SprintID != nil {
		ret.Sprint, err = s.s.Get(ctx, tx, *u.SprintID, logger)
		if err != nil {
			return nil, err
		}
	}

	ret.Reports, err = s.rt.GetByStandupID(ctx, tx, u.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
