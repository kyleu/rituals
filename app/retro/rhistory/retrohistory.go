package rhistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/retro/history"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*RetroHistory)(nil)

type RetroHistory struct {
	Slug      string    `json:"slug,omitzero"`
	RetroID   uuid.UUID `json:"retroID,omitzero"`
	RetroName string    `json:"retroName,omitzero"`
	Created   time.Time `json:"created,omitzero"`
}

func NewRetroHistory(slug string) *RetroHistory {
	return &RetroHistory{Slug: slug}
}

func (r *RetroHistory) Clone() *RetroHistory {
	return &RetroHistory{Slug: r.Slug, RetroID: r.RetroID, RetroName: r.RetroName, Created: r.Created}
}

func (r *RetroHistory) String() string {
	return r.Slug
}

func (r *RetroHistory) TitleString() string {
	return r.String()
}

func RandomRetroHistory() *RetroHistory {
	return &RetroHistory{
		Slug:      util.RandomString(12),
		RetroID:   util.UUID(),
		RetroName: util.RandomString(12),
		Created:   util.TimeCurrent(),
	}
}

func (r *RetroHistory) Strings() []string {
	return []string{r.Slug, r.RetroID.String(), r.RetroName, util.TimeToFull(&r.Created)}
}

func (r *RetroHistory) ToCSV() ([]string, [][]string) {
	return RetroHistoryFieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *RetroHistory) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(r.Slug))...)
}

func (r *RetroHistory) Breadcrumb(extra ...string) string {
	return r.TitleString() + "||" + r.WebPath(extra...) + "**history"
}

func (r *RetroHistory) ToData() []any {
	return []any{r.Slug, r.RetroID, r.RetroName, r.Created}
}

var RetroHistoryFieldDescs = util.FieldDescs{
	{Key: "slug", Title: "Slug", Type: "string"},
	{Key: "retroID", Title: "Retro ID", Type: "uuid"},
	{Key: "retroName", Title: "Retro Name", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
}
