// Content managed by Project Forge, see [projectforge.md] for details.
package uhistory

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "standup_history"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"slug", "standup_id", "standup_name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	Slug        string    `db:"slug"`
	StandupID   uuid.UUID `db:"standup_id"`
	StandupName string    `db:"standup_name"`
	Created     time.Time `db:"created"`
}

func (r *row) ToStandupHistory() *StandupHistory {
	if r == nil {
		return nil
	}
	return &StandupHistory{
		Slug:        r.Slug,
		StandupID:   r.StandupID,
		StandupName: r.StandupName,
		Created:     r.Created,
	}
}

type rows []*row

func (x rows) ToStandupHistories() StandupHistories {
	ret := make(StandupHistories, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToStandupHistory())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}
