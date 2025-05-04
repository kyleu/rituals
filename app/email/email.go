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
	ID         uuid.UUID     `json:"id,omitempty"`
	Recipients []string      `json:"recipients,omitempty"`
	Subject    string        `json:"subject,omitempty"`
	Data       util.ValueMap `json:"data,omitempty"`
	Plain      string        `json:"plain,omitempty"`
	HTML       string        `json:"html,omitempty"`
	UserID     uuid.UUID     `json:"userID,omitempty"`
	Status     string        `json:"status,omitempty"`
	Created    time.Time     `json:"created,omitempty"`
}

func NewEmail(id uuid.UUID) *Email {
	return &Email{ID: id}
}

func (e *Email) Clone() *Email {
	return &Email{
		e.ID, util.ArrayCopy(e.Recipients), e.Subject, e.Data.Clone(), e.Plain, e.HTML, e.UserID, e.Status, e.Created,
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
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "recipients", Title: "Recipients", Description: "", Type: "[]string"},
	{Key: "subject", Title: "Subject", Description: "", Type: "string"},
	{Key: "data", Title: "Data", Description: "", Type: "map"},
	{Key: "plain", Title: "Plain", Description: "", Type: "string"},
	{Key: "html", Title: "HTML", Description: "", Type: "string"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "status", Title: "Status", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
