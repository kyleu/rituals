package thistory

import "github.com/kyleu/rituals/app/util"

func (t *TeamHistory) ToMap() util.ValueMap {
	return util.ValueMap{"slug": t.Slug, "teamID": t.TeamID, "teamName": t.TeamName, "created": t.Created}
}

func TeamHistoryFromMap(m util.ValueMap, setPK bool) (*TeamHistory, util.ValueMap, error) {
	ret := &TeamHistory{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "slug":
			if setPK {
				ret.Slug, err = m.ParseString(k, true, true)
			}
		case "teamID":
			retTeamID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retTeamID != nil {
				ret.TeamID = *retTeamID
			}
		case "teamName":
			ret.TeamName, err = m.ParseString(k, true, true)
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

func (t *TeamHistory) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "slug", V: t.Slug}, {K: "teamID", V: t.TeamID}, {K: "teamName", V: t.TeamName}, {K: "created", V: t.Created}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
