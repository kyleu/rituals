package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

type FullEstimate struct {
	Estimate    *estimate.Estimate              `json:"estimate"`
	Histories   ehistory.EstimateHistories      `json:"histories"`
	Members     emember.EstimateMembers         `json:"members"`
	Permissions epermission.EstimatePermissions `json:"permissions"`
}

func (s *Service) LoadEstimate(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*FullEstimate, error) {
	bySlug, err := s.e.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no estimate found with slug [%s]", slug)
		}
		e, err := s.e.Get(ctx, nil, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no estimate found with slug [%s]", slug)
		}
		bySlug = estimate.Estimates{e}
	}
	e := bySlug[0]

	hist, err := s.eh.GetByEstimateID(ctx, nil, e.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	members, err := s.em.GetByEstimateID(ctx, nil, e.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	perms, err := s.ep.GetByEstimateID(ctx, nil, e.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	return &FullEstimate{Estimate: e, Histories: hist, Members: members, Permissions: perms}, nil
}
