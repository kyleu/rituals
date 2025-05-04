package estimate

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "estimate"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "slug", "title", "icon", "status", "team_id", "sprint_id", "choices", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	ID       uuid.UUID          `db:"id" json:"id"`
	Slug     string             `db:"slug" json:"slug"`
	Title    string             `db:"title" json:"title"`
	Icon     string             `db:"icon" json:"icon"`
	Status   enum.SessionStatus `db:"status" json:"status"`
	TeamID   *uuid.UUID         `db:"team_id" json:"team_id"`
	SprintID *uuid.UUID         `db:"sprint_id" json:"sprint_id"`
	Choices  json.RawMessage    `db:"choices" json:"choices"`
	Created  time.Time          `db:"created" json:"created"`
	Updated  *time.Time         `db:"updated" json:"updated"`
}

func (r *row) ToEstimate() *Estimate {
	if r == nil {
		return nil
	}
	var choicesArg []string
	_ = util.FromJSON(r.Choices, &choicesArg)
	return &Estimate{
		ID:       r.ID,
		Slug:     r.Slug,
		Title:    r.Title,
		Icon:     r.Icon,
		Status:   r.Status,
		TeamID:   r.TeamID,
		SprintID: r.SprintID,
		Choices:  choicesArg,
		Created:  r.Created,
		Updated:  r.Updated,
	}
}

type rows []*row

func (x rows) ToEstimates() Estimates {
	return lo.Map(x, func(d *row, _ int) *Estimate {
		return d.ToEstimate()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
