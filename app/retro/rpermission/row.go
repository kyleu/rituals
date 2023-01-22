// Content managed by Project Forge, see [projectforge.md] for details.
package rpermission

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

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
	RetroID uuid.UUID `db:"retro_id"`
	Key     string    `db:"key"`
	Value   string    `db:"value"`
	Access  string    `db:"access"`
	Created time.Time `db:"created"`
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
	ret := make(RetroPermissions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRetroPermission())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"retro_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}
