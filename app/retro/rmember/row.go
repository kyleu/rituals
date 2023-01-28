// Content managed by Project Forge, see [projectforge.md] for details.
package rmember

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "retro_member"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"retro_id", "user_id", "name", "picture", "role", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	RetroID uuid.UUID         `db:"retro_id"`
	UserID  uuid.UUID         `db:"user_id"`
	Name    string            `db:"name"`
	Picture string            `db:"picture"`
	Role    enum.MemberStatus `db:"role"`
	Created time.Time         `db:"created"`
	Updated *time.Time        `db:"updated"`
}

func (r *row) ToRetroMember() *RetroMember {
	if r == nil {
		return nil
	}
	return &RetroMember{
		RetroID: r.RetroID,
		UserID:  r.UserID,
		Name:    r.Name,
		Picture: r.Picture,
		Role:    r.Role,
		Created: r.Created,
		Updated: r.Updated,
	}
}

type rows []*row

func (x rows) ToRetroMembers() RetroMembers {
	ret := make(RetroMembers, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRetroMember())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"retro_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}