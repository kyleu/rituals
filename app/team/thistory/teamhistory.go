package thistory

import (
	"net/url"
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
	return util.StringPath(paths...)
}

var _ svc.Model = (*TeamHistory)(nil)

type TeamHistory struct {
	Slug     string    `json:"slug,omitzero"`
	TeamID   uuid.UUID `json:"teamID,omitzero"`
	TeamName string    `json:"teamName,omitzero"`
	Created  time.Time `json:"created,omitzero"`
}

func NewTeamHistory(slug string) *TeamHistory {
	return &TeamHistory{Slug: slug}
}

func (t *TeamHistory) Clone() *TeamHistory {
	if t == nil {
		return nil
	}
	return &TeamHistory{Slug: t.Slug, TeamID: t.TeamID, TeamName: t.TeamName, Created: t.Created}
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
	return util.StringPath(append(paths, url.QueryEscape(t.Slug))...)
}

func (t *TeamHistory) Breadcrumb(extra ...string) string {
	return t.TitleString() + "||" + t.WebPath(extra...) + "**history"
}

func (t *TeamHistory) ToData() []any {
	return []any{t.Slug, t.TeamID, t.TeamName, t.Created}
}

var TeamHistoryFieldDescs = util.FieldDescs{
	{Key: "slug", Title: "Slug", Type: "string"},
	{Key: "teamID", Title: "Team ID", Type: "uuid"},
	{Key: "teamName", Title: "Team Name", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
}
