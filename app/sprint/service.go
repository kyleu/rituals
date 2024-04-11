// Package sprint - Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var (
	_ svc.ServiceID[*Sprint, Sprints, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Sprints]                 = (*Service)(nil)
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["sprint"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("sprint", &filter.Ordering{Column: "created"})
}
