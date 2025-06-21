package upermission

import "github.com/kyleu/rituals/app/util"

func (s *StandupPermission) ToMap() util.ValueMap {
	return util.ValueMap{"standupID": s.StandupID, "key": s.Key, "value": s.Value, "access": s.Access, "created": s.Created}
}

func StandupPermissionFromMap(m util.ValueMap, setPK bool) (*StandupPermission, util.ValueMap, error) {
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

//nolint:lll
func (s *StandupPermission) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "standupID", V: s.StandupID}, {K: "key", V: s.Key}, {K: "value", V: s.Value}, {K: "access", V: s.Access}, {K: "created", V: s.Created}}
	return util.NewOrderedMap(false, 4, pairs...)
}
