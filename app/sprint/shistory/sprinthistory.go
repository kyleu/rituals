package shistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

var _ svc.Model = (*SprintHistory)(nil)

type SprintHistory struct {
	Slug       string    `json:"slug,omitempty"`
	SprintID   uuid.UUID `json:"sprintID,omitempty"`
	SprintName string    `json:"sprintName,omitempty"`
	Created    time.Time `json:"created,omitempty"`
}

func New(slug string) *SprintHistory {
	return &SprintHistory{Slug: slug}
}

func (s *SprintHistory) Clone() *SprintHistory {
	return &SprintHistory{s.Slug, s.SprintID, s.SprintName, s.Created}
}

func (s *SprintHistory) String() string {
	return s.Slug
}

func (s *SprintHistory) TitleString() string {
	return s.String()
}

func Random() *SprintHistory {
	return &SprintHistory{
		Slug:       util.RandomString(12),
		SprintID:   util.UUID(),
		SprintName: util.RandomString(12),
		Created:    util.TimeCurrent(),
	}
}

func (s *SprintHistory) Strings() []string {
	return []string{s.Slug, s.SprintID.String(), s.SprintName, util.TimeToFull(&s.Created)}
}

func (s *SprintHistory) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *SprintHistory) WebPath() string {
	return "/admin/db/sprint/history/" + url.QueryEscape(s.Slug)
}

func (s *SprintHistory) ToData() []any {
	return []any{s.Slug, s.SprintID, s.SprintName, s.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "sprintID", Title: "Sprint ID", Description: "", Type: "uuid"},
	{Key: "sprintName", Title: "Sprint Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
