package story

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/estimate/story"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Story)(nil)

type Story struct {
	ID         uuid.UUID          `json:"id,omitzero"`
	EstimateID uuid.UUID          `json:"estimateID,omitzero"`
	Idx        int                `json:"idx,omitzero"`
	UserID     uuid.UUID          `json:"userID,omitzero"`
	Title      string             `json:"title,omitzero"`
	Status     enum.SessionStatus `json:"status,omitzero"`
	FinalVote  string             `json:"finalVote,omitzero"`
	Created    time.Time          `json:"created,omitzero"`
	Updated    *time.Time         `json:"updated,omitzero"`
}

func NewStory(id uuid.UUID) *Story {
	return &Story{ID: id}
}

func (s *Story) Clone() *Story {
	return &Story{
		ID: s.ID, EstimateID: s.EstimateID, Idx: s.Idx, UserID: s.UserID, Title: s.Title, Status: s.Status,
		FinalVote: s.FinalVote, Created: s.Created, Updated: s.Updated,
	}
}

func (s *Story) String() string {
	return s.ID.String()
}

func (s *Story) TitleString() string {
	if xx := s.Title; xx != "" {
		return xx
	}
	return s.String()
}

func RandomStory() *Story {
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
	return StoryFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *Story) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.ID.String()))...)
}

func (s *Story) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**story"
}

func (s *Story) ToData() []any {
	return []any{s.ID, s.EstimateID, s.Idx, s.UserID, s.Title, s.Status, s.FinalVote, s.Created, s.Updated}
}

var StoryFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "estimateID", Title: "Estimate ID", Type: "uuid"},
	{Key: "idx", Title: "Idx", Type: "int"},
	{Key: "userID", Title: "User ID", Type: "uuid"},
	{Key: "title", Title: "Title", Type: "string"},
	{Key: "status", Title: "Status", Type: "enum(session_status)"},
	{Key: "finalVote", Title: "Final Vote", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
}
