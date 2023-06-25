// Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "team_history"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"slug", "team_id", "team_name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	Slug     string    `db:"slug"`
	TeamID   uuid.UUID `db:"team_id"`
	TeamName string    `db:"team_name"`
	Created  time.Time `db:"created"`
}

func (r *row) ToTeamHistory() *TeamHistory {
	if r == nil {
		return nil
	}
	return &TeamHistory{
		Slug:     r.Slug,
		TeamID:   r.TeamID,
		TeamName: r.TeamName,
		Created:  r.Created,
	}
}

type rows []*row

func (x rows) ToTeamHistories() TeamHistories {
	return lo.Map(x, func(d *row, _ int) *TeamHistory {
		return d.ToTeamHistory()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}
