package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

type FullRetro struct {
	Retro       *retro.Retro                 `json:"retro"`
	Histories   rhistory.RetroHistories      `json:"histories"`
	Members     rmember.RetroMembers         `json:"members"`
	Permissions rpermission.RetroPermissions `json:"permissions"`
}

func (s *Service) LoadRetro(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*FullRetro, error) {
	bySlug, err := s.r.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no retro Retros with slug [%s]", slug)
		}
		r, err := s.r.Get(ctx, nil, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no retro Retros with slug [%s]", slug)
		}
		bySlug = retro.Retros{r}
	}
	r := bySlug[0]

	hist, err := s.rh.GetByRetroID(ctx, nil, r.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	members, err := s.rm.GetByRetroID(ctx, nil, r.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	perms, err := s.rp.GetByRetroID(ctx, nil, r.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	return &FullRetro{Retro: r, Histories: hist, Members: members, Permissions: perms}, nil
}
