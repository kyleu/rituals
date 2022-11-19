// Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "sprint_permission"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"sprint_id", "key", "value", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	SprintID uuid.UUID `db:"sprint_id"`
	Key      string    `db:"key"`
	Value    string    `db:"value"`
	Access   string    `db:"access"`
	Created  time.Time `db:"created"`
}

func (d *dto) ToSprintPermission() *SprintPermission {
	if d == nil {
		return nil
	}
	return &SprintPermission{
		SprintID: d.SprintID,
		Key:      d.Key,
		Value:    d.Value,
		Access:   d.Access,
		Created:  d.Created,
	}
}

type dtos []*dto

func (x dtos) ToSprintPermissions() SprintPermissions {
	ret := make(SprintPermissions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToSprintPermission())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"sprint_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}
