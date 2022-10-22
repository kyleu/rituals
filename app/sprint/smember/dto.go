// Content managed by Project Forge, see [projectforge.md] for details.
package smember

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "sprint_member"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"sprint_id", "user_id", "name", "picture", "role", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	SprintID uuid.UUID  `db:"sprint_id"`
	UserID   uuid.UUID  `db:"user_id"`
	Name     string     `db:"name"`
	Picture  string     `db:"picture"`
	Role     string     `db:"role"`
	Created  time.Time  `db:"created"`
	Updated  *time.Time `db:"updated"`
}

func (d *dto) ToSprintMember() *SprintMember {
	if d == nil {
		return nil
	}
	return &SprintMember{
		SprintID: d.SprintID,
		UserID:   d.UserID,
		Name:     d.Name,
		Picture:  d.Picture,
		Role:     d.Role,
		Created:  d.Created,
		Updated:  d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToSprintMembers() SprintMembers {
	ret := make(SprintMembers, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToSprintMember())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"sprint_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
