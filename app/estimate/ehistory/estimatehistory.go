package ehistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/estimate/history"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*EstimateHistory)(nil)

type EstimateHistory struct {
	Slug         string    `json:"slug,omitzero"`
	EstimateID   uuid.UUID `json:"estimateID,omitzero"`
	EstimateName string    `json:"estimateName,omitzero"`
	Created      time.Time `json:"created,omitzero"`
}

func NewEstimateHistory(slug string) *EstimateHistory {
	return &EstimateHistory{Slug: slug}
}

func (e *EstimateHistory) Clone() *EstimateHistory {
	return &EstimateHistory{Slug: e.Slug, EstimateID: e.EstimateID, EstimateName: e.EstimateName, Created: e.Created}
}

func (e *EstimateHistory) String() string {
	return e.Slug
}

func (e *EstimateHistory) TitleString() string {
	return e.String()
}

func RandomEstimateHistory() *EstimateHistory {
	return &EstimateHistory{
		Slug:         util.RandomString(12),
		EstimateID:   util.UUID(),
		EstimateName: util.RandomString(12),
		Created:      util.TimeCurrent(),
	}
}

func (e *EstimateHistory) Strings() []string {
	return []string{e.Slug, e.EstimateID.String(), e.EstimateName, util.TimeToFull(&e.Created)}
}

func (e *EstimateHistory) ToCSV() ([]string, [][]string) {
	return EstimateHistoryFieldDescs.Keys(), [][]string{e.Strings()}
}

func (e *EstimateHistory) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(e.Slug))...)
}

func (e *EstimateHistory) Breadcrumb(extra ...string) string {
	return e.TitleString() + "||" + e.WebPath(extra...) + "**history"
}

func (e *EstimateHistory) ToData() []any {
	return []any{e.Slug, e.EstimateID, e.EstimateName, e.Created}
}

var EstimateHistoryFieldDescs = util.FieldDescs{
	{Key: "slug", Title: "Slug", Description: "", Type: "string"},
	{Key: "estimateID", Title: "Estimate ID", Description: "", Type: "uuid"},
	{Key: "estimateName", Title: "Estimate Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
