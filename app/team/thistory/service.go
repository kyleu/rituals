// Package thistory - Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import (
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var _ svc.ServiceID[*TeamHistory, TeamHistories, string] = (*Service)(nil)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["thistory"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("thistory", &filter.Ordering{Column: "created"})
}
