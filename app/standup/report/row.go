// Content managed by Project Forge, see [projectforge.md] for details.
package report

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "report"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "standup_id", "day", "user_id", "content", "html", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID        uuid.UUID  `db:"id"`
	StandupID uuid.UUID  `db:"standup_id"`
	Day       time.Time  `db:"day"`
	UserID    uuid.UUID  `db:"user_id"`
	Content   string     `db:"content"`
	HTML      string     `db:"html"`
	Created   time.Time  `db:"created"`
	Updated   *time.Time `db:"updated"`
}

func (r *row) ToReport() *Report {
	if r == nil {
		return nil
	}
	return &Report{
		ID:        r.ID,
		StandupID: r.StandupID,
		Day:       r.Day,
		UserID:    r.UserID,
		Content:   r.Content,
		HTML:      r.HTML,
		Created:   r.Created,
		Updated:   r.Updated,
	}
}

type rows []*row

func (x rows) ToReports() Reports {
	return lo.Map(x, func(d *row, _ int) *Report {
		return d.ToReport()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
