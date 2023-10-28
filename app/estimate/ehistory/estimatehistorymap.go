// Package ehistory - Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import "github.com/kyleu/rituals/app/util"

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
