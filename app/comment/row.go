// Content managed by Project Forge, see [projectforge.md] for details.
package comment

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "comment"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "svc", "model_id", "user_id", "content", "html", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID         `db:"id"`
	Svc     enum.ModelService `db:"svc"`
	ModelID uuid.UUID         `db:"model_id"`
	UserID  uuid.UUID         `db:"user_id"`
	Content string            `db:"content"`
	HTML    string            `db:"html"`
	Created time.Time         `db:"created"`
}

func (r *row) ToComment() *Comment {
	if r == nil {
		return nil
	}
	return &Comment{
		ID:      r.ID,
		Svc:     r.Svc,
		ModelID: r.ModelID,
		UserID:  r.UserID,
		Content: r.Content,
		HTML:    r.HTML,
		Created: r.Created,
	}
}

type rows []*row

func (x rows) ToComments() Comments {
	return lo.Map(x, func(d *row, _ int) *Comment {
		return d.ToComment()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
