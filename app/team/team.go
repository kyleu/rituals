package team

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

var _ svc.Model = (*Team)(nil)

type Team struct {
	ID      uuid.UUID          `json:"id,omitempty"`
	Slug    string             `json:"slug,omitempty"`
	Title   string             `json:"title,omitempty"`
	Icon    string             `json:"icon,omitempty"`
	Status  enum.SessionStatus `json:"status,omitempty"`
	Created time.Time          `json:"created,omitempty"`
	Updated *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Team {
	return &Team{ID: id}
}

func (t *Team) Clone() *Team {
	return &Team{t.ID, t.Slug, t.Title, t.Icon, t.Status, t.Created, t.Updated}
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

func Random() *Team {
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
	return FieldDescs.Keys(), [][]string{t.Strings()}
}

func (t *Team) WebPath() string {
	return "/admin/db/team/" + t.ID.String()
}

func (t *Team) ToData() []any {
	return []any{t.ID, t.Slug, t.Title, t.Icon, t.Status, t.Created, t.Updated}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "title", Title: "Title", Description: "", Type: "string"},
	{Key: "icon", Title: "Icon", Description: "", Type: "string"},
	{Key: "status", Title: "Status", Description: "", Type: "enum(session_status)"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
