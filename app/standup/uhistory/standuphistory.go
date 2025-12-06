package uhistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/standup/history"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*StandupHistory)(nil)

type StandupHistory struct {
	Slug        string    `json:"slug,omitzero"`
	StandupID   uuid.UUID `json:"standupID,omitzero"`
	StandupName string    `json:"standupName,omitzero"`
	Created     time.Time `json:"created,omitzero"`
}

func NewStandupHistory(slug string) *StandupHistory {
	return &StandupHistory{Slug: slug}
}

func (s *StandupHistory) Clone() *StandupHistory {
	if s == nil {
		return nil
	}
	return &StandupHistory{Slug: s.Slug, StandupID: s.StandupID, StandupName: s.StandupName, Created: s.Created}
}

func (s *StandupHistory) String() string {
	return s.Slug
}

func (s *StandupHistory) TitleString() string {
	return s.String()
}

func RandomStandupHistory() *StandupHistory {
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
	return StandupHistoryFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *StandupHistory) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.Slug))...)
}

func (s *StandupHistory) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**history"
}

func (s *StandupHistory) ToData() []any {
	return []any{s.Slug, s.StandupID, s.StandupName, s.Created}
}

var StandupHistoryFieldDescs = util.FieldDescs{
	{Key: "slug", Title: "Slug", Type: "string"},
	{Key: "standupID", Title: "Standup ID", Type: "uuid"},
	{Key: "standupName", Title: "Standup Name", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
}
