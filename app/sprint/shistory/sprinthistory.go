// Package shistory - Content managed by Project Forge, see [projectforge.md] for details.
package shistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type SprintHistory struct {
	Slug       string    `json:"slug,omitempty"`
	SprintID   uuid.UUID `json:"sprintID,omitempty"`
	SprintName string    `json:"sprintName,omitempty"`
	Created    time.Time `json:"created,omitempty"`
}

func New(slug string) *SprintHistory {
	return &SprintHistory{Slug: slug}
}

func (s *SprintHistory) Clone() *SprintHistory {
	return &SprintHistory{s.Slug, s.SprintID, s.SprintName, s.Created}
}

func (s *SprintHistory) String() string {
	return s.Slug
}

func (s *SprintHistory) TitleString() string {
	return s.String()
}

func Random() *SprintHistory {
	return &SprintHistory{
		Slug:       util.RandomString(12),
		SprintID:   util.UUID(),
		SprintName: util.RandomString(12),
		Created:    util.TimeCurrent(),
	}
}

func (s *SprintHistory) WebPath() string {
	return "/admin/db/sprint/history/" + url.QueryEscape(s.Slug)
}

func (s *SprintHistory) ToData() []any {
	return []any{s.Slug, s.SprintID, s.SprintName, s.Created}
}
