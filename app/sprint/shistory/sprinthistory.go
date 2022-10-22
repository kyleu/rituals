// Content managed by Project Forge, see [projectforge.md] for details.
package shistory

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/util"
)

type SprintHistory struct {
	Slug       string    `json:"slug"`
	SprintID   uuid.UUID `json:"sprintID"`
	SprintName string    `json:"sprintName"`
	Created    time.Time `json:"created"`
}

func New(slug string) *SprintHistory {
	return &SprintHistory{Slug: slug}
}

func Random() *SprintHistory {
	return &SprintHistory{
		Slug:       util.RandomString(12),
		SprintID:   util.UUID(),
		SprintName: util.RandomString(12),
		Created:    time.Now(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*SprintHistory, error) {
	ret := &SprintHistory{}
	var err error
	if setPK {
		ret.Slug, err = m.ParseString("slug", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	retSprintID, e := m.ParseUUID("sprintID", true, true)
	if e != nil {
		return nil, e
	}
	if retSprintID != nil {
		ret.SprintID = *retSprintID
	}
	ret.SprintName, err = m.ParseString("sprintName", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (s *SprintHistory) Clone() *SprintHistory {
	return &SprintHistory{
		Slug:       s.Slug,
		SprintID:   s.SprintID,
		SprintName: s.SprintName,
		Created:    s.Created,
	}
}

func (s *SprintHistory) String() string {
	return s.Slug
}

func (s *SprintHistory) TitleString() string {
	return s.String()
}

func (s *SprintHistory) WebPath() string {
	return "/admin/db/sprint/history" + "/" + s.Slug
}

func (s *SprintHistory) Diff(sx *SprintHistory) util.Diffs {
	var diffs util.Diffs
	if s.Slug != sx.Slug {
		diffs = append(diffs, util.NewDiff("slug", s.Slug, sx.Slug))
	}
	if s.SprintID != sx.SprintID {
		diffs = append(diffs, util.NewDiff("sprintID", s.SprintID.String(), sx.SprintID.String()))
	}
	if s.SprintName != sx.SprintName {
		diffs = append(diffs, util.NewDiff("sprintName", s.SprintName, sx.SprintName))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}

func (s *SprintHistory) ToData() []any {
	return []any{s.Slug, s.SprintID, s.SprintName, s.Created}
}

type SprintHistories []*SprintHistory

func (s SprintHistories) Get(slug string) *SprintHistory {
	for _, x := range s {
		if x.Slug == slug {
			return x
		}
	}
	return nil
}

func (s SprintHistories) Clone() SprintHistories {
	return slices.Clone(s)
}
