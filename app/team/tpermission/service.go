package tpermission

import (
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var _ svc.Service[*TeamPermission, TeamPermissions] = (*Service)(nil)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["tpermission"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("tpermission", &filter.Ordering{Column: "created"})
}
