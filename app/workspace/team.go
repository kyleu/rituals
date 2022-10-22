package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) LoadTeam(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*team.Team, error) {
	bySlug, err := s.t.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		return nil, errors.Errorf("no team found with slug [%s]", slug)
	}
	return bySlug[0], nil
}
