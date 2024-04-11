// Package ehistory - Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import (
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/svc"
)

var _ svc.ServiceID[*EstimateHistory, EstimateHistories, string] = (*Service)(nil)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["ehistory"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("ehistory", &filter.Ordering{Column: "created"})
}
