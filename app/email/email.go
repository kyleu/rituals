// Content managed by Project Forge, see [projectforge.md] for details.
package email

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type Email struct {
	ID         uuid.UUID     `json:"id"`
	Recipients []string      `json:"recipients"`
	Subject    string        `json:"subject"`
	Data       util.ValueMap `json:"data"`
	Plain      string        `json:"plain"`
	HTML       string        `json:"html"`
	UserID     uuid.UUID     `json:"userID"`
	Status     string        `json:"status"`
	Created    time.Time     `json:"created"`
}

func New(id uuid.UUID) *Email {
	return &Email{ID: id}
}

func Random() *Email {
	return &Email{
		ID:         util.UUID(),
		Recipients: nil,
		Subject:    util.RandomString(12),
		Data:       util.RandomValueMap(4),
		Plain:      util.RandomString(12),
		HTML:       "<h3>" + util.RandomString(6) + "</h3>",
		UserID:     util.UUID(),
		Status:     util.RandomString(12),
		Created:    util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Email, error) {
	ret := &Email{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Recipients, err = m.ParseArrayString("recipients", true, true)
	if err != nil {
		return nil, err
	}
	ret.Subject, err = m.ParseString("subject", true, true)
	if err != nil {
		return nil, err
	}
	ret.Data, err = m.ParseMap("data", true, true)
	if err != nil {
		return nil, err
	}
	ret.Plain, err = m.ParseString("plain", true, true)
	if err != nil {
		return nil, err
	}
	ret.HTML, err = m.ParseString("html", true, true)
	if err != nil {
		return nil, err
	}
	retUserID, e := m.ParseUUID("userID", true, true)
	if e != nil {
		return nil, e
	}
	if retUserID != nil {
		ret.UserID = *retUserID
	}
	ret.Status, err = m.ParseString("status", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (e *Email) Clone() *Email {
	return &Email{e.ID, e.Recipients, e.Subject, e.Data.Clone(), e.Plain, e.HTML, e.UserID, e.Status, e.Created}
}

func (e *Email) String() string {
	return e.ID.String()
}

func (e *Email) TitleString() string {
	return e.String()
}

func (e *Email) WebPath() string {
	return "/admin/db/email/" + e.ID.String()
}

func (e *Email) Diff(ex *Email) util.Diffs {
	var diffs util.Diffs
	if e.ID != ex.ID {
		diffs = append(diffs, util.NewDiff("id", e.ID.String(), ex.ID.String()))
	}
	diffs = append(diffs, util.DiffObjects(e.Recipients, ex.Recipients, "recipients")...)
	if e.Subject != ex.Subject {
		diffs = append(diffs, util.NewDiff("subject", e.Subject, ex.Subject))
	}
	diffs = append(diffs, util.DiffObjects(e.Data, ex.Data, "data")...)
	if e.Plain != ex.Plain {
		diffs = append(diffs, util.NewDiff("plain", e.Plain, ex.Plain))
	}
	if e.HTML != ex.HTML {
		diffs = append(diffs, util.NewDiff("html", e.HTML, ex.HTML))
	}
	if e.UserID != ex.UserID {
		diffs = append(diffs, util.NewDiff("userID", e.UserID.String(), ex.UserID.String()))
	}
	if e.Status != ex.Status {
		diffs = append(diffs, util.NewDiff("status", e.Status, ex.Status))
	}
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}

func (e *Email) ToData() []any {
	return []any{e.ID, e.Recipients, e.Subject, e.Data, e.Plain, e.HTML, e.UserID, e.Status, e.Created}
}
