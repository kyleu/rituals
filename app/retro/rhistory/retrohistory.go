package rhistory

import (
	"net/url"
	"path"
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
	return path.Join(paths...)
}

var _ svc.Model = (*RetroHistory)(nil)

type RetroHistory struct {
	Slug      string    `json:"slug,omitempty"`
	RetroID   uuid.UUID `json:"retroID,omitempty"`
	RetroName string    `json:"retroName,omitempty"`
	Created   time.Time `json:"created,omitempty"`
}

func NewRetroHistory(slug string) *RetroHistory {
	return &RetroHistory{Slug: slug}
}

func (r *RetroHistory) Clone() *RetroHistory {
	return &RetroHistory{r.Slug, r.RetroID, r.RetroName, r.Created}
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
	return path.Join(append(paths, url.QueryEscape(r.Slug))...)
}

func (r *RetroHistory) Breadcrumb(extra ...string) string {
	return r.TitleString() + "||" + r.WebPath(extra...) + "**history"
}

func (r *RetroHistory) ToData() []any {
	return []any{r.Slug, r.RetroID, r.RetroName, r.Created}
}

var RetroHistoryFieldDescs = util.FieldDescs{
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "retroID", Title: "Retro ID", Description: "", Type: "uuid"},
	{Key: "retroName", Title: "Retro Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
