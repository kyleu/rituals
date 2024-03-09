// Package retro - Content managed by Project Forge, see [projectforge.md] for details.
package retro

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "retro"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "slug", "title", "icon", "status", "team_id", "sprint_id", "categories", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID         uuid.UUID          `db:"id" json:"id"`
	Slug       string             `db:"slug" json:"slug"`
	Title      string             `db:"title" json:"title"`
	Icon       string             `db:"icon" json:"icon"`
	Status     enum.SessionStatus `db:"status" json:"status"`
	TeamID     *uuid.UUID         `db:"team_id" json:"team_id"`
	SprintID   *uuid.UUID         `db:"sprint_id" json:"sprint_id"`
	Categories json.RawMessage    `db:"categories" json:"categories"`
	Created    time.Time          `db:"created" json:"created"`
	Updated    *time.Time         `db:"updated" json:"updated"`
}

func (r *row) ToRetro() *Retro {
	if r == nil {
		return nil
	}
	var categoriesArg []string
	_ = util.FromJSON(r.Categories, &categoriesArg)
	return &Retro{
		ID:         r.ID,
		Slug:       r.Slug,
		Title:      r.Title,
		Icon:       r.Icon,
		Status:     r.Status,
		TeamID:     r.TeamID,
		SprintID:   r.SprintID,
		Categories: categoriesArg,
		Created:    r.Created,
		Updated:    r.Updated,
	}
}

type rows []*row

func (x rows) ToRetros() Retros {
	return lo.Map(x, func(d *row, _ int) *Retro {
		return d.ToRetro()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
