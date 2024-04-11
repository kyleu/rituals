// Package estimate - Content managed by Project Forge, see [projectforge.md] for details.
package estimate

import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var (
	_ svc.ServiceID[*Estimate, Estimates, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Estimates]                   = (*Service)(nil)
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["estimate"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("estimate", &filter.Ordering{Column: "created"})
}
