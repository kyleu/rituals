package rhistory

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*RetroHistory, util.ValueMap, error) {
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
