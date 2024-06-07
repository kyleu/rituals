// Package thistory - Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*TeamHistory, util.ValueMap, error) {
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
