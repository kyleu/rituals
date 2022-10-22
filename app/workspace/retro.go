package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) LoadRetro(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*retro.Retro, error) {
	bySlug, err := s.r.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		return nil, errors.Errorf("no retro found with slug [%s]", slug)
	}
	return bySlug[0], nil
}
