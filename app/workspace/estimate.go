package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) LoadEstimate(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*estimate.Estimate, error) {
	bySlug, err := s.e.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		return nil, errors.Errorf("no estimate found with slug [%s]", slug)
	}
	return bySlug[0], nil
}
