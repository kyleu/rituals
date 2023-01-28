// Content managed by Project Forge, see [projectforge.md] for details.
package tpermission

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "team_permission"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"team_id", "key", "value", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	TeamID  uuid.UUID `db:"team_id"`
	Key     string    `db:"key"`
	Value   string    `db:"value"`
	Access  string    `db:"access"`
	Created time.Time `db:"created"`
}

func (r *row) ToTeamPermission() *TeamPermission {
	if r == nil {
		return nil
	}
	return &TeamPermission{
		TeamID:  r.TeamID,
		Key:     r.Key,
		Value:   r.Value,
		Access:  r.Access,
		Created: r.Created,
	}
}

type rows []*row

func (x rows) ToTeamPermissions() TeamPermissions {
	ret := make(TeamPermissions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToTeamPermission())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"team_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}