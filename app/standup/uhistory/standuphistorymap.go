// Package uhistory - Content managed by Project Forge, see [projectforge.md] for details.
package uhistory

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*StandupHistory, util.ValueMap, error) {
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
