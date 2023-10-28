// Package uhistory - Content managed by Project Forge, see [projectforge.md] for details.
package uhistory

import "github.com/kyleu/rituals/app/util"

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
