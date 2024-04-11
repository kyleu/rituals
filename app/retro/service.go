// Package retro - Content managed by Project Forge, see [projectforge.md] for details.
package retro

import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var (
	_ svc.ServiceID[*Retro, Retros, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Retros]                = (*Service)(nil)
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["retro"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("retro", &filter.Ordering{Column: "created"})
}
