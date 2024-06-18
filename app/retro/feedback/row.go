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
	ID       uuid.UUID  `db:"id" json:"id"`
	RetroID  uuid.UUID  `db:"retro_id" json:"retro_id"`
	Idx      int        `db:"idx" json:"idx"`
	UserID   uuid.UUID  `db:"user_id" json:"user_id"`
	Category string     `db:"category" json:"category"`
	Content  string     `db:"content" json:"content"`
	HTML     string     `db:"html" json:"html"`
	Created  time.Time  `db:"created" json:"created"`
	Updated  *time.Time `db:"updated" json:"updated"`
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
