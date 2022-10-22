// Content managed by Project Forge, see [projectforge.md] for details.
package rhistory

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/util"
)

type RetroHistory struct {
	Slug      string    `json:"slug"`
	RetroID   uuid.UUID `json:"retroID"`
	RetroName string    `json:"retroName"`
	Created   time.Time `json:"created"`
}

func New(slug string) *RetroHistory {
	return &RetroHistory{Slug: slug}
}

func Random() *RetroHistory {
	return &RetroHistory{
		Slug:      util.RandomString(12),
		RetroID:   util.UUID(),
		RetroName: util.RandomString(12),
		Created:   time.Now(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*RetroHistory, error) {
	ret := &RetroHistory{}
	var err error
	if setPK {
		ret.Slug, err = m.ParseString("slug", true, true)
		if err != nil {
			return nil, err
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
	ret.RetroName, err = m.ParseString("retroName", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (r *RetroHistory) Clone() *RetroHistory {
	return &RetroHistory{
		Slug:      r.Slug,
		RetroID:   r.RetroID,
		RetroName: r.RetroName,
		Created:   r.Created,
	}
}

func (r *RetroHistory) String() string {
	return r.Slug
}

func (r *RetroHistory) TitleString() string {
	return r.String()
}

func (r *RetroHistory) WebPath() string {
	return "/admin/db/retro/history" + "/" + r.Slug
}

func (r *RetroHistory) Diff(rx *RetroHistory) util.Diffs {
	var diffs util.Diffs
	if r.Slug != rx.Slug {
		diffs = append(diffs, util.NewDiff("slug", r.Slug, rx.Slug))
	}
	if r.RetroID != rx.RetroID {
		diffs = append(diffs, util.NewDiff("retroID", r.RetroID.String(), rx.RetroID.String()))
	}
	if r.RetroName != rx.RetroName {
		diffs = append(diffs, util.NewDiff("retroName", r.RetroName, rx.RetroName))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}

func (r *RetroHistory) ToData() []any {
	return []any{r.Slug, r.RetroID, r.RetroName, r.Created}
}

type RetroHistories []*RetroHistory

func (r RetroHistories) Get(slug string) *RetroHistory {
	for _, x := range r {
		if x.Slug == slug {
			return x
		}
	}
	return nil
}

func (r RetroHistories) Clone() RetroHistories {
	return slices.Clone(r)
}
