package estimate

import (
	"net/url"
	"path"
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
	return path.Join(paths...)
}

var _ svc.Model = (*Estimate)(nil)

type Estimate struct {
	ID       uuid.UUID          `json:"id,omitempty"`
	Slug     string             `json:"slug,omitempty"`
	Title    string             `json:"title,omitempty"`
	Icon     string             `json:"icon,omitempty"`
	Status   enum.SessionStatus `json:"status,omitempty"`
	TeamID   *uuid.UUID         `json:"teamID,omitempty"`
	SprintID *uuid.UUID         `json:"sprintID,omitempty"`
	Choices  []string           `json:"choices,omitempty"`
	Created  time.Time          `json:"created,omitempty"`
	Updated  *time.Time         `json:"updated,omitempty"`
}

func NewEstimate(id uuid.UUID) *Estimate {
	return &Estimate{ID: id}
}

func (e *Estimate) Clone() *Estimate {
	return &Estimate{
		e.ID, e.Slug, e.Title, e.Icon, e.Status, e.TeamID, e.SprintID, util.ArrayCopy(e.Choices), e.Created, e.Updated,
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
	return path.Join(append(paths, url.QueryEscape(e.ID.String()))...)
}

func (e *Estimate) ToData() []any {
	return []any{e.ID, e.Slug, e.Title, e.Icon, e.Status, e.TeamID, e.SprintID, e.Choices, e.Created, e.Updated}
}

var EstimateFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "title", Title: "Title", Description: "", Type: "string"},
	{Key: "icon", Title: "Icon", Description: "", Type: "string"},
	{Key: "status", Title: "Status", Description: "", Type: "enum(session_status)"},
	{Key: "teamID", Title: "Team ID", Description: "", Type: "uuid"},
	{Key: "sprintID", Title: "Sprint ID", Description: "", Type: "uuid"},
	{Key: "choices", Title: "Choices", Description: "", Type: "[]string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
