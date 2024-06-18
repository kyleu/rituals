package email

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

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
	ID         uuid.UUID       `db:"id" json:"id"`
	Recipients json.RawMessage `db:"recipients" json:"recipients"`
	Subject    string          `db:"subject" json:"subject"`
	Data       json.RawMessage `db:"data" json:"data"`
	Plain      string          `db:"plain" json:"plain"`
	HTML       string          `db:"html" json:"html"`
	UserID     uuid.UUID       `db:"user_id" json:"user_id"`
	Status     string          `db:"status" json:"status"`
	Created    time.Time       `db:"created" json:"created"`
}

func (r *row) ToEmail() *Email {
	if r == nil {
		return nil
	}
	var recipientsArg []string
	_ = util.FromJSON(r.Recipients, &recipientsArg)
	dataArg, _ := util.FromJSONMap(r.Data)
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
	return lo.Map(x, func(d *row, _ int) *Email {
		return d.ToEmail()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
