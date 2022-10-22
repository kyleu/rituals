// Content managed by Project Forge, see [projectforge.md] for details.
package team

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Team struct {
	ID      uuid.UUID          `json:"id"`
	Slug    string             `json:"slug"`
	Title   string             `json:"title"`
	Status  enum.SessionStatus `json:"status"`
	Owner   uuid.UUID          `json:"owner"`
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
		Status:  enum.SessionStatus(util.RandomString(12)),
		Owner:   util.UUID(),
		Created: time.Now(),
		Updated: util.NowPointer(),
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
	retStatus, err := m.ParseString("status", true, true)
	if err != nil {
		return nil, err
	}
	ret.Status = enum.SessionStatus(retStatus)
	retOwner, e := m.ParseUUID("owner", true, true)
	if e != nil {
		return nil, e
	}
	if retOwner != nil {
		ret.Owner = *retOwner
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (t *Team) Clone() *Team {
	return &Team{
		ID:      t.ID,
		Slug:    t.Slug,
		Title:   t.Title,
		Status:  t.Status,
		Owner:   t.Owner,
		Created: t.Created,
		Updated: t.Updated,
	}
}

func (t *Team) String() string {
	return t.ID.String()
}

func (t *Team) TitleString() string {
	return t.Title
}

func (t *Team) WebPath() string {
	return "/team" + "/" + t.ID.String()
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
	if t.Status != tx.Status {
		diffs = append(diffs, util.NewDiff("status", string(t.Status), string(tx.Status)))
	}
	if t.Owner != tx.Owner {
		diffs = append(diffs, util.NewDiff("owner", t.Owner.String(), tx.Owner.String()))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}

func (t *Team) ToData() []any {
	return []any{t.ID, t.Slug, t.Title, t.Status, t.Owner, t.Created, t.Updated}
}

type Teams []*Team

func (t Teams) Get(id uuid.UUID) *Team {
	for _, x := range t {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (t Teams) Clone() Teams {
	return slices.Clone(t)
}
