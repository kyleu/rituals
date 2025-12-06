package estimate

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/estimate"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Estimate)(nil)

type Estimate struct {
	ID       uuid.UUID          `json:"id,omitzero"`
	Slug     string             `json:"slug,omitzero"`
	Title    string             `json:"title,omitzero"`
	Icon     string             `json:"icon,omitzero"`
	Status   enum.SessionStatus `json:"status,omitzero"`
	TeamID   *uuid.UUID         `json:"teamID,omitzero"`
	SprintID *uuid.UUID         `json:"sprintID,omitzero"`
	Choices  []string           `json:"choices,omitempty"`
	Created  time.Time          `json:"created,omitzero"`
	Updated  *time.Time         `json:"updated,omitzero"`
}

func NewEstimate(id uuid.UUID) *Estimate {
	return &Estimate{ID: id}
}

func (e *Estimate) Clone() *Estimate {
	if e == nil {
		return nil
	}
	return &Estimate{
		ID: e.ID, Slug: e.Slug, Title: e.Title, Icon: e.Icon, Status: e.Status, TeamID: e.TeamID, SprintID: e.SprintID,
		Choices: util.ArrayCopy(e.Choices), Created: e.Created, Updated: e.Updated,
	}
}

func (e *Estimate) String() string {
	return e.ID.String()
}

func (e *Estimate) TitleString() string {
	if xx := e.Title; xx != "" {
		return xx
	}
	return e.String()
}

func RandomEstimate() *Estimate {
	return &Estimate{
		ID:       util.UUID(),
		Slug:     util.RandomString(12),
		Title:    util.RandomString(12),
		Icon:     util.RandomString(12),
		Status:   enum.AllSessionStatuses.Random(),
		TeamID:   util.UUIDP(),
		SprintID: util.UUIDP(),
		Choices:  []string{util.RandomString(12), util.RandomString(12)},
		Created:  util.TimeCurrent(),
		Updated:  util.TimeCurrentP(),
	}
}

//nolint:lll
func (e *Estimate) Strings() []string {
	return []string{e.ID.String(), e.Slug, e.Title, e.Icon, e.Status.String(), util.StringNullable(e.TeamID), util.StringNullable(e.SprintID), util.ToJSONCompact(e.Choices), util.TimeToFull(&e.Created), util.TimeToFull(e.Updated)}
}

func (e *Estimate) ToCSV() ([]string, [][]string) {
	return EstimateFieldDescs.Keys(), [][]string{e.Strings()}
}

func (e *Estimate) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(e.ID.String()))...)
}

func (e *Estimate) Breadcrumb(extra ...string) string {
	return e.TitleString() + "||" + e.WebPath(extra...) + "**estimate"
}

func (e *Estimate) ToData() []any {
	return []any{e.ID, e.Slug, e.Title, e.Icon, e.Status, e.TeamID, e.SprintID, e.Choices, e.Created, e.Updated}
}

var EstimateFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "slug", Title: "Slug", Type: "string"},
	{Key: "title", Title: "Title", Type: "string"},
	{Key: "icon", Title: "Icon", Type: "string"},
	{Key: "status", Title: "Status", Type: "enum(session_status)"},
	{Key: "teamID", Title: "Team ID", Type: "uuid"},
	{Key: "sprintID", Title: "Sprint ID", Type: "uuid"},
	{Key: "choices", Title: "Choices", Type: "[]string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
}
