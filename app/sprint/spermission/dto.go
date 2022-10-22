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
	columns       = []string{"sprint_id", "k", "v", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	SprintID uuid.UUID `db:"sprint_id"`
	K        string    `db:"k"`
	V        string    `db:"v"`
	Access   string    `db:"access"`
	Created  time.Time `db:"created"`
}

func (d *dto) ToSprintPermission() *SprintPermission {
	if d == nil {
		return nil
	}
	return &SprintPermission{
		SprintID: d.SprintID,
		K:        d.K,
		V:        d.V,
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
	return fmt.Sprintf("\"sprint_id\" = $%d and \"k\" = $%d and \"v\" = $%d", idx+1, idx+2, idx+3)
}
