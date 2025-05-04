package ehistory

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "estimate_history"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"slug", "estimate_id", "estimate_name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	Slug         string    `db:"slug" json:"slug"`
	EstimateID   uuid.UUID `db:"estimate_id" json:"estimate_id"`
	EstimateName string    `db:"estimate_name" json:"estimate_name"`
	Created      time.Time `db:"created" json:"created"`
}

func (r *row) ToEstimateHistory() *EstimateHistory {
	if r == nil {
		return nil
	}
	return &EstimateHistory{
		Slug:         r.Slug,
		EstimateID:   r.EstimateID,
		EstimateName: r.EstimateName,
		Created:      r.Created,
	}
}

type rows []*row

func (x rows) ToEstimateHistories() EstimateHistories {
	return lo.Map(x, func(d *row, _ int) *EstimateHistory {
		return d.ToEstimateHistory()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}
