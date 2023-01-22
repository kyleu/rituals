// Content managed by Project Forge, see [projectforge.md] for details.
package smember

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "sprint_member"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"sprint_id", "user_id", "name", "picture", "role", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	SprintID uuid.UUID         `db:"sprint_id"`
	UserID   uuid.UUID         `db:"user_id"`
	Name     string            `db:"name"`
	Picture  string            `db:"picture"`
	Role     enum.MemberStatus `db:"role"`
	Created  time.Time         `db:"created"`
	Updated  *time.Time        `db:"updated"`
}

func (r *row) ToSprintMember() *SprintMember {
	if r == nil {
		return nil
	}
	return &SprintMember{
		SprintID: r.SprintID,
		UserID:   r.UserID,
		Name:     r.Name,
		Picture:  r.Picture,
		Role:     r.Role,
		Created:  r.Created,
		Updated:  r.Updated,
	}
}

type rows []*row

func (x rows) ToSprintMembers() SprintMembers {
	ret := make(SprintMembers, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToSprintMember())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"sprint_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
