// Content managed by Project Forge, see [projectforge.md] for details.
package standup

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "standup"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "slug", "title", "icon", "status", "team_id", "sprint_id", "owner", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID       uuid.UUID          `db:"id"`
	Slug     string             `db:"slug"`
	Title    string             `db:"title"`
	Icon     string             `db:"icon"`
	Status   enum.SessionStatus `db:"status"`
	TeamID   *uuid.UUID         `db:"team_id"`
	SprintID *uuid.UUID         `db:"sprint_id"`
	Owner    uuid.UUID          `db:"owner"`
	Created  time.Time          `db:"created"`
	Updated  *time.Time         `db:"updated"`
}

func (r *row) ToStandup() *Standup {
	if r == nil {
		return nil
	}
	return &Standup{
		ID:       r.ID,
		Slug:     r.Slug,
		Title:    r.Title,
		Icon:     r.Icon,
		Status:   r.Status,
		TeamID:   r.TeamID,
		SprintID: r.SprintID,
		Owner:    r.Owner,
		Created:  r.Created,
		Updated:  r.Updated,
	}
}

type rows []*row

func (x rows) ToStandups() Standups {
	ret := make(Standups, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToStandup())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
