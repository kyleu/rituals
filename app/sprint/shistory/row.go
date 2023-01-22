// Content managed by Project Forge, see [projectforge.md] for details.
package shistory

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "sprint_history"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"slug", "sprint_id", "sprint_name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	Slug       string    `db:"slug"`
	SprintID   uuid.UUID `db:"sprint_id"`
	SprintName string    `db:"sprint_name"`
	Created    time.Time `db:"created"`
}

func (r *row) ToSprintHistory() *SprintHistory {
	if r == nil {
		return nil
	}
	return &SprintHistory{
		Slug:       r.Slug,
		SprintID:   r.SprintID,
		SprintName: r.SprintName,
		Created:    r.Created,
	}
}

type rows []*row

func (x rows) ToSprintHistories() SprintHistories {
	ret := make(SprintHistories, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToSprintHistory())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}
