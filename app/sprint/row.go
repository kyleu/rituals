package sprint

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "sprint"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "slug", "title", "icon", "status", "team_id", "start_date", "end_date", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID        uuid.UUID          `db:"id" json:"id"`
	Slug      string             `db:"slug" json:"slug"`
	Title     string             `db:"title" json:"title"`
	Icon      string             `db:"icon" json:"icon"`
	Status    enum.SessionStatus `db:"status" json:"status"`
	TeamID    *uuid.UUID         `db:"team_id" json:"team_id"`
	StartDate *time.Time         `db:"start_date" json:"start_date"`
	EndDate   *time.Time         `db:"end_date" json:"end_date"`
	Created   time.Time          `db:"created" json:"created"`
	Updated   *time.Time         `db:"updated" json:"updated"`
}

func (r *row) ToSprint() *Sprint {
	if r == nil {
		return nil
	}
	return &Sprint{
		ID:        r.ID,
		Slug:      r.Slug,
		Title:     r.Title,
		Icon:      r.Icon,
		Status:    r.Status,
		TeamID:    r.TeamID,
		StartDate: r.StartDate,
		EndDate:   r.EndDate,
		Created:   r.Created,
		Updated:   r.Updated,
	}
}

type rows []*row

func (x rows) ToSprints() Sprints {
	return lo.Map(x, func(d *row, _ int) *Sprint {
		return d.ToSprint()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
