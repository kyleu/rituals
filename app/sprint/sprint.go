package sprint

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/sprint"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Sprint)(nil)

type Sprint struct {
	ID        uuid.UUID          `json:"id,omitzero"`
	Slug      string             `json:"slug,omitzero"`
	Title     string             `json:"title,omitzero"`
	Icon      string             `json:"icon,omitzero"`
	Status    enum.SessionStatus `json:"status,omitzero"`
	TeamID    *uuid.UUID         `json:"teamID,omitzero"`
	StartDate *time.Time         `json:"startDate,omitzero"`
	EndDate   *time.Time         `json:"endDate,omitzero"`
	Created   time.Time          `json:"created,omitzero"`
	Updated   *time.Time         `json:"updated,omitzero"`
}

func NewSprint(id uuid.UUID) *Sprint {
	return &Sprint{ID: id}
}

func (s *Sprint) Clone() *Sprint {
	if s == nil {
		return nil
	}
	return &Sprint{
		ID: s.ID, Slug: s.Slug, Title: s.Title, Icon: s.Icon, Status: s.Status, TeamID: s.TeamID, StartDate: s.StartDate,
		EndDate: s.EndDate, Created: s.Created, Updated: s.Updated,
	}
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

func RandomSprint() *Sprint {
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
	return SprintFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *Sprint) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.ID.String()))...)
}

func (s *Sprint) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**sprint"
}

func (s *Sprint) ToData() []any {
	return []any{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.StartDate, s.EndDate, s.Created, s.Updated}
}

var SprintFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "slug", Title: "Slug", Type: "string"},
	{Key: "title", Title: "Title", Type: "string"},
	{Key: "icon", Title: "Icon", Type: "string"},
	{Key: "status", Title: "Status", Type: "enum(session_status)"},
	{Key: "teamID", Title: "Team ID", Type: "uuid"},
	{Key: "startDate", Title: "Start Date", Type: "date"},
	{Key: "endDate", Title: "End Date", Type: "date"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
}
