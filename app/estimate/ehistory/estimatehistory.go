// Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type EstimateHistory struct {
	Slug         string    `json:"slug"`
	EstimateID   uuid.UUID `json:"estimateID"`
	EstimateName string    `json:"estimateName"`
	Created      time.Time `json:"created"`
}

func New(slug string) *EstimateHistory {
	return &EstimateHistory{Slug: slug}
}

func Random() *EstimateHistory {
	return &EstimateHistory{
		Slug:         util.RandomString(12),
		EstimateID:   util.UUID(),
		EstimateName: util.RandomString(12),
		Created:      time.Now(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*EstimateHistory, error) {
	ret := &EstimateHistory{}
	var err error
	if setPK {
		ret.Slug, err = m.ParseString("slug", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	retEstimateID, e := m.ParseUUID("estimateID", true, true)
	if e != nil {
		return nil, e
	}
	if retEstimateID != nil {
		ret.EstimateID = *retEstimateID
	}
	ret.EstimateName, err = m.ParseString("estimateName", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (e *EstimateHistory) Clone() *EstimateHistory {
	return &EstimateHistory{
		Slug:         e.Slug,
		EstimateID:   e.EstimateID,
		EstimateName: e.EstimateName,
		Created:      e.Created,
	}
}

func (e *EstimateHistory) String() string {
	return e.Slug
}

func (e *EstimateHistory) TitleString() string {
	return e.String()
}

func (e *EstimateHistory) WebPath() string {
	return "/admin/db/estimate/history" + "/" + e.Slug
}

func (e *EstimateHistory) Diff(ex *EstimateHistory) util.Diffs {
	var diffs util.Diffs
	if e.Slug != ex.Slug {
		diffs = append(diffs, util.NewDiff("slug", e.Slug, ex.Slug))
	}
	if e.EstimateID != ex.EstimateID {
		diffs = append(diffs, util.NewDiff("estimateID", e.EstimateID.String(), ex.EstimateID.String()))
	}
	if e.EstimateName != ex.EstimateName {
		diffs = append(diffs, util.NewDiff("estimateName", e.EstimateName, ex.EstimateName))
	}
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}

func (e *EstimateHistory) ToData() []any {
	return []any{e.Slug, e.EstimateID, e.EstimateName, e.Created}
}
