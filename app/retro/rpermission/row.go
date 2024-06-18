package rpermission

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "retro_permission"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"retro_id", "key", "value", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	RetroID uuid.UUID `db:"retro_id" json:"retro_id"`
	Key     string    `db:"key" json:"key"`
	Value   string    `db:"value" json:"value"`
	Access  string    `db:"access" json:"access"`
	Created time.Time `db:"created" json:"created"`
}

func (r *row) ToRetroPermission() *RetroPermission {
	if r == nil {
		return nil
	}
	return &RetroPermission{
		RetroID: r.RetroID,
		Key:     r.Key,
		Value:   r.Value,
		Access:  r.Access,
		Created: r.Created,
	}
}

type rows []*row

func (x rows) ToRetroPermissions() RetroPermissions {
	return lo.Map(x, func(d *row, _ int) *RetroPermission {
		return d.ToRetroPermission()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"retro_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}
