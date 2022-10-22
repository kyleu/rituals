// Content managed by Project Forge, see [projectforge.md] for details.
package estimate

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "estimate"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "slug", "title", "status", "team_id", "sprint_id", "owner", "choices", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID       uuid.UUID          `db:"id"`
	Slug     string             `db:"slug"`
	Title    string             `db:"title"`
	Status   enum.SessionStatus `db:"status"`
	TeamID   *uuid.UUID         `db:"team_id"`
	SprintID *uuid.UUID         `db:"sprint_id"`
	Owner    uuid.UUID          `db:"owner"`
	Choices  json.RawMessage    `db:"choices"`
	Created  time.Time          `db:"created"`
	Updated  *time.Time         `db:"updated"`
}

func (d *dto) ToEstimate() *Estimate {
	if d == nil {
		return nil
	}
	choicesArg := []string{}
	_ = util.FromJSON(d.Choices, &choicesArg)
	return &Estimate{
		ID:       d.ID,
		Slug:     d.Slug,
		Title:    d.Title,
		Status:   d.Status,
		TeamID:   d.TeamID,
		SprintID: d.SprintID,
		Owner:    d.Owner,
		Choices:  choicesArg,
		Created:  d.Created,
		Updated:  d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToEstimates() Estimates {
	ret := make(Estimates, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToEstimate())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
