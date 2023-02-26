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
	columns       = []string{"id", "slug", "title", "icon", "status", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID          `db:"id"`
	Slug    string             `db:"slug"`
	Title   string             `db:"title"`
	Icon    string             `db:"icon"`
	Status  enum.SessionStatus `db:"status"`
	Created time.Time          `db:"created"`
	Updated *time.Time         `db:"updated"`
}

func (r *row) ToTeam() *Team {
	if r == nil {
		return nil
	}
	return &Team{
		ID:      r.ID,
		Slug:    r.Slug,
		Title:   r.Title,
		Icon:    r.Icon,
		Status:  r.Status,
		Created: r.Created,
		Updated: r.Updated,
	}
}

type rows []*row

func (x rows) ToTeams() Teams {
	ret := make(Teams, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToTeam())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
