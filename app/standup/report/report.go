// Package report - Content managed by Project Forge, see [projectforge.md] for details.
package report

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

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

func New(id uuid.UUID) *Report {
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

func Random() *Report {
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

func (r *Report) WebPath() string {
	return "/admin/db/standup/report/" + r.ID.String()
}

func (r *Report) ToData() []any {
	return []any{r.ID, r.StandupID, r.Day, r.UserID, r.Content, r.HTML, r.Created, r.Updated}
}
