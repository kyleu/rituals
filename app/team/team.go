package team

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/team"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Team)(nil)

type Team struct {
	ID      uuid.UUID          `json:"id,omitzero"`
	Slug    string             `json:"slug,omitzero"`
	Title   string             `json:"title,omitzero"`
	Icon    string             `json:"icon,omitzero"`
	Status  enum.SessionStatus `json:"status,omitzero"`
	Created time.Time          `json:"created,omitzero"`
	Updated *time.Time         `json:"updated,omitzero"`
}

func NewTeam(id uuid.UUID) *Team {
	return &Team{ID: id}
}

func (t *Team) Clone() *Team {
	if t == nil {
		return nil
	}
	return &Team{
		ID: t.ID, Slug: t.Slug, Title: t.Title, Icon: t.Icon, Status: t.Status, Created: t.Created, Updated: t.Updated,
	}
}

func (t *Team) String() string {
	return t.ID.String()
}

func (t *Team) TitleString() string {
	if xx := t.Title; xx != "" {
		return xx
	}
	return t.String()
}

func RandomTeam() *Team {
	return &Team{
		ID:      util.UUID(),
		Slug:    util.RandomString(12),
		Title:   util.RandomString(12),
		Icon:    util.RandomString(12),
		Status:  enum.AllSessionStatuses.Random(),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func (t *Team) Strings() []string {
	return []string{t.ID.String(), t.Slug, t.Title, t.Icon, t.Status.String(), util.TimeToFull(&t.Created), util.TimeToFull(t.Updated)}
}

func (t *Team) ToCSV() ([]string, [][]string) {
	return TeamFieldDescs.Keys(), [][]string{t.Strings()}
}

func (t *Team) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(t.ID.String()))...)
}

func (t *Team) Breadcrumb(extra ...string) string {
	return t.TitleString() + "||" + t.WebPath(extra...) + "**team"
}

func (t *Team) ToData() []any {
	return []any{t.ID, t.Slug, t.Title, t.Icon, t.Status, t.Created, t.Updated}
}

var TeamFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "slug", Title: "Slug", Type: "string"},
	{Key: "title", Title: "Title", Type: "string"},
	{Key: "icon", Title: "Icon", Type: "string"},
	{Key: "status", Title: "Status", Type: "enum(session_status)"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
}
