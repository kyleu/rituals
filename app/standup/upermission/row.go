// Package upermission - Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "standup_permission"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"standup_id", "key", "value", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	StandupID uuid.UUID `db:"standup_id"`
	Key       string    `db:"key"`
	Value     string    `db:"value"`
	Access    string    `db:"access"`
	Created   time.Time `db:"created"`
}

func (r *row) ToStandupPermission() *StandupPermission {
	if r == nil {
		return nil
	}
	return &StandupPermission{
		StandupID: r.StandupID,
		Key:       r.Key,
		Value:     r.Value,
		Access:    r.Access,
		Created:   r.Created,
	}
}

type rows []*row

func (x rows) ToStandupPermissions() StandupPermissions {
	return lo.Map(x, func(d *row, _ int) *StandupPermission {
		return d.ToStandupPermission()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"standup_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}
