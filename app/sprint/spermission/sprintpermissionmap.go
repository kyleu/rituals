// Package spermission - Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*SprintPermission, util.ValueMap, error) {
	ret := &SprintPermission{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "sprintID":
			if setPK {
				retSprintID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retSprintID != nil {
					ret.SprintID = *retSprintID
				}
			}
		case "key":
			if setPK {
				ret.Key, err = m.ParseString(k, true, true)
			}
		case "value":
			if setPK {
				ret.Value, err = m.ParseString(k, true, true)
			}
		case "access":
			ret.Access, err = m.ParseString(k, true, true)
		default:
			extra[k] = v
		}
		if err != nil {
			return nil, nil, err
		}
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, extra, nil
}
