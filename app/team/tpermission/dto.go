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
	columns       = []string{"team_id", "k", "v", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	TeamID  uuid.UUID `db:"team_id"`
	K       string    `db:"k"`
	V       string    `db:"v"`
	Access  string    `db:"access"`
	Created time.Time `db:"created"`
}

func (d *dto) ToTeamPermission() *TeamPermission {
	if d == nil {
		return nil
	}
	return &TeamPermission{
		TeamID:  d.TeamID,
		K:       d.K,
		V:       d.V,
		Access:  d.Access,
		Created: d.Created,
	}
}

type dtos []*dto

func (x dtos) ToTeamPermissions() TeamPermissions {
	ret := make(TeamPermissions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToTeamPermission())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"team_id\" = $%d and \"k\" = $%d and \"v\" = $%d", idx+1, idx+2, idx+3)
}
