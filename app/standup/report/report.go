// Content managed by Project Forge, see [projectforge.md] for details.
package report

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type Report struct {
	ID        uuid.UUID  `json:"id"`
	StandupID uuid.UUID  `json:"standupID"`
	D         time.Time  `json:"d"`
	UserID    uuid.UUID  `json:"userID"`
	Content   string     `json:"content"`
	HTML      string     `json:"html"`
	Created   time.Time  `json:"created"`
	Updated   *time.Time `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Report {
	return &Report{ID: id}
}

func Random() *Report {
	return &Report{
		ID:        util.UUID(),
		StandupID: util.UUID(),
		D:         time.Now(),
		UserID:    util.UUID(),
		Content:   util.RandomString(12),
		HTML:      util.RandomString(12),
		Created:   time.Now(),
		Updated:   util.NowPointer(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Report, error) {
	ret := &Report{}
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
	retStandupID, e := m.ParseUUID("standupID", true, true)
	if e != nil {
		return nil, e
	}
	if retStandupID != nil {
		ret.StandupID = *retStandupID
	}
	retD, e := m.ParseTime("d", true, true)
	if e != nil {
		return nil, e
	}
	if retD != nil {
		ret.D = *retD
	}
	retUserID, e := m.ParseUUID("userID", true, true)
	if e != nil {
		return nil, e
	}
	if retUserID != nil {
		ret.UserID = *retUserID
	}
	ret.Content, err = m.ParseString("content", true, true)
	if err != nil {
		return nil, err
	}
	ret.HTML, err = m.ParseString("html", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (r *Report) Clone() *Report {
	return &Report{
		ID:        r.ID,
		StandupID: r.StandupID,
		D:         r.D,
		UserID:    r.UserID,
		Content:   r.Content,
		HTML:      r.HTML,
		Created:   r.Created,
		Updated:   r.Updated,
	}
}

func (r *Report) String() string {
	return r.ID.String()
}

func (r *Report) TitleString() string {
	return r.String()
}

func (r *Report) WebPath() string {
	return "/admin/db/standup/report" + "/" + r.ID.String()
}

func (r *Report) Diff(rx *Report) util.Diffs {
	var diffs util.Diffs
	if r.ID != rx.ID {
		diffs = append(diffs, util.NewDiff("id", r.ID.String(), rx.ID.String()))
	}
	if r.StandupID != rx.StandupID {
		diffs = append(diffs, util.NewDiff("standupID", r.StandupID.String(), rx.StandupID.String()))
	}
	if r.D != rx.D {
		diffs = append(diffs, util.NewDiff("d", r.D.String(), rx.D.String()))
	}
	if r.UserID != rx.UserID {
		diffs = append(diffs, util.NewDiff("userID", r.UserID.String(), rx.UserID.String()))
	}
	if r.Content != rx.Content {
		diffs = append(diffs, util.NewDiff("content", r.Content, rx.Content))
	}
	if r.HTML != rx.HTML {
		diffs = append(diffs, util.NewDiff("html", r.HTML, rx.HTML))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}

func (r *Report) ToData() []any {
	return []any{r.ID, r.StandupID, r.D, r.UserID, r.Content, r.HTML, r.Created, r.Updated}
}
