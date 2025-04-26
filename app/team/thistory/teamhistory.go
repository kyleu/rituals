package thistory

import (
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/team/history"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*TeamHistory)(nil)

type TeamHistory struct {
	Slug     string    `json:"slug,omitempty"`
	TeamID   uuid.UUID `json:"teamID,omitempty"`
	TeamName string    `json:"teamName,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

func NewTeamHistory(slug string) *TeamHistory {
	return &TeamHistory{Slug: slug}
}

func (t *TeamHistory) Clone() *TeamHistory {
	return &TeamHistory{t.Slug, t.TeamID, t.TeamName, t.Created}
}

func (t *TeamHistory) String() string {
	return t.Slug
}

func (t *TeamHistory) TitleString() string {
	return t.String()
}

func RandomTeamHistory() *TeamHistory {
	return &TeamHistory{
		Slug:     util.RandomString(12),
		TeamID:   util.UUID(),
		TeamName: util.RandomString(12),
		Created:  util.TimeCurrent(),
	}
}

func (t *TeamHistory) Strings() []string {
	return []string{t.Slug, t.TeamID.String(), t.TeamName, util.TimeToFull(&t.Created)}
}

func (t *TeamHistory) ToCSV() ([]string, [][]string) {
	return TeamHistoryFieldDescs.Keys(), [][]string{t.Strings()}
}

func (t *TeamHistory) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(t.Slug))...)
}

func (t *TeamHistory) Breadcrumb(extra ...string) string {
	return t.TitleString() + "||" + t.WebPath(extra...) + "**history"
}

func (t *TeamHistory) ToData() []any {
	return []any{t.Slug, t.TeamID, t.TeamName, t.Created}
}

var TeamHistoryFieldDescs = util.FieldDescs{
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "teamID", Title: "Team ID", Description: "", Type: "uuid"},
	{Key: "teamName", Title: "Team Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
