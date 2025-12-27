package rpermission

import "github.com/kyleu/rituals/app/util"

func (r *RetroPermission) ToMap() util.ValueMap {
	return util.ValueMap{"retroID": r.RetroID, "key": r.Key, "value": r.Value, "access": r.Access, "created": r.Created}
}

func RetroPermissionFromMap(m util.ValueMap, setPK bool) (*RetroPermission, util.ValueMap, error) {
	ret := &RetroPermission{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "retroID":
			if setPK {
				retRetroID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retRetroID != nil {
					ret.RetroID = *retRetroID
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
func (r *RetroPermission) ToOrderedMap() *util.OrderedMap[any] {
	if r == nil {
		return nil
	}
	pairs := util.OrderedPairs[any]{{K: "retroID", V: r.RetroID}, {K: "key", V: r.Key}, {K: "value", V: r.Value}, {K: "access", V: r.Access}, {K: "created", V: r.Created}}
	return util.NewOrderedMap(false, 4, pairs...)
}
