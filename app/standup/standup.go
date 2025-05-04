package standup

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/standup"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Standup)(nil)

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

func NewStandup(id uuid.UUID) *Standup {
	return &Standup{ID: id}
}

func (s *Standup) Clone() *Standup {
	return &Standup{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.SprintID, s.Created, s.Updated}
}

func (s *Standup) String() string {
	return s.ID.String()
}

func (s *Standup) TitleString() string {
	if xx := s.Title; xx != "" {
		return xx
	}
	return s.String()
}

func RandomStandup() *Standup {
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

//nolint:lll
func (s *Standup) Strings() []string {
	return []string{s.ID.String(), s.Slug, s.Title, s.Icon, s.Status.String(), util.StringNullable(s.TeamID), util.StringNullable(s.SprintID), util.TimeToFull(&s.Created), util.TimeToFull(s.Updated)}
}

func (s *Standup) ToCSV() ([]string, [][]string) {
	return StandupFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *Standup) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.ID.String()))...)
}

func (s *Standup) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**standup"
}

func (s *Standup) ToData() []any {
	return []any{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.SprintID, s.Created, s.Updated}
}

var StandupFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "title", Title: "Title", Description: "", Type: "string"},
	{Key: "icon", Title: "Icon", Description: "", Type: "string"},
	{Key: "status", Title: "Status", Description: "", Type: "enum(session_status)"},
	{Key: "teamID", Title: "Team ID", Description: "", Type: "uuid"},
	{Key: "sprintID", Title: "Sprint ID", Description: "", Type: "uuid"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
