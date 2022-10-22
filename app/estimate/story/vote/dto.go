// Content managed by Project Forge, see [projectforge.md] for details.
package vote

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "vote"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"story_id", "user_id", "choice", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	StoryID uuid.UUID  `db:"story_id"`
	UserID  uuid.UUID  `db:"user_id"`
	Choice  string     `db:"choice"`
	Created time.Time  `db:"created"`
	Updated *time.Time `db:"updated"`
}

func (d *dto) ToVote() *Vote {
	if d == nil {
		return nil
	}
	return &Vote{
		StoryID: d.StoryID,
		UserID:  d.UserID,
		Choice:  d.Choice,
		Created: d.Created,
		Updated: d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToVotes() Votes {
	ret := make(Votes, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToVote())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"story_id\" = $%d and \"user_id\" = $%d", idx+1, idx+2)
}
