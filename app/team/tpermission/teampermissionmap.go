// Package tpermission - Content managed by Project Forge, see [projectforge.md] for details.
package tpermission

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*TeamPermission, error) {
	ret := &TeamPermission{}
	var err error
	if setPK {
		retTeamID, e := m.ParseUUID("teamID", true, true)
		if e != nil {
			return nil, e
		}
		if retTeamID != nil {
			ret.TeamID = *retTeamID
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
