package action

import (
	"encoding/json"
	"fmt"
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
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID         `db:"id" json:"id"`
	Svc     enum.ModelService `db:"svc" json:"svc"`
	ModelID uuid.UUID         `db:"model_id" json:"model_id"`
	UserID  uuid.UUID         `db:"user_id" json:"user_id"`
	Act     string            `db:"act" json:"act"`
	Content json.RawMessage   `db:"content" json:"content"`
	Note    string            `db:"note" json:"note"`
	Created time.Time         `db:"created" json:"created"`
}

func (r *row) ToAction() *Action {
	if r == nil {
		return nil
	}
	contentArg, _ := util.FromJSONMap(r.Content)
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
