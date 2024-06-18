package shistory

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

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
	Slug       string    `db:"slug" json:"slug"`
	SprintID   uuid.UUID `db:"sprint_id" json:"sprint_id"`
	SprintName string    `db:"sprint_name" json:"sprint_name"`
	Created    time.Time `db:"created" json:"created"`
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
	return lo.Map(x, func(d *row, _ int) *SprintHistory {
		return d.ToSprintHistory()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}
