// Package shistory - Content managed by Project Forge, see [projectforge.md] for details.
package shistory

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type SprintHistory struct {
	Slug       string    `json:"slug,omitempty"`
	SprintID   uuid.UUID `json:"sprintID,omitempty"`
	SprintName string    `json:"sprintName,omitempty"`
	Created    time.Time `json:"created,omitempty"`
}

func New(slug string) *SprintHistory {
	return &SprintHistory{Slug: slug}
}

func Random() *SprintHistory {
	return &SprintHistory{
		Slug:       util.RandomString(12),
		SprintID:   util.UUID(),
		SprintName: util.RandomString(12),
		Created:    util.TimeCurrent(),
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
	return &SprintHistory{s.Slug, s.SprintID, s.SprintName, s.Created}
}

func (s *SprintHistory) String() string {
	return s.Slug
}

func (s *SprintHistory) TitleString() string {
	return s.String()
}

func (s *SprintHistory) WebPath() string {
	return "/admin/db/sprint/history/" + url.QueryEscape(s.Slug)
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
