package sprint

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

var _ svc.Model = (*Sprint)(nil)

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
	if xx := s.Title; xx != "" {
		return xx
	}
	return s.String()
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

//nolint:lll
func (s *Sprint) Strings() []string {
	return []string{s.ID.String(), s.Slug, s.Title, s.Icon, s.Status.String(), util.StringNullable(s.TeamID), util.TimeToYMD(s.StartDate), util.TimeToYMD(s.EndDate), util.TimeToFull(&s.Created), util.TimeToFull(s.Updated)}
}

func (s *Sprint) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *Sprint) WebPath() string {
	return "/admin/db/sprint/" + s.ID.String()
}

func (s *Sprint) ToData() []any {
	return []any{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.StartDate, s.EndDate, s.Created, s.Updated}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "title", Title: "Title", Description: "", Type: "string"},
	{Key: "icon", Title: "Icon", Description: "", Type: "string"},
	{Key: "status", Title: "Status", Description: "", Type: "enum(session_status)"},
	{Key: "teamID", Title: "Team ID", Description: "", Type: "uuid"},
	{Key: "startDate", Title: "Start Date", Description: "", Type: "date"},
	{Key: "endDate", Title: "End Date", Description: "", Type: "date"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
