package report

import (
	"net/url"
	"path"
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
	return path.Join(paths...)
}

var _ svc.Model = (*Report)(nil)

type Report struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	StandupID uuid.UUID  `json:"standupID,omitempty"`
	Day       time.Time  `json:"day,omitempty"`
	UserID    uuid.UUID  `json:"userID,omitempty"`
	Content   string     `json:"content,omitempty"`
	HTML      string     `json:"html,omitempty"`
	Created   time.Time  `json:"created,omitempty"`
	Updated   *time.Time `json:"updated,omitempty"`
}

func NewReport(id uuid.UUID) *Report {
	return &Report{ID: id}
}

func (r *Report) Clone() *Report {
	return &Report{r.ID, r.StandupID, r.Day, r.UserID, r.Content, r.HTML, r.Created, r.Updated}
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
	return path.Join(append(paths, url.QueryEscape(r.ID.String()))...)
}

func (r *Report) Breadcrumb(extra ...string) string {
	return r.TitleString() + "||" + r.WebPath(extra...) + "**file-alt"
}

func (r *Report) ToData() []any {
	return []any{r.ID, r.StandupID, r.Day, r.UserID, r.Content, r.HTML, r.Created, r.Updated}
}

var ReportFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "standupID", Title: "Standup ID", Description: "", Type: "uuid"},
	{Key: "day", Title: "Day", Description: "", Type: "date"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "content", Title: "Content", Description: "", Type: "string"},
	{Key: "html", Title: "HTML", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
