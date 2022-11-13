package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type FullRetro struct {
	Retro       *retro.Retro                 `json:"retro"`
	Histories   rhistory.RetroHistories      `json:"histories"`
	Members     rmember.RetroMembers         `json:"members"`
	Permissions rpermission.RetroPermissions `json:"permissions"`
	Team        *team.Team                   `json:"team"`
	Sprint      *sprint.Sprint               `json:"sprint"`
	Feedbacks   feedback.Feedbacks           `json:"feedbacks"`
}

func (s *Service) CreateRetro(
	ctx context.Context, id uuid.UUID, slug string, title string, user uuid.UUID, name string, teamID *uuid.UUID, sprintID *uuid.UUID, logger util.Logger,
) (*retro.Retro, error) {
	model := &retro.Retro{
		ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, Owner: user, TeamID: teamID, SprintID: sprintID, Created: time.Now(),
	}
	err := s.r.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save retro")
	}
	member := &rmember.RetroMember{RetroID: model.ID, UserID: user, Name: name, Role: enum.MemberStatusOwner, Created: time.Now()}
	err = s.rm.Create(ctx, nil, logger, member)
	if err != nil {
		return nil, errors.Wrap(err, "unable to save retro member")
	}
	return model, nil
}

func (s *Service) LoadRetro(ctx context.Context, slug string, user uuid.UUID, tx *sqlx.Tx, params filter.ParamSet, logger util.Logger) (*FullRetro, error) {
	bySlug, err := s.r.GetBySlug(ctx, tx, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no retros found with slug [%s]", slug)
		}
		r, err := s.r.Get(ctx, tx, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no retros found with slug [%s]", slug)
		}
		bySlug = retro.Retros{r}
	}
	r := bySlug[0]
	ret := &FullRetro{Retro: r}

	ret.Histories, err = s.rh.GetByRetroID(ctx, tx, r.ID, params.Get("rhistory", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	ret.Members, err = s.rm.GetByRetroID(ctx, tx, r.ID, params.Get("rmember", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	ret.Permissions, err = s.rp.GetByRetroID(ctx, tx, r.ID, params.Get("rpermission", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	if r.TeamID != nil {
		ret.Team, err = s.t.Get(ctx, tx, *r.TeamID, logger)
		if err != nil {
			return nil, err
		}
	}
	if r.SprintID != nil {
		ret.Sprint, err = s.s.Get(ctx, tx, *r.SprintID, logger)
		if err != nil {
			return nil, err
		}
	}

	ret.Feedbacks, err = s.f.GetByRetroID(ctx, tx, r.ID, params.Get("feedback", nil, logger), logger)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
