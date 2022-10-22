// Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "team_history"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"slug", "team_id", "team_name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	Slug     string    `db:"slug"`
	TeamID   uuid.UUID `db:"team_id"`
	TeamName string    `db:"team_name"`
	Created  time.Time `db:"created"`
}

func (d *dto) ToTeamHistory() *TeamHistory {
	if d == nil {
		return nil
	}
	return &TeamHistory{
		Slug:     d.Slug,
		TeamID:   d.TeamID,
		TeamName: d.TeamName,
		Created:  d.Created,
	}
}

type dtos []*dto

func (x dtos) ToTeamHistories() TeamHistories {
	ret := make(TeamHistories, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToTeamHistory())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}
