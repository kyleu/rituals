package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

type FullTeam struct {
	Team        *team.Team                  `json:"team"`
	Histories   thistory.TeamHistories      `json:"histories"`
	Members     tmember.TeamMembers         `json:"members"`
	Permissions tpermission.TeamPermissions `json:"permissions"`
}

func (s *Service) LoadTeam(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*FullTeam, error) {
	bySlug, err := s.t.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no team Teams with slug [%s]", slug)
		}
		t, err := s.t.Get(ctx, nil, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no team Teams with slug [%s]", slug)
		}
		bySlug = team.Teams{t}
	}
	t := bySlug[0]

	hist, err := s.th.GetByTeamID(ctx, nil, t.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	members, err := s.tm.GetByTeamID(ctx, nil, t.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	perms, err := s.tp.GetByTeamID(ctx, nil, t.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	return &FullTeam{Team: t, Histories: hist, Members: members, Permissions: perms}, nil
}
