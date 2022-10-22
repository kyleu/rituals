// Content managed by Project Forge, see [projectforge.md] for details.
package rmember

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "retro_member"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"retro_id", "user_id", "name", "picture", "role", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	RetroID uuid.UUID  `db:"retro_id"`
	UserID  uuid.UUID  `db:"user_id"`
	Name    string     `db:"name"`
	Picture string     `db:"picture"`
	Role    string     `db:"role"`
	Created time.Time  `db:"created"`
	Updated *time.Time `db:"updated"`
}

func (d *dto) ToRetroMember() *RetroMember {
	if d == nil {
		return nil
	}
	return &RetroMember{
		RetroID: d.RetroID,
		UserID:  d.UserID,
		Name:    d.Name,
		Picture: d.Picture,
		Role:    d.Role,
		Created: d.Created,
		Updated: d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToRetroMembers() RetroMembers {
	ret := make(RetroMembers, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRetroMember())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"retro_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
