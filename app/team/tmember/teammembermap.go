package tmember

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func (t *TeamMember) ToMap() util.ValueMap {
	return util.ValueMap{"teamID": t.TeamID, "userID": t.UserID, "name": t.Name, "picture": t.Picture, "role": t.Role, "created": t.Created, "updated": t.Updated}
}

func TeamMemberFromMap(m util.ValueMap, setPK bool) (*TeamMember, util.ValueMap, error) {
	ret := &TeamMember{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "teamID":
			if setPK {
				retTeamID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retTeamID != nil {
					ret.TeamID = *retTeamID
				}
			}
		case "userID":
			if setPK {
				retUserID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retUserID != nil {
					ret.UserID = *retUserID
				}
			}
		case "name":
			ret.Name, err = m.ParseString(k, true, true)
		case "picture":
			ret.Picture, err = m.ParseString(k, true, true)
		case "role":
			retRole, err := m.ParseString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
			ret.Role = enum.AllMemberStatuses.Get(retRole, nil)
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
func (t *TeamMember) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "teamID", V: t.TeamID}, {K: "userID", V: t.UserID}, {K: "name", V: t.Name}, {K: "picture", V: t.Picture}, {K: "role", V: t.Role}, {K: "created", V: t.Created}, {K: "updated", V: t.Updated}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
