package email

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/email"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Email)(nil)

type Email struct {
	ID         uuid.UUID     `json:"id,omitzero"`
	Recipients []string      `json:"recipients,omitempty"`
	Subject    string        `json:"subject,omitzero"`
	Data       util.ValueMap `json:"data,omitempty"`
	Plain      string        `json:"plain,omitzero"`
	HTML       string        `json:"html,omitzero"`
	UserID     uuid.UUID     `json:"userID,omitzero"`
	Status     string        `json:"status,omitzero"`
	Created    time.Time     `json:"created,omitzero"`
}

func NewEmail(id uuid.UUID) *Email {
	return &Email{ID: id}
}

func (e *Email) Clone() *Email {
	return &Email{
		ID: e.ID, Recipients: util.ArrayCopy(e.Recipients), Subject: e.Subject, Data: e.Data.Clone(), Plain: e.Plain,
		HTML: e.HTML, UserID: e.UserID, Status: e.Status, Created: e.Created,
	}
}

func (e *Email) String() string {
	return e.ID.String()
}

func (e *Email) TitleString() string {
	return e.String()
}

func RandomEmail() *Email {
	return &Email{
		ID:         util.UUID(),
		Recipients: []string{util.RandomString(12), util.RandomString(12)},
		Subject:    util.RandomString(12),
		Data:       util.RandomValueMap(4),
		Plain:      util.RandomString(12),
		HTML:       "<h3>" + util.RandomString(6) + "</h3>",
		UserID:     util.UUID(),
		Status:     util.RandomString(12),
		Created:    util.TimeCurrent(),
	}
}

//nolint:lll
func (e *Email) Strings() []string {
	return []string{e.ID.String(), util.ToJSONCompact(e.Recipients), e.Subject, util.ToJSONCompact(e.Data), e.Plain, e.HTML, e.UserID.String(), e.Status, util.TimeToFull(&e.Created)}
}

func (e *Email) ToCSV() ([]string, [][]string) {
	return EmailFieldDescs.Keys(), [][]string{e.Strings()}
}

func (e *Email) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(e.ID.String()))...)
}

func (e *Email) Breadcrumb(extra ...string) string {
	return e.TitleString() + "||" + e.WebPath(extra...) + "**email"
}

func (e *Email) ToData() []any {
	return []any{e.ID, e.Recipients, e.Subject, e.Data, e.Plain, e.HTML, e.UserID, e.Status, e.Created}
}

var EmailFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "recipients", Title: "Recipients", Type: "[]string"},
	{Key: "subject", Title: "Subject", Type: "string"},
	{Key: "data", Title: "Data", Type: "map"},
	{Key: "plain", Title: "Plain", Type: "string"},
	{Key: "html", Title: "HTML", Type: "string"},
	{Key: "userID", Title: "User ID", Type: "uuid"},
	{Key: "status", Title: "Status", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
}
