// Package upermission - Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*StandupPermission, util.ValueMap, error) {
	ret := &StandupPermission{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "standupID":
			if setPK {
				retStandupID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retStandupID != nil {
					ret.StandupID = *retStandupID
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
