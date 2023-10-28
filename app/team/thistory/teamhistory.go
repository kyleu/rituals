// Package thistory - Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type TeamHistory struct {
	Slug     string    `json:"slug,omitempty"`
	TeamID   uuid.UUID `json:"teamID,omitempty"`
	TeamName string    `json:"teamName,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

func New(slug string) *TeamHistory {
	return &TeamHistory{Slug: slug}
}

func (t *TeamHistory) Clone() *TeamHistory {
	return &TeamHistory{t.Slug, t.TeamID, t.TeamName, t.Created}
}

func (t *TeamHistory) String() string {
	return t.Slug
}

func (t *TeamHistory) TitleString() string {
	return t.String()
}

func Random() *TeamHistory {
	return &TeamHistory{
		Slug:     util.RandomString(12),
		TeamID:   util.UUID(),
		TeamName: util.RandomString(12),
		Created:  util.TimeCurrent(),
	}
}

func (t *TeamHistory) WebPath() string {
	return "/admin/db/team/history/" + url.QueryEscape(t.Slug)
}

func (t *TeamHistory) ToData() []any {
	return []any{t.Slug, t.TeamID, t.TeamName, t.Created}
}
