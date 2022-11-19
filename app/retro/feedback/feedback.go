// Content managed by Project Forge, see [projectforge.md] for details.
package feedback

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type Feedback struct {
	ID       uuid.UUID  `json:"id"`
	RetroID  uuid.UUID  `json:"retroID"`
	Idx      int        `json:"idx"`
	UserID   uuid.UUID  `json:"userID"`
	Category string     `json:"category"`
	Content  string     `json:"content"`
	HTML     string     `json:"html"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Feedback {
	return &Feedback{ID: id}
}

func Random() *Feedback {
	return &Feedback{
		ID:       util.UUID(),
		RetroID:  util.UUID(),
		Idx:      util.RandomInt(10000),
		UserID:   util.UUID(),
		Category: util.RandomString(12),
		Content:  util.RandomString(12),
		HTML:     util.RandomString(12),
		Created:  time.Now(),
		Updated:  util.NowPointer(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Feedback, error) {
	ret := &Feedback{}
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
	retRetroID, e := m.ParseUUID("retroID", true, true)
	if e != nil {
		return nil, e
	}
	if retRetroID != nil {
		ret.RetroID = *retRetroID
	}
	ret.Idx, err = m.ParseInt("idx", true, true)
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
	ret.Category, err = m.ParseString("category", true, true)
	if err != nil {
		return nil, err
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

func (f *Feedback) Clone() *Feedback {
	return &Feedback{
		ID:       f.ID,
		RetroID:  f.RetroID,
		Idx:      f.Idx,
		UserID:   f.UserID,
		Category: f.Category,
		Content:  f.Content,
		HTML:     f.HTML,
		Created:  f.Created,
		Updated:  f.Updated,
	}
}

func (f *Feedback) String() string {
	return f.ID.String()
}

func (f *Feedback) TitleString() string {
	return f.String()
}

func (f *Feedback) WebPath() string {
	return "/admin/db/retro/feedback/" + f.ID.String()
}

func (f *Feedback) Diff(fx *Feedback) util.Diffs {
	var diffs util.Diffs
	if f.ID != fx.ID {
		diffs = append(diffs, util.NewDiff("id", f.ID.String(), fx.ID.String()))
	}
	if f.RetroID != fx.RetroID {
		diffs = append(diffs, util.NewDiff("retroID", f.RetroID.String(), fx.RetroID.String()))
	}
	if f.Idx != fx.Idx {
		diffs = append(diffs, util.NewDiff("idx", fmt.Sprint(f.Idx), fmt.Sprint(fx.Idx)))
	}
	if f.UserID != fx.UserID {
		diffs = append(diffs, util.NewDiff("userID", f.UserID.String(), fx.UserID.String()))
	}
	if f.Category != fx.Category {
		diffs = append(diffs, util.NewDiff("category", f.Category, fx.Category))
	}
	if f.Content != fx.Content {
		diffs = append(diffs, util.NewDiff("content", f.Content, fx.Content))
	}
	if f.HTML != fx.HTML {
		diffs = append(diffs, util.NewDiff("html", f.HTML, fx.HTML))
	}
	if f.Created != fx.Created {
		diffs = append(diffs, util.NewDiff("created", f.Created.String(), fx.Created.String()))
	}
	return diffs
}

func (f *Feedback) ToData() []any {
	return []any{f.ID, f.RetroID, f.Idx, f.UserID, f.Category, f.Content, f.HTML, f.Created, f.Updated}
}
