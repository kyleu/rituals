package shistory

import "github.com/kyleu/rituals/app/util"

func (s *SprintHistory) ToMap() util.ValueMap {
	return util.ValueMap{"slug": s.Slug, "sprintID": s.SprintID, "sprintName": s.SprintName, "created": s.Created}
}

func SprintHistoryFromMap(m util.ValueMap, setPK bool) (*SprintHistory, util.ValueMap, error) {
	ret := &SprintHistory{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "slug":
			if setPK {
				ret.Slug, err = m.ParseString(k, true, true)
			}
		case "sprintID":
			retSprintID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retSprintID != nil {
				ret.SprintID = *retSprintID
			}
		case "sprintName":
			ret.SprintName, err = m.ParseString(k, true, true)
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

func (s *SprintHistory) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "slug", V: s.Slug}, {K: "sprintID", V: s.SprintID}, {K: "sprintName", V: s.SprintName}, {K: "created", V: s.Created}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
