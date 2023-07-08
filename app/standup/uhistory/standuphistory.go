// Content managed by Project Forge, see [projectforge.md] for details.
package uhistory

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type StandupHistory struct {
	Slug        string    `json:"slug"`
	StandupID   uuid.UUID `json:"standupID"`
	StandupName string    `json:"standupName"`
	Created     time.Time `json:"created"`
}

func New(slug string) *StandupHistory {
	return &StandupHistory{Slug: slug}
}

func Random() *StandupHistory {
	return &StandupHistory{
		Slug:        util.RandomString(12),
		StandupID:   util.UUID(),
		StandupName: util.RandomString(12),
		Created:     util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*StandupHistory, error) {
	ret := &StandupHistory{}
	var err error
	if setPK {
		ret.Slug, err = m.ParseString("slug", true, true)
		if err != nil {
			return nil, err
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
	ret.StandupName, err = m.ParseString("standupName", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (s *StandupHistory) Clone() *StandupHistory {
	return &StandupHistory{
		Slug:        s.Slug,
		StandupID:   s.StandupID,
		StandupName: s.StandupName,
		Created:     s.Created,
	}
}

func (s *StandupHistory) String() string {
	return s.Slug
}

func (s *StandupHistory) TitleString() string {
	return s.String()
}

func (s *StandupHistory) WebPath() string {
	return "/admin/db/standup/history/" + s.Slug
}

func (s *StandupHistory) Diff(sx *StandupHistory) util.Diffs {
	var diffs util.Diffs
	if s.Slug != sx.Slug {
		diffs = append(diffs, util.NewDiff("slug", s.Slug, sx.Slug))
	}
	if s.StandupID != sx.StandupID {
		diffs = append(diffs, util.NewDiff("standupID", s.StandupID.String(), sx.StandupID.String()))
	}
	if s.StandupName != sx.StandupName {
		diffs = append(diffs, util.NewDiff("standupName", s.StandupName, sx.StandupName))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}

func (s *StandupHistory) ToData() []any {
	return []any{s.Slug, s.StandupID, s.StandupName, s.Created}
}
