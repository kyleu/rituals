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

type dto struct {
	RetroID uuid.UUID `db:"retro_id"`
	Key     string    `db:"key"`
	Value   string    `db:"value"`
	Access  string    `db:"access"`
	Created time.Time `db:"created"`
}

func (d *dto) ToRetroPermission() *RetroPermission {
	if d == nil {
		return nil
	}
	return &RetroPermission{
		RetroID: d.RetroID,
		Key:     d.Key,
		Value:   d.Value,
		Access:  d.Access,
		Created: d.Created,
	}
}

type dtos []*dto

func (x dtos) ToRetroPermissions() RetroPermissions {
	ret := make(RetroPermissions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRetroPermission())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"retro_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}
