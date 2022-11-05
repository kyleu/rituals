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

type dto struct {
	TeamID  uuid.UUID         `db:"team_id"`
	UserID  uuid.UUID         `db:"user_id"`
	Name    string            `db:"name"`
	Picture string            `db:"picture"`
	Role    enum.MemberStatus `db:"role"`
	Created time.Time         `db:"created"`
	Updated *time.Time        `db:"updated"`
}

func (d *dto) ToTeamMember() *TeamMember {
	if d == nil {
		return nil
	}
	return &TeamMember{
		TeamID:  d.TeamID,
		UserID:  d.UserID,
		Name:    d.Name,
		Picture: d.Picture,
		Role:    d.Role,
		Created: d.Created,
		Updated: d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToTeamMembers() TeamMembers {
	ret := make(TeamMembers, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToTeamMember())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"team_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
