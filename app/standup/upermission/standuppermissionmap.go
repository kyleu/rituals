// Package upermission - Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*StandupPermission, error) {
	ret := &StandupPermission{}
	var err error
	if setPK {
		retStandupID, e := m.ParseUUID("standupID", true, true)
		if e != nil {
			return nil, e
		}
		if retStandupID != nil {
			ret.StandupID = *retStandupID
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
