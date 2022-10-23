package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

type FullSprint struct {
	Sprint      *sprint.Sprint                `json:"sprint"`
	Histories   shistory.SprintHistories      `json:"histories"`
	Members     smember.SprintMembers         `json:"members"`
	Permissions spermission.SprintPermissions `json:"permissions"`
}

func (s *Service) LoadSprint(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*FullSprint, error) {
	bySlug, err := s.s.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no sprint found with slug [%s]", slug)
		}
		s, err := s.s.Get(ctx, nil, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no sprint found with slug [%s]", slug)
		}
		bySlug = sprint.Sprints{s}
	}
	spr := bySlug[0]

	hist, err := s.sh.GetBySprintID(ctx, nil, spr.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	members, err := s.sm.GetBySprintID(ctx, nil, spr.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	perms, err := s.sp.GetBySprintID(ctx, nil, spr.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	return &FullSprint{Sprint: spr, Histories: hist, Members: members, Permissions: perms}, nil
}
