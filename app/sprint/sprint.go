// Package sprint - Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Sprint struct {
	ID        uuid.UUID          `json:"id,omitempty"`
	Slug      string             `json:"slug,omitempty"`
	Title     string             `json:"title,omitempty"`
	Icon      string             `json:"icon,omitempty"`
	Status    enum.SessionStatus `json:"status,omitempty"`
	TeamID    *uuid.UUID         `json:"teamID,omitempty"`
	StartDate *time.Time         `json:"startDate,omitempty"`
	EndDate   *time.Time         `json:"endDate,omitempty"`
	Created   time.Time          `json:"created,omitempty"`
	Updated   *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Sprint {
	return &Sprint{ID: id}
}

func (s *Sprint) Clone() *Sprint {
	return &Sprint{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.StartDate, s.EndDate, s.Created, s.Updated}
}

func (s *Sprint) String() string {
	return s.ID.String()
}

func (s *Sprint) TitleString() string {
	return s.Title
}

func Random() *Sprint {
	return &Sprint{
		ID:        util.UUID(),
		Slug:      util.RandomString(12),
		Title:     util.RandomString(12),
		Icon:      util.RandomString(12),
		Status:    enum.AllSessionStatuses.Random(),
		TeamID:    util.UUIDP(),
		StartDate: util.TimeCurrentP(),
		EndDate:   util.TimeCurrentP(),
		Created:   util.TimeCurrent(),
		Updated:   util.TimeCurrentP(),
	}
}

func (s *Sprint) WebPath() string {
	return "/admin/db/sprint/" + s.ID.String()
}

func (s *Sprint) ToData() []any {
	return []any{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.StartDate, s.EndDate, s.Created, s.Updated}
}
