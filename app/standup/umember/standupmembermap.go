// Package umember - Content managed by Project Forge, see [projectforge.md] for details.
package umember

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func FromMap(m util.ValueMap, setPK bool) (*StandupMember, util.ValueMap, error) {
	ret := &StandupMember{}
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
		case "userID":
			if setPK {
				retUserID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retUserID != nil {
					ret.UserID = *retUserID
				}
			}
		case "name":
			ret.Name, err = m.ParseString(k, true, true)
		case "picture":
			ret.Picture, err = m.ParseString(k, true, true)
		case "role":
			retRole, err := m.ParseString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
			ret.Role = enum.AllMemberStatuses.Get(retRole, nil)
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
