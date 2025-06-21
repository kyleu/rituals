package rhistory

import "github.com/kyleu/rituals/app/util"

func (r *RetroHistory) ToMap() util.ValueMap {
	return util.ValueMap{"slug": r.Slug, "retroID": r.RetroID, "retroName": r.RetroName, "created": r.Created}
}

func RetroHistoryFromMap(m util.ValueMap, setPK bool) (*RetroHistory, util.ValueMap, error) {
	ret := &RetroHistory{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "slug":
			if setPK {
				ret.Slug, err = m.ParseString(k, true, true)
			}
		case "retroID":
			retRetroID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retRetroID != nil {
				ret.RetroID = *retRetroID
			}
		case "retroName":
			ret.RetroName, err = m.ParseString(k, true, true)
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

func (r *RetroHistory) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "slug", V: r.Slug}, {K: "retroID", V: r.RetroID}, {K: "retroName", V: r.RetroName}, {K: "created", V: r.Created}}
	return util.NewOrderedMap(false, 4, pairs...)
}
