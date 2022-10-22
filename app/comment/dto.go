// Content managed by Project Forge, see [projectforge.md] for details.
package comment

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "comment"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "svc", "model_id", "target_type", "target_id", "user_id", "content", "html", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID         uuid.UUID         `db:"id"`
	Svc        enum.ModelService `db:"svc"`
	ModelID    uuid.UUID         `db:"model_id"`
	TargetType string            `db:"target_type"`
	TargetID   uuid.UUID         `db:"target_id"`
	UserID     uuid.UUID         `db:"user_id"`
	Content    string            `db:"content"`
	HTML       string            `db:"html"`
	Created    time.Time         `db:"created"`
}

func (d *dto) ToComment() *Comment {
	if d == nil {
		return nil
	}
	return &Comment{
		ID:         d.ID,
		Svc:        d.Svc,
		ModelID:    d.ModelID,
		TargetType: d.TargetType,
		TargetID:   d.TargetID,
		UserID:     d.UserID,
		Content:    d.Content,
		HTML:       d.HTML,
		Created:    d.Created,
	}
}

type dtos []*dto

func (x dtos) ToComments() Comments {
	ret := make(Comments, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToComment())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
