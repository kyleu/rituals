// Package spermission - Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*SprintPermission, error) {
	ret := &SprintPermission{}
	var err error
	if setPK {
		retSprintID, e := m.ParseUUID("sprintID", true, true)
		if e != nil {
			return nil, e
		}
		if retSprintID != nil {
			ret.SprintID = *retSprintID
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
