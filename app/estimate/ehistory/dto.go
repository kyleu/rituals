// Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "estimate_history"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"slug", "estimate_id", "estimate_name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	Slug         string    `db:"slug"`
	EstimateID   uuid.UUID `db:"estimate_id"`
	EstimateName string    `db:"estimate_name"`
	Created      time.Time `db:"created"`
}

func (d *dto) ToEstimateHistory() *EstimateHistory {
	if d == nil {
		return nil
	}
	return &EstimateHistory{
		Slug:         d.Slug,
		EstimateID:   d.EstimateID,
		EstimateName: d.EstimateName,
		Created:      d.Created,
	}
}

type dtos []*dto

func (x dtos) ToEstimateHistories() EstimateHistories {
	ret := make(EstimateHistories, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToEstimateHistory())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"slug\" = $%d", idx+1)
}