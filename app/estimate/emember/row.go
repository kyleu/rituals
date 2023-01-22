// Content managed by Project Forge, see [projectforge.md] for details.
package emember

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "estimate_member"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"estimate_id", "user_id", "name", "picture", "role", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	EstimateID uuid.UUID         `db:"estimate_id"`
	UserID     uuid.UUID         `db:"user_id"`
	Name       string            `db:"name"`
	Picture    string            `db:"picture"`
	Role       enum.MemberStatus `db:"role"`
	Created    time.Time         `db:"created"`
	Updated    *time.Time        `db:"updated"`
}

func (r *row) ToEstimateMember() *EstimateMember {
	if r == nil {
		return nil
	}
	return &EstimateMember{
		EstimateID: r.EstimateID,
		UserID:     r.UserID,
		Name:       r.Name,
		Picture:    r.Picture,
		Role:       r.Role,
		Created:    r.Created,
		Updated:    r.Updated,
	}
}

type rows []*row

func (x rows) ToEstimateMembers() EstimateMembers {
	ret := make(EstimateMembers, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToEstimateMember())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"estimate_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
