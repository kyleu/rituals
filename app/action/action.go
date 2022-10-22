// Content managed by Project Forge, see [projectforge.md] for details.
package action

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Action struct {
	ID      uuid.UUID         `json:"id"`
	Svc     enum.ModelService `json:"svc"`
	ModelID uuid.UUID         `json:"modelID"`
	UserID  uuid.UUID         `json:"userID"`
	Act     string            `json:"act"`
	Content string            `json:"content"`
	Note    string            `json:"note"`
	Created time.Time         `json:"created"`
}

func New(id uuid.UUID) *Action {
	return &Action{ID: id}
}

func Random() *Action {
	return &Action{
		ID:      util.UUID(),
		Svc:     enum.ModelService(util.RandomString(12)),
		ModelID: util.UUID(),
		UserID:  util.UUID(),
		Act:     util.RandomString(12),
		Content: util.RandomString(12),
		Note:    util.RandomString(12),
		Created: time.Now(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Action, error) {
	ret := &Action{}
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
	retSvc, err := m.ParseString("svc", true, true)
	if err != nil {
		return nil, err
	}
	ret.Svc = enum.ModelService(retSvc)
	retModelID, e := m.ParseUUID("modelID", true, true)
	if e != nil {
		return nil, e
	}
	if retModelID != nil {
		ret.ModelID = *retModelID
	}
	retUserID, e := m.ParseUUID("userID", true, true)
	if e != nil {
		return nil, e
	}
	if retUserID != nil {
		ret.UserID = *retUserID
	}
	ret.Act, err = m.ParseString("act", true, true)
	if err != nil {
		return nil, err
	}
	ret.Content, err = m.ParseString("content", true, true)
	if err != nil {
		return nil, err
	}
	ret.Note, err = m.ParseString("note", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (a *Action) Clone() *Action {
	return &Action{
		ID:      a.ID,
		Svc:     a.Svc,
		ModelID: a.ModelID,
		UserID:  a.UserID,
		Act:     a.Act,
		Content: a.Content,
		Note:    a.Note,
		Created: a.Created,
	}
}

func (a *Action) String() string {
	return a.ID.String()
}

func (a *Action) TitleString() string {
	return a.String()
}

func (a *Action) WebPath() string {
	return "/admin/db/action" + "/" + a.ID.String()
}

func (a *Action) Diff(ax *Action) util.Diffs {
	var diffs util.Diffs
	if a.ID != ax.ID {
		diffs = append(diffs, util.NewDiff("id", a.ID.String(), ax.ID.String()))
	}
	if a.Svc != ax.Svc {
		diffs = append(diffs, util.NewDiff("svc", string(a.Svc), string(ax.Svc)))
	}
	if a.ModelID != ax.ModelID {
		diffs = append(diffs, util.NewDiff("modelID", a.ModelID.String(), ax.ModelID.String()))
	}
	if a.UserID != ax.UserID {
		diffs = append(diffs, util.NewDiff("userID", a.UserID.String(), ax.UserID.String()))
	}
	if a.Act != ax.Act {
		diffs = append(diffs, util.NewDiff("act", a.Act, ax.Act))
	}
	if a.Content != ax.Content {
		diffs = append(diffs, util.NewDiff("content", a.Content, ax.Content))
	}
	if a.Note != ax.Note {
		diffs = append(diffs, util.NewDiff("note", a.Note, ax.Note))
	}
	if a.Created != ax.Created {
		diffs = append(diffs, util.NewDiff("created", a.Created.String(), ax.Created.String()))
	}
	return diffs
}

func (a *Action) ToData() []any {
	return []any{a.ID, a.Svc, a.ModelID, a.UserID, a.Act, a.Content, a.Note, a.Created}
}

type Actions []*Action

func (a Actions) Get(id uuid.UUID) *Action {
	for _, x := range a {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (a Actions) Clone() Actions {
	return slices.Clone(a)
}
