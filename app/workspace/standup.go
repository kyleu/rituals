package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) LoadStandup(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*standup.Standup, error) {
	bySlug, err := s.u.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		return nil, errors.Errorf("no standup found with slug [%s]", slug)
	}
	return bySlug[0], nil
}
