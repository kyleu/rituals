// Package email - Content managed by Project Forge, see [projectforge.md] for details.
package email

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

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

func (e *Email) WebPath() string {
	return "/admin/db/email/" + e.ID.String()
}

func (e *Email) ToData() []any {
	return []any{e.ID, e.Recipients, e.Subject, e.Data, e.Plain, e.HTML, e.UserID, e.Status, e.Created}
}
