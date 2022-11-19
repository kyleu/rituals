// Content managed by Project Forge, see [projectforge.md] for details.
package epermission

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "estimate_permission"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"estimate_id", "key", "value", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	EstimateID uuid.UUID `db:"estimate_id"`
	Key        string    `db:"key"`
	Value      string    `db:"value"`
	Access     string    `db:"access"`
	Created    time.Time `db:"created"`
}

func (d *dto) ToEstimatePermission() *EstimatePermission {
	if d == nil {
		return nil
	}
	return &EstimatePermission{
		EstimateID: d.EstimateID,
		Key:        d.Key,
		Value:      d.Value,
		Access:     d.Access,
		Created:    d.Created,
	}
}

type dtos []*dto

func (x dtos) ToEstimatePermissions() EstimatePermissions {
	ret := make(EstimatePermissions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToEstimatePermission())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"estimate_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}
