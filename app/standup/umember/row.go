package umember

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "standup_member"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"standup_id", "user_id", "name", "picture", "role", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	StandupID uuid.UUID         `db:"standup_id" json:"standup_id"`
	UserID    uuid.UUID         `db:"user_id" json:"user_id"`
	Name      string            `db:"name" json:"name"`
	Picture   string            `db:"picture" json:"picture"`
	Role      enum.MemberStatus `db:"role" json:"role"`
	Created   time.Time         `db:"created" json:"created"`
	Updated   *time.Time        `db:"updated" json:"updated"`
}

func (r *row) ToStandupMember() *StandupMember {
	if r == nil {
		return nil
	}
	return &StandupMember{
		StandupID: r.StandupID,
		UserID:    r.UserID,
		Name:      r.Name,
		Picture:   r.Picture,
		Role:      r.Role,
		Created:   r.Created,
		Updated:   r.Updated,
	}
}

type rows []*row

func (x rows) ToStandupMembers() StandupMembers {
	return lo.Map(x, func(d *row, _ int) *StandupMember {
		return d.ToStandupMember()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"standup_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
