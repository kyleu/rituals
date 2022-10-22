package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) LoadSprint(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*sprint.Sprint, error) {
	bySlug, err := s.s.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		return nil, errors.Errorf("no sprint found with slug [%s]", slug)
	}
	return bySlug[0], nil
}
