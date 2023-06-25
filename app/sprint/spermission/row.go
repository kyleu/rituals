// Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "sprint_permission"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"sprint_id", "key", "value", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	SprintID uuid.UUID `db:"sprint_id"`
	Key      string    `db:"key"`
	Value    string    `db:"value"`
	Access   string    `db:"access"`
	Created  time.Time `db:"created"`
}

func (r *row) ToSprintPermission() *SprintPermission {
	if r == nil {
		return nil
	}
	return &SprintPermission{
		SprintID: r.SprintID,
		Key:      r.Key,
		Value:    r.Value,
		Access:   r.Access,
		Created:  r.Created,
	}
}

type rows []*row

func (x rows) ToSprintPermissions() SprintPermissions {
	return lo.Map(x, func(d *row, _ int) *SprintPermission {
		return d.ToSprintPermission()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"sprint_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}
