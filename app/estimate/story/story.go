// Package story - Content managed by Project Forge, see [projectforge.md] for details.
package story

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Story struct {
	ID         uuid.UUID          `json:"id,omitempty"`
	EstimateID uuid.UUID          `json:"estimateID,omitempty"`
	Idx        int                `json:"idx,omitempty"`
	UserID     uuid.UUID          `json:"userID,omitempty"`
	Title      string             `json:"title,omitempty"`
	Status     enum.SessionStatus `json:"status,omitempty"`
	FinalVote  string             `json:"finalVote,omitempty"`
	Created    time.Time          `json:"created,omitempty"`
	Updated    *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Story {
	return &Story{ID: id}
}

func (s *Story) Clone() *Story {
	return &Story{s.ID, s.EstimateID, s.Idx, s.UserID, s.Title, s.Status, s.FinalVote, s.Created, s.Updated}
}

func (s *Story) String() string {
	return s.ID.String()
}

func (s *Story) TitleString() string {
	return s.Title
}

func Random() *Story {
	return &Story{
		ID:         util.UUID(),
		EstimateID: util.UUID(),
		Idx:        util.RandomInt(10000),
		UserID:     util.UUID(),
		Title:      util.RandomString(12),
		Status:     enum.AllSessionStatuses.Random(),
		FinalVote:  util.RandomString(12),
		Created:    util.TimeCurrent(),
		Updated:    util.TimeCurrentP(),
	}
}

func (s *Story) WebPath() string {
	return "/admin/db/estimate/story/" + s.ID.String()
}

func (s *Story) ToData() []any {
	return []any{s.ID, s.EstimateID, s.Idx, s.UserID, s.Title, s.Status, s.FinalVote, s.Created, s.Updated}
}
