// Package uhistory - Content managed by Project Forge, see [projectforge.md] for details.
package uhistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type StandupHistory struct {
	Slug        string    `json:"slug,omitempty"`
	StandupID   uuid.UUID `json:"standupID,omitempty"`
	StandupName string    `json:"standupName,omitempty"`
	Created     time.Time `json:"created,omitempty"`
}

func New(slug string) *StandupHistory {
	return &StandupHistory{Slug: slug}
}

func (s *StandupHistory) Clone() *StandupHistory {
	return &StandupHistory{s.Slug, s.StandupID, s.StandupName, s.Created}
}

func (s *StandupHistory) String() string {
	return s.Slug
}

func (s *StandupHistory) TitleString() string {
	return s.String()
}

func Random() *StandupHistory {
	return &StandupHistory{
		Slug:        util.RandomString(12),
		StandupID:   util.UUID(),
		StandupName: util.RandomString(12),
		Created:     util.TimeCurrent(),
	}
}

func (s *StandupHistory) WebPath() string {
	return "/admin/db/standup/history/" + url.QueryEscape(s.Slug)
}

func (s *StandupHistory) ToData() []any {
	return []any{s.Slug, s.StandupID, s.StandupName, s.Created}
}
