package workspace

import (
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
)

type Service struct {
	t *team.Service
	s *sprint.Service
	e *estimate.Service
	u *standup.Service
	r *retro.Service
}

func NewService(t *team.Service, s *sprint.Service, e *estimate.Service, u *standup.Service, r *retro.Service) *Service {
	return &Service{t: t, s: s, e: e, u: u, r: r}
}
