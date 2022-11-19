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

type dto struct {
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

func (d *dto) ToRetro() *Retro {
	if d == nil {
		return nil
	}
	categoriesArg := []string{}
	_ = util.FromJSON(d.Categories, &categoriesArg)
	return &Retro{
		ID:         d.ID,
		Slug:       d.Slug,
		Title:      d.Title,
		Icon:       d.Icon,
		Status:     d.Status,
		TeamID:     d.TeamID,
		SprintID:   d.SprintID,
		Owner:      d.Owner,
		Categories: categoriesArg,
		Created:    d.Created,
		Updated:    d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToRetros() Retros {
	ret := make(Retros, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRetro())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
