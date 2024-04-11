// Package action - Content managed by Project Forge, see [projectforge.md] for details.
package action

import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var _ svc.ServiceID[*Action, Actions, uuid.UUID] = (*Service)(nil)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["action"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("action", &filter.Ordering{Column: "created"})
}
