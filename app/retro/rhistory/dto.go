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

type dto struct {
	Slug      string    `db:"slug"`
	RetroID   uuid.UUID `db:"retro_id"`
	RetroName string    `db:"retro_name"`
	Created   time.Time `db:"created"`
}

func (d *dto) ToRetroHistory() *RetroHistory {
	if d == nil {
		return nil
	}
	return &RetroHistory{
		Slug:      d.Slug,
		RetroID:   d.RetroID,
		RetroName: d.RetroName,
		Created:   d.Created,
	}
}

type dtos []*dto

func (x dtos) ToRetroHistories() RetroHistories {
	ret := make(RetroHistories, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRetroHistory())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}
