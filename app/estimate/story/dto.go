// Content managed by Project Forge, see [projectforge.md] for details.
package story

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "story"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "estimate_id", "idx", "user_id", "title", "status", "final_vote", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID         uuid.UUID          `db:"id"`
	EstimateID uuid.UUID          `db:"estimate_id"`
	Idx        int                `db:"idx"`
	UserID     uuid.UUID          `db:"user_id"`
	Title      string             `db:"title"`
	Status     enum.SessionStatus `db:"status"`
	FinalVote  string             `db:"final_vote"`
	Created    time.Time          `db:"created"`
	Updated    *time.Time         `db:"updated"`
}

func (d *dto) ToStory() *Story {
	if d == nil {
		return nil
	}
	return &Story{
		ID:         d.ID,
		EstimateID: d.EstimateID,
		Idx:        d.Idx,
		UserID:     d.UserID,
		Title:      d.Title,
		Status:     d.Status,
		FinalVote:  d.FinalVote,
		Created:    d.Created,
		Updated:    d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToStories() Stories {
	ret := make(Stories, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToStory())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
