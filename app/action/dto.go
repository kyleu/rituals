// Content managed by Project Forge, see [projectforge.md] for details.
package action

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "action"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "svc", "model_id", "user_id", "act", "content", "note", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID      uuid.UUID         `db:"id"`
	Svc     enum.ModelService `db:"svc"`
	ModelID uuid.UUID         `db:"model_id"`
	UserID  uuid.UUID         `db:"user_id"`
	Act     string            `db:"act"`
	Content string            `db:"content"`
	Note    string            `db:"note"`
	Created time.Time         `db:"created"`
}

func (d *dto) ToAction() *Action {
	if d == nil {
		return nil
	}
	return &Action{
		ID:      d.ID,
		Svc:     d.Svc,
		ModelID: d.ModelID,
		UserID:  d.UserID,
		Act:     d.Act,
		Content: d.Content,
		Note:    d.Note,
		Created: d.Created,
	}
}

type dtos []*dto

func (x dtos) ToActions() Actions {
	ret := make(Actions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToAction())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
