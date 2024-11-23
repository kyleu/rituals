package uhistory

import "github.com/kyleu/rituals/app/util"

func (s *StandupHistory) ToMap() util.ValueMap {
	return util.ValueMap{"slug": s.Slug, "standupID": s.StandupID, "standupName": s.StandupName, "created": s.Created}
}

func StandupHistoryFromMap(m util.ValueMap, setPK bool) (*StandupHistory, util.ValueMap, error) {
	ret := &StandupHistory{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "slug":
			if setPK {
				ret.Slug, err = m.ParseString(k, true, true)
			}
		case "standupID":
			retStandupID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retStandupID != nil {
				ret.StandupID = *retStandupID
			}
		case "standupName":
			ret.StandupName, err = m.ParseString(k, true, true)
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

func (s *StandupHistory) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "slug", V: s.Slug}, {K: "standupID", V: s.StandupID}, {K: "standupName", V: s.StandupName}, {K: "created", V: s.Created}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
