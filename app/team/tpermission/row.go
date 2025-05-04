package tpermission

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "team_permission"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"team_id", "key", "value", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	TeamID  uuid.UUID `db:"team_id" json:"team_id"`
	Key     string    `db:"key" json:"key"`
	Value   string    `db:"value" json:"value"`
	Access  string    `db:"access" json:"access"`
	Created time.Time `db:"created" json:"created"`
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
	return lo.Map(x, func(d *row, _ int) *TeamPermission {
		return d.ToTeamPermission()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"team_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}
