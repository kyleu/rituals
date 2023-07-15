// Content managed by Project Forge, see [projectforge.md] for details.
package team

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Team struct {
	ID      uuid.UUID          `json:"id"`
	Slug    string             `json:"slug"`
	Title   string             `json:"title"`
	Icon    string             `json:"icon"`
	Status  enum.SessionStatus `json:"status"`
	Created time.Time          `json:"created"`
	Updated *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Team {
	return &Team{ID: id}
}

func Random() *Team {
	return &Team{
		ID:      util.UUID(),
		Slug:    util.RandomString(12),
		Title:   util.RandomString(12),
		Icon:    util.RandomString(12),
		Status:  enum.SessionStatus(util.RandomString(12)),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Team, error) {
	ret := &Team{}
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
	ret.Slug, err = m.ParseString("slug", true, true)
	if err != nil {
		return nil, err
	}
	ret.Title, err = m.ParseString("title", true, true)
	if err != nil {
		return nil, err
	}
	ret.Icon, err = m.ParseString("icon", true, true)
	if err != nil {
		return nil, err
	}
	retStatus, err := m.ParseString("status", true, true)
	if err != nil {
		return nil, err
	}
	ret.Status = enum.SessionStatus(retStatus)
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (t *Team) Clone() *Team {
	return &Team{t.ID, t.Slug, t.Title, t.Icon, t.Status, t.Created, t.Updated}
}

func (t *Team) String() string {
	return t.ID.String()
}

func (t *Team) TitleString() string {
	return t.Title
}

func (t *Team) WebPath() string {
	return "/admin/db/team/" + t.ID.String()
}

func (t *Team) Diff(tx *Team) util.Diffs {
	var diffs util.Diffs
	if t.ID != tx.ID {
		diffs = append(diffs, util.NewDiff("id", t.ID.String(), tx.ID.String()))
	}
	if t.Slug != tx.Slug {
		diffs = append(diffs, util.NewDiff("slug", t.Slug, tx.Slug))
	}
	if t.Title != tx.Title {
		diffs = append(diffs, util.NewDiff("title", t.Title, tx.Title))
	}
	if t.Icon != tx.Icon {
		diffs = append(diffs, util.NewDiff("icon", t.Icon, tx.Icon))
	}
	if t.Status != tx.Status {
		diffs = append(diffs, util.NewDiff("status", string(t.Status), string(tx.Status)))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}

func (t *Team) ToData() []any {
	return []any{t.ID, t.Slug, t.Title, t.Icon, t.Status, t.Created, t.Updated}
}
