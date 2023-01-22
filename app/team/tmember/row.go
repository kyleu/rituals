// Content managed by Project Forge, see [projectforge.md] for details.
package tmember

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "team_member"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"team_id", "user_id", "name", "picture", "role", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	TeamID  uuid.UUID         `db:"team_id"`
	UserID  uuid.UUID         `db:"user_id"`
	Name    string            `db:"name"`
	Picture string            `db:"picture"`
	Role    enum.MemberStatus `db:"role"`
	Created time.Time         `db:"created"`
	Updated *time.Time        `db:"updated"`
}

func (r *row) ToTeamMember() *TeamMember {
	if r == nil {
		return nil
	}
	return &TeamMember{
		TeamID:  r.TeamID,
		UserID:  r.UserID,
		Name:    r.Name,
		Picture: r.Picture,
		Role:    r.Role,
		Created: r.Created,
		Updated: r.Updated,
	}
}

type rows []*row

func (x rows) ToTeamMembers() TeamMembers {
	ret := make(TeamMembers, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToTeamMember())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"team_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
