// Package shistory - Content managed by Project Forge, see [projectforge.md] for details.
package shistory

import "github.com/kyleu/rituals/app/util"

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
