// Package epermission - Content managed by Project Forge, see [projectforge.md] for details.
package epermission

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*EstimatePermission, error) {
	ret := &EstimatePermission{}
	var err error
	if setPK {
		retEstimateID, e := m.ParseUUID("estimateID", true, true)
		if e != nil {
			return nil, e
		}
		if retEstimateID != nil {
			ret.EstimateID = *retEstimateID
		}
		ret.Key, err = m.ParseString("key", true, true)
		if err != nil {
			return nil, err
		}
		ret.Value, err = m.ParseString("value", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Access, err = m.ParseString("access", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
