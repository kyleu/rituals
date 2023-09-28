// Package feedback - Content managed by Project Forge, see [projectforge.md] for details.
package feedback

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "feedback"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "retro_id", "idx", "user_id", "category", "content", "html", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID       uuid.UUID  `db:"id"`
	RetroID  uuid.UUID  `db:"retro_id"`
	Idx      int        `db:"idx"`
	UserID   uuid.UUID  `db:"user_id"`
	Category string     `db:"category"`
	Content  string     `db:"content"`
	HTML     string     `db:"html"`
	Created  time.Time  `db:"created"`
	Updated  *time.Time `db:"updated"`
}

func (r *row) ToFeedback() *Feedback {
	if r == nil {
		return nil
	}
	return &Feedback{
		ID:       r.ID,
		RetroID:  r.RetroID,
		Idx:      r.Idx,
		UserID:   r.UserID,
		Category: r.Category,
		Content:  r.Content,
		HTML:     r.HTML,
		Created:  r.Created,
		Updated:  r.Updated,
	}
}

type rows []*row

func (x rows) ToFeedbacks() Feedbacks {
	return lo.Map(x, func(d *row, _ int) *Feedback {
		return d.ToFeedback()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
