// Content managed by Project Forge, see [projectforge.md] for details.
package action

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

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

type row struct {
	ID      uuid.UUID         `db:"id"`
	Svc     enum.ModelService `db:"svc"`
	ModelID uuid.UUID         `db:"model_id"`
	UserID  uuid.UUID         `db:"user_id"`
	Act     string            `db:"act"`
	Content json.RawMessage   `db:"content"`
	Note    string            `db:"note"`
	Created time.Time         `db:"created"`
}

func (r *row) ToAction() *Action {
	if r == nil {
		return nil
	}
	contentArg := util.ValueMap{}
	_ = util.FromJSON(r.Content, &contentArg)
	return &Action{
		ID:      r.ID,
		Svc:     r.Svc,
		ModelID: r.ModelID,
		UserID:  r.UserID,
		Act:     r.Act,
		Content: contentArg,
		Note:    r.Note,
		Created: r.Created,
	}
}

type rows []*row

func (x rows) ToActions() Actions {
	return lo.Map(x, func(d *row, _ int) *Action {
		return d.ToAction()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
