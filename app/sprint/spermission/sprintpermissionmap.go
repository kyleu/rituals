package spermission

import "github.com/kyleu/rituals/app/util"

func (s *SprintPermission) ToMap() util.ValueMap {
	return util.ValueMap{"sprintID": s.SprintID, "key": s.Key, "value": s.Value, "access": s.Access, "created": s.Created}
}

func SprintPermissionFromMap(m util.ValueMap, setPK bool) (*SprintPermission, util.ValueMap, error) {
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

//nolint:lll
func (s *SprintPermission) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "sprintID", V: s.SprintID}, {K: "key", V: s.Key}, {K: "value", V: s.Value}, {K: "access", V: s.Access}, {K: "created", V: s.Created}}
	return util.NewOrderedMap(false, 4, pairs...)
}
