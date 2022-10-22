// Content managed by Project Forge, see [projectforge.md] for details.
package umember

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "standup_member"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"standup_id", "user_id", "name", "picture", "role", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	StandupID uuid.UUID  `db:"standup_id"`
	UserID    uuid.UUID  `db:"user_id"`
	Name      string     `db:"name"`
	Picture   string     `db:"picture"`
	Role      string     `db:"role"`
	Created   time.Time  `db:"created"`
	Updated   *time.Time `db:"updated"`
}

func (d *dto) ToStandupMember() *StandupMember {
	if d == nil {
		return nil
	}
	return &StandupMember{
		StandupID: d.StandupID,
		UserID:    d.UserID,
		Name:      d.Name,
		Picture:   d.Picture,
		Role:      d.Role,
		Created:   d.Created,
		Updated:   d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToStandupMembers() StandupMembers {
	ret := make(StandupMembers, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToStandupMember())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"standup_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
