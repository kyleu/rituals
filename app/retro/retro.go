package retro

import (
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/retro"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*Retro)(nil)

type Retro struct {
	ID         uuid.UUID          `json:"id,omitempty"`
	Slug       string             `json:"slug,omitempty"`
	Title      string             `json:"title,omitempty"`
	Icon       string             `json:"icon,omitempty"`
	Status     enum.SessionStatus `json:"status,omitempty"`
	TeamID     *uuid.UUID         `json:"teamID,omitempty"`
	SprintID   *uuid.UUID         `json:"sprintID,omitempty"`
	Categories []string           `json:"categories,omitempty"`
	Created    time.Time          `json:"created,omitempty"`
	Updated    *time.Time         `json:"updated,omitempty"`
}

func NewRetro(id uuid.UUID) *Retro {
	return &Retro{ID: id}
}

func (r *Retro) Clone() *Retro {
	return &Retro{r.ID, r.Slug, r.Title, r.Icon, r.Status, r.TeamID, r.SprintID, r.Categories, r.Created, r.Updated}
}

func (r *Retro) String() string {
	return r.ID.String()
}

func (r *Retro) TitleString() string {
	if xx := r.Title; xx != "" {
		return xx
	}
	return r.String()
}

func RandomRetro() *Retro {
	return &Retro{
		ID:         util.UUID(),
		Slug:       util.RandomString(12),
		Title:      util.RandomString(12),
		Icon:       util.RandomString(12),
		Status:     enum.AllSessionStatuses.Random(),
		TeamID:     util.UUIDP(),
		SprintID:   util.UUIDP(),
		Categories: []string{util.RandomString(12), util.RandomString(12)},
		Created:    util.TimeCurrent(),
		Updated:    util.TimeCurrentP(),
	}
}

//nolint:lll
func (r *Retro) Strings() []string {
	return []string{r.ID.String(), r.Slug, r.Title, r.Icon, r.Status.String(), util.StringNullable(r.TeamID), util.StringNullable(r.SprintID), util.ToJSONCompact(r.Categories), util.TimeToFull(&r.Created), util.TimeToFull(r.Updated)}
}

func (r *Retro) ToCSV() ([]string, [][]string) {
	return RetroFieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *Retro) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(r.ID.String()))...)
}

func (r *Retro) ToData() []any {
	return []any{r.ID, r.Slug, r.Title, r.Icon, r.Status, r.TeamID, r.SprintID, r.Categories, r.Created, r.Updated}
}

var RetroFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "title", Title: "Title", Description: "", Type: "string"},
	{Key: "icon", Title: "Icon", Description: "", Type: "string"},
	{Key: "status", Title: "Status", Description: "", Type: "enum(session_status)"},
	{Key: "teamID", Title: "Team ID", Description: "", Type: "uuid"},
	{Key: "sprintID", Title: "Sprint ID", Description: "", Type: "uuid"},
	{Key: "categories", Title: "Categories", Description: "", Type: "[]string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
