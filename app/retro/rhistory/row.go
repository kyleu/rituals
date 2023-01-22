// Content managed by Project Forge, see [projectforge.md] for details.
package rhistory

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "retro_history"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"slug", "retro_id", "retro_name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	Slug      string    `db:"slug"`
	RetroID   uuid.UUID `db:"retro_id"`
	RetroName string    `db:"retro_name"`
	Created   time.Time `db:"created"`
}

func (r *row) ToRetroHistory() *RetroHistory {
	if r == nil {
		return nil
	}
	return &RetroHistory{
		Slug:      r.Slug,
		RetroID:   r.RetroID,
		RetroName: r.RetroName,
		Created:   r.Created,
	}
}

type rows []*row

func (x rows) ToRetroHistories() RetroHistories {
	ret := make(RetroHistories, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRetroHistory())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}
