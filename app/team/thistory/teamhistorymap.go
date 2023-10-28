// Package thistory - Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*TeamHistory, error) {
	ret := &TeamHistory{}
	var err error
	if setPK {
		ret.Slug, err = m.ParseString("slug", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	retTeamID, e := m.ParseUUID("teamID", true, true)
	if e != nil {
		return nil, e
	}
	if retTeamID != nil {
		ret.TeamID = *retTeamID
	}
	ret.TeamName, err = m.ParseString("teamName", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
