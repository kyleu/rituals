package team

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "team"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "slug", "title", "icon", "status", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID          `db:"id" json:"id"`
	Slug    string             `db:"slug" json:"slug"`
	Title   string             `db:"title" json:"title"`
	Icon    string             `db:"icon" json:"icon"`
	Status  enum.SessionStatus `db:"status" json:"status"`
	Created time.Time          `db:"created" json:"created"`
	Updated *time.Time         `db:"updated" json:"updated"`
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
	return lo.Map(x, func(d *row, _ int) *Team {
		return d.ToTeam()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
