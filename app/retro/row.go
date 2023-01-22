// Content managed by Project Forge, see [projectforge.md] for details.
package retro

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "retro"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "slug", "title", "icon", "status", "team_id", "sprint_id", "owner", "categories", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID         uuid.UUID          `db:"id"`
	Slug       string             `db:"slug"`
	Title      string             `db:"title"`
	Icon       string             `db:"icon"`
	Status     enum.SessionStatus `db:"status"`
	TeamID     *uuid.UUID         `db:"team_id"`
	SprintID   *uuid.UUID         `db:"sprint_id"`
	Owner      uuid.UUID          `db:"owner"`
	Categories json.RawMessage    `db:"categories"`
	Created    time.Time          `db:"created"`
	Updated    *time.Time         `db:"updated"`
}

func (r *row) ToRetro() *Retro {
	if r == nil {
		return nil
	}
	categoriesArg := []string{}
	_ = util.FromJSON(r.Categories, &categoriesArg)
	return &Retro{
		ID:         r.ID,
		Slug:       r.Slug,
		Title:      r.Title,
		Icon:       r.Icon,
		Status:     r.Status,
		TeamID:     r.TeamID,
		SprintID:   r.SprintID,
		Owner:      r.Owner,
		Categories: categoriesArg,
		Created:    r.Created,
		Updated:    r.Updated,
	}
}

type rows []*row

func (x rows) ToRetros() Retros {
	ret := make(Retros, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRetro())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
