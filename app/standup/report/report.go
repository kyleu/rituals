package report

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/standup/report"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Report)(nil)

type Report struct {
	ID        uuid.UUID  `json:"id,omitzero"`
	StandupID uuid.UUID  `json:"standupID,omitzero"`
	Day       time.Time  `json:"day,omitzero"`
	UserID    uuid.UUID  `json:"userID,omitzero"`
	Content   string     `json:"content,omitzero"`
	HTML      string     `json:"html,omitzero"`
	Created   time.Time  `json:"created,omitzero"`
	Updated   *time.Time `json:"updated,omitzero"`
}

func NewReport(id uuid.UUID) *Report {
	return &Report{ID: id}
}

func (r *Report) Clone() *Report {
	return &Report{
		ID: r.ID, StandupID: r.StandupID, Day: r.Day, UserID: r.UserID, Content: r.Content, HTML: r.HTML, Created: r.Created,
		Updated: r.Updated,
	}
}

func (r *Report) String() string {
	return r.ID.String()
}

func (r *Report) TitleString() string {
	return r.String()
}

func RandomReport() *Report {
	return &Report{
		ID:        util.UUID(),
		StandupID: util.UUID(),
		Day:       util.TimeCurrent(),
		UserID:    util.UUID(),
		Content:   util.RandomString(12),
		HTML:      "<h3>" + util.RandomString(6) + "</h3>",
		Created:   util.TimeCurrent(),
		Updated:   util.TimeCurrentP(),
	}
}

//nolint:lll
func (r *Report) Strings() []string {
	return []string{r.ID.String(), r.StandupID.String(), util.TimeToYMD(&r.Day), r.UserID.String(), r.Content, r.HTML, util.TimeToFull(&r.Created), util.TimeToFull(r.Updated)}
}

func (r *Report) ToCSV() ([]string, [][]string) {
	return ReportFieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *Report) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(r.ID.String()))...)
}

func (r *Report) Breadcrumb(extra ...string) string {
	return r.TitleString() + "||" + r.WebPath(extra...) + "**file-alt"
}

func (r *Report) ToData() []any {
	return []any{r.ID, r.StandupID, r.Day, r.UserID, r.Content, r.HTML, r.Created, r.Updated}
}

var ReportFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "standupID", Title: "Standup ID", Type: "uuid"},
	{Key: "day", Title: "Day", Type: "date"},
	{Key: "userID", Title: "User ID", Type: "uuid"},
	{Key: "content", Title: "Content", Type: "string"},
	{Key: "html", Title: "HTML", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
}
