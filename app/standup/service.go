package standup

import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var (
	_ svc.ServiceID[*Standup, Standups, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Standups]                  = (*Service)(nil)
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["standup"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("standup", &filter.Ordering{Column: "created"})
}
