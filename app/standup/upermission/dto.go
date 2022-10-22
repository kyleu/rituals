// Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "standup_permission"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"standup_id", "k", "v", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	StandupID uuid.UUID `db:"standup_id"`
	K         string    `db:"k"`
	V         string    `db:"v"`
	Access    string    `db:"access"`
	Created   time.Time `db:"created"`
}

func (d *dto) ToStandupPermission() *StandupPermission {
	if d == nil {
		return nil
	}
	return &StandupPermission{
		StandupID: d.StandupID,
		K:         d.K,
		V:         d.V,
		Access:    d.Access,
		Created:   d.Created,
	}
}

type dtos []*dto

func (x dtos) ToStandupPermissions() StandupPermissions {
	ret := make(StandupPermissions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToStandupPermission())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"standup_id\" = $%d and \"k\" = $%d and \"v\" = $%d", idx+1, idx+2, idx+3)
}
