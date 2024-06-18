package story

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

var _ svc.Model = (*Story)(nil)

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

//nolint:lll
func (s *Story) Strings() []string {
	return []string{s.ID.String(), s.EstimateID.String(), fmt.Sprint(s.Idx), s.UserID.String(), s.Title, s.Status.String(), s.FinalVote, util.TimeToFull(&s.Created), util.TimeToFull(s.Updated)}
}

func (s *Story) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *Story) WebPath() string {
	return "/admin/db/estimate/story/" + s.ID.String()
}

func (s *Story) ToData() []any {
	return []any{s.ID, s.EstimateID, s.Idx, s.UserID, s.Title, s.Status, s.FinalVote, s.Created, s.Updated}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "estimateID", Title: "Estimate ID", Description: "", Type: "uuid"},
	{Key: "idx", Title: "Idx", Description: "", Type: "int"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "title", Title: "Title", Description: "", Type: "string"},
	{Key: "status", Title: "Status", Description: "", Type: "enum(session_status)"},
	{Key: "finalVote", Title: "Final Vote", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
