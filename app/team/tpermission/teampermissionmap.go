package tpermission

import "github.com/kyleu/rituals/app/util"

func (t *TeamPermission) ToMap() util.ValueMap {
	return util.ValueMap{"teamID": t.TeamID, "key": t.Key, "value": t.Value, "access": t.Access, "created": t.Created}
}

func TeamPermissionFromMap(m util.ValueMap, setPK bool) (*TeamPermission, util.ValueMap, error) {
	ret := &TeamPermission{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "teamID":
			if setPK {
				retTeamID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retTeamID != nil {
					ret.TeamID = *retTeamID
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
