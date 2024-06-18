package uhistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

var _ svc.Model = (*StandupHistory)(nil)

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

func (s *StandupHistory) Strings() []string {
	return []string{s.Slug, s.StandupID.String(), s.StandupName, util.TimeToFull(&s.Created)}
}

func (s *StandupHistory) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *StandupHistory) WebPath() string {
	return "/admin/db/standup/history/" + url.QueryEscape(s.Slug)
}

func (s *StandupHistory) ToData() []any {
	return []any{s.Slug, s.StandupID, s.StandupName, s.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "standupID", Title: "Standup ID", Description: "", Type: "uuid"},
	{Key: "standupName", Title: "Standup Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
