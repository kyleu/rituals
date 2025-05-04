package story

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "story"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "estimate_id", "idx", "user_id", "title", "status", "final_vote", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	ID         uuid.UUID          `db:"id" json:"id"`
	EstimateID uuid.UUID          `db:"estimate_id" json:"estimate_id"`
	Idx        int                `db:"idx" json:"idx"`
	UserID     uuid.UUID          `db:"user_id" json:"user_id"`
	Title      string             `db:"title" json:"title"`
	Status     enum.SessionStatus `db:"status" json:"status"`
	FinalVote  string             `db:"final_vote" json:"final_vote"`
	Created    time.Time          `db:"created" json:"created"`
	Updated    *time.Time         `db:"updated" json:"updated"`
}

func (r *row) ToStory() *Story {
	if r == nil {
		return nil
	}
	return &Story{
		ID:         r.ID,
		EstimateID: r.EstimateID,
		Idx:        r.Idx,
		UserID:     r.UserID,
		Title:      r.Title,
		Status:     r.Status,
		FinalVote:  r.FinalVote,
		Created:    r.Created,
		Updated:    r.Updated,
	}
}

type rows []*row

func (x rows) ToStories() Stories {
	return lo.Map(x, func(d *row, _ int) *Story {
		return d.ToStory()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
