package user

import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var (
	_ svc.ServiceID[*User, Users, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Users]               = (*Service)(nil)
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["user"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("user", &filter.Ordering{Column: "created"})
}
