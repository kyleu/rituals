// Content managed by Project Forge, see [projectforge.md] for details.
package email

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "email"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "recipients", "subject", "data", "plain", "html", "user_id", "status", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID         uuid.UUID       `db:"id"`
	Recipients json.RawMessage `db:"recipients"`
	Subject    string          `db:"subject"`
	Data       json.RawMessage `db:"data"`
	Plain      string          `db:"plain"`
	HTML       string          `db:"html"`
	UserID     uuid.UUID       `db:"user_id"`
	Status     string          `db:"status"`
	Created    time.Time       `db:"created"`
}

func (r *row) ToEmail() *Email {
	if r == nil {
		return nil
	}
	recipientsArg := []string{}
	_ = util.FromJSON(r.Recipients, &recipientsArg)
	dataArg := util.ValueMap{}
	_ = util.FromJSON(r.Data, &dataArg)
	return &Email{
		ID:         r.ID,
		Recipients: recipientsArg,
		Subject:    r.Subject,
		Data:       dataArg,
		Plain:      r.Plain,
		HTML:       r.HTML,
		UserID:     r.UserID,
		Status:     r.Status,
		Created:    r.Created,
	}
}

type rows []*row

func (x rows) ToEmails() Emails {
	ret := make(Emails, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToEmail())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
