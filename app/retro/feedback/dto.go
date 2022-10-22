// Content managed by Project Forge, see [projectforge.md] for details.
package feedback

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "feedback"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "retro_id", "idx", "user_id", "category", "content", "html", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
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

func (d *dto) ToFeedback() *Feedback {
	if d == nil {
		return nil
	}
	return &Feedback{
		ID:       d.ID,
		RetroID:  d.RetroID,
		Idx:      d.Idx,
		UserID:   d.UserID,
		Category: d.Category,
		Content:  d.Content,
		HTML:     d.HTML,
		Created:  d.Created,
		Updated:  d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToFeedbacks() Feedbacks {
	ret := make(Feedbacks, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToFeedback())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
