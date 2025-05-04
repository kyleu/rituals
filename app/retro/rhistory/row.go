package rhistory

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "retro_history"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"slug", "retro_id", "retro_name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	Slug      string    `db:"slug" json:"slug"`
	RetroID   uuid.UUID `db:"retro_id" json:"retro_id"`
	RetroName string    `db:"retro_name" json:"retro_name"`
	Created   time.Time `db:"created" json:"created"`
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
	return lo.Map(x, func(d *row, _ int) *RetroHistory {
		return d.ToRetroHistory()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}
