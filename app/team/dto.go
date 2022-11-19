// Content managed by Project Forge, see [projectforge.md] for details.
package team

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "team"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "slug", "title", "icon", "status", "owner", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID      uuid.UUID          `db:"id"`
	Slug    string             `db:"slug"`
	Title   string             `db:"title"`
	Icon    string             `db:"icon"`
	Status  enum.SessionStatus `db:"status"`
	Owner   uuid.UUID          `db:"owner"`
	Created time.Time          `db:"created"`
	Updated *time.Time         `db:"updated"`
}

func (d *dto) ToTeam() *Team {
	if d == nil {
		return nil
	}
	return &Team{
		ID:      d.ID,
		Slug:    d.Slug,
		Title:   d.Title,
		Icon:    d.Icon,
		Status:  d.Status,
		Owner:   d.Owner,
		Created: d.Created,
		Updated: d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToTeams() Teams {
	ret := make(Teams, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToTeam())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
