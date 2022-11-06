// Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "sprint"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "slug", "title", "status", "team_id", "owner", "start_date", "end_date", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID        uuid.UUID          `db:"id"`
	Slug      string             `db:"slug"`
	Title     string             `db:"title"`
	Status    enum.SessionStatus `db:"status"`
	TeamID    *uuid.UUID         `db:"team_id"`
	Owner     uuid.UUID          `db:"owner"`
	StartDate *time.Time         `db:"start_date"`
	EndDate   *time.Time         `db:"end_date"`
	Created   time.Time          `db:"created"`
	Updated   *time.Time         `db:"updated"`
}

func (d *dto) ToSprint() *Sprint {
	if d == nil {
		return nil
	}
	return &Sprint{
		ID:        d.ID,
		Slug:      d.Slug,
		Title:     d.Title,
		Status:    d.Status,
		TeamID:    d.TeamID,
		Owner:     d.Owner,
		StartDate: d.StartDate,
		EndDate:   d.EndDate,
		Created:   d.Created,
		Updated:   d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToSprints() Sprints {
	ret := make(Sprints, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToSprint())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
