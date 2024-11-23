package ehistory

import "github.com/kyleu/rituals/app/util"

func (e *EstimateHistory) ToMap() util.ValueMap {
	return util.ValueMap{"slug": e.Slug, "estimateID": e.EstimateID, "estimateName": e.EstimateName, "created": e.Created}
}

func EstimateHistoryFromMap(m util.ValueMap, setPK bool) (*EstimateHistory, util.ValueMap, error) {
	ret := &EstimateHistory{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "slug":
			if setPK {
				ret.Slug, err = m.ParseString(k, true, true)
			}
		case "estimateID":
			retEstimateID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retEstimateID != nil {
				ret.EstimateID = *retEstimateID
			}
		case "estimateName":
			ret.EstimateName, err = m.ParseString(k, true, true)
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
func (e *EstimateHistory) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "slug", V: e.Slug}, {K: "estimateID", V: e.EstimateID}, {K: "estimateName", V: e.EstimateName}, {K: "created", V: e.Created}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
