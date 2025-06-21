package epermission

import "github.com/kyleu/rituals/app/util"

func (e *EstimatePermission) ToMap() util.ValueMap {
	return util.ValueMap{"estimateID": e.EstimateID, "key": e.Key, "value": e.Value, "access": e.Access, "created": e.Created}
}

func EstimatePermissionFromMap(m util.ValueMap, setPK bool) (*EstimatePermission, util.ValueMap, error) {
	ret := &EstimatePermission{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "estimateID":
			if setPK {
				retEstimateID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retEstimateID != nil {
					ret.EstimateID = *retEstimateID
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
func (e *EstimatePermission) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "estimateID", V: e.EstimateID}, {K: "key", V: e.Key}, {K: "value", V: e.Value}, {K: "access", V: e.Access}, {K: "created", V: e.Created}}
	return util.NewOrderedMap(false, 4, pairs...)
}
