// Package vote - Content managed by Project Forge, see [projectforge.md] for details.
package vote

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "vote"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"story_id", "user_id", "choice", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	StoryID uuid.UUID  `db:"story_id"`
	UserID  uuid.UUID  `db:"user_id"`
	Choice  string     `db:"choice"`
	Created time.Time  `db:"created"`
	Updated *time.Time `db:"updated"`
}

func (r *row) ToVote() *Vote {
	if r == nil {
		return nil
	}
	return &Vote{
		StoryID: r.StoryID,
		UserID:  r.UserID,
		Choice:  r.Choice,
		Created: r.Created,
		Updated: r.Updated,
	}
}

type rows []*row

func (x rows) ToVotes() Votes {
	return lo.Map(x, func(d *row, _ int) *Vote {
		return d.ToVote()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"story_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
