// Content managed by Project Forge, see [projectforge.md] for details.
package report

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "report"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "standup_id", "day", "user_id", "content", "html", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID        uuid.UUID  `db:"id"`
	StandupID uuid.UUID  `db:"standup_id"`
	Day       time.Time  `db:"day"`
	UserID    uuid.UUID  `db:"user_id"`
	Content   string     `db:"content"`
	HTML      string     `db:"html"`
	Created   time.Time  `db:"created"`
	Updated   *time.Time `db:"updated"`
}

func (d *dto) ToReport() *Report {
	if d == nil {
		return nil
	}
	return &Report{
		ID:        d.ID,
		StandupID: d.StandupID,
		Day:       d.Day,
		UserID:    d.UserID,
		Content:   d.Content,
		HTML:      d.HTML,
		Created:   d.Created,
		Updated:   d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToReports() Reports {
	ret := make(Reports, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToReport())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
