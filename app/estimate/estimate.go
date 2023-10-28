// Package estimate - Content managed by Project Forge, see [projectforge.md] for details.
package estimate

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Estimate struct {
	ID       uuid.UUID          `json:"id,omitempty"`
	Slug     string             `json:"slug,omitempty"`
	Title    string             `json:"title,omitempty"`
	Icon     string             `json:"icon,omitempty"`
	Status   enum.SessionStatus `json:"status,omitempty"`
	TeamID   *uuid.UUID         `json:"teamID,omitempty"`
	SprintID *uuid.UUID         `json:"sprintID,omitempty"`
	Choices  []string           `json:"choices,omitempty"`
	Created  time.Time          `json:"created,omitempty"`
	Updated  *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Estimate {
	return &Estimate{ID: id}
}

func (e *Estimate) Clone() *Estimate {
	return &Estimate{e.ID, e.Slug, e.Title, e.Icon, e.Status, e.TeamID, e.SprintID, e.Choices, e.Created, e.Updated}
}

func (e *Estimate) String() string {
	return e.ID.String()
}

func (e *Estimate) TitleString() string {
	return e.Title
}

func Random() *Estimate {
	return &Estimate{
		ID:       util.UUID(),
		Slug:     util.RandomString(12),
		Title:    util.RandomString(12),
		Icon:     util.RandomString(12),
		Status:   enum.AllSessionStatuses.Random(),
		TeamID:   util.UUIDP(),
		SprintID: util.UUIDP(),
		Choices:  []string{util.RandomString(12), util.RandomString(12)},
		Created:  util.TimeCurrent(),
		Updated:  util.TimeCurrentP(),
	}
}

func (e *Estimate) WebPath() string {
	return "/admin/db/estimate/" + e.ID.String()
}

func (e *Estimate) ToData() []any {
	return []any{e.ID, e.Slug, e.Title, e.Icon, e.Status, e.TeamID, e.SprintID, e.Choices, e.Created, e.Updated}
}
