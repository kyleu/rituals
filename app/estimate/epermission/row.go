package epermission

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "estimate_permission"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"estimate_id", "key", "value", "access", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	EstimateID uuid.UUID `db:"estimate_id" json:"estimate_id"`
	Key        string    `db:"key" json:"key"`
	Value      string    `db:"value" json:"value"`
	Access     string    `db:"access" json:"access"`
	Created    time.Time `db:"created" json:"created"`
}

func (r *row) ToEstimatePermission() *EstimatePermission {
	if r == nil {
		return nil
	}
	return &EstimatePermission{
		EstimateID: r.EstimateID,
		Key:        r.Key,
		Value:      r.Value,
		Access:     r.Access,
		Created:    r.Created,
	}
}

type rows []*row

func (x rows) ToEstimatePermissions() EstimatePermissions {
	return lo.Map(x, func(d *row, _ int) *EstimatePermission {
		return d.ToEstimatePermission()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"estimate_id\" = $%d and \"key\" = $%d and \"value\" = $%d", idx+1, idx+2, idx+3)
}
