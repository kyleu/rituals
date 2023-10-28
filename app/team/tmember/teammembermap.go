// Package tmember - Content managed by Project Forge, see [projectforge.md] for details.
package tmember

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func FromMap(m util.ValueMap, setPK bool) (*TeamMember, error) {
	ret := &TeamMember{}
	var err error
	if setPK {
		retTeamID, e := m.ParseUUID("teamID", true, true)
		if e != nil {
			return nil, e
		}
		if retTeamID != nil {
			ret.TeamID = *retTeamID
		}
		retUserID, e := m.ParseUUID("userID", true, true)
		if e != nil {
			return nil, e
		}
		if retUserID != nil {
			ret.UserID = *retUserID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	ret.Picture, err = m.ParseString("picture", true, true)
	if err != nil {
		return nil, err
	}
	retRole, err := m.ParseString("role", true, true)
	if err != nil {
		return nil, err
	}
	ret.Role = enum.AllMemberStatuses.Get(retRole, nil)
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
