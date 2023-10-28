// Package standup - Content managed by Project Forge, see [projectforge.md] for details.
package standup

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Standup struct {
	ID       uuid.UUID          `json:"id,omitempty"`
	Slug     string             `json:"slug,omitempty"`
	Title    string             `json:"title,omitempty"`
	Icon     string             `json:"icon,omitempty"`
	Status   enum.SessionStatus `json:"status,omitempty"`
	TeamID   *uuid.UUID         `json:"teamID,omitempty"`
	SprintID *uuid.UUID         `json:"sprintID,omitempty"`
	Created  time.Time          `json:"created,omitempty"`
	Updated  *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Standup {
	return &Standup{ID: id}
}

func (s *Standup) Clone() *Standup {
	return &Standup{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.SprintID, s.Created, s.Updated}
}

func (s *Standup) String() string {
	return s.ID.String()
}

func (s *Standup) TitleString() string {
	return s.Title
}

func Random() *Standup {
	return &Standup{
		ID:       util.UUID(),
		Slug:     util.RandomString(12),
		Title:    util.RandomString(12),
		Icon:     util.RandomString(12),
		Status:   enum.AllSessionStatuses.Random(),
		TeamID:   util.UUIDP(),
		SprintID: util.UUIDP(),
		Created:  util.TimeCurrent(),
		Updated:  util.TimeCurrentP(),
	}
}

func (s *Standup) WebPath() string {
	return "/admin/db/standup/" + s.ID.String()
}

func (s *Standup) ToData() []any {
	return []any{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.SprintID, s.Created, s.Updated}
}
