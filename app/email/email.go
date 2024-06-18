package email

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

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

func New(id uuid.UUID) *Email {
	return &Email{ID: id}
}

func (e *Email) Clone() *Email {
	return &Email{e.ID, e.Recipients, e.Subject, e.Data.Clone(), e.Plain, e.HTML, e.UserID, e.Status, e.Created}
}

func (e *Email) String() string {
	return e.ID.String()
}

func (e *Email) TitleString() string {
	return e.String()
}

func Random() *Email {
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
	return []string{e.ID.String(), util.ToJSON(e.Recipients), e.Subject, util.ToJSON(e.Data), e.Plain, e.HTML, e.UserID.String(), e.Status, util.TimeToFull(&e.Created)}
}

func (e *Email) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{e.Strings()}
}

func (e *Email) WebPath() string {
	return "/admin/db/email/" + e.ID.String()
}

func (e *Email) ToData() []any {
	return []any{e.ID, e.Recipients, e.Subject, e.Data, e.Plain, e.HTML, e.UserID, e.Status, e.Created}
}

var FieldDescs = util.FieldDescs{
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
