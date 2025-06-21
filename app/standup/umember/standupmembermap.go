package umember

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (s *StandupMember) ToMap() util.ValueMap {
	return util.ValueMap{"standupID": s.StandupID, "userID": s.UserID, "name": s.Name, "picture": s.Picture, "role": s.Role, "created": s.Created, "updated": s.Updated}
}

func StandupMemberFromMap(m util.ValueMap, setPK bool) (*StandupMember, util.ValueMap, error) {
	ret := &StandupMember{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "standupID":
			if setPK {
				retStandupID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retStandupID != nil {
					ret.StandupID = *retStandupID
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
func (s *StandupMember) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "standupID", V: s.StandupID}, {K: "userID", V: s.UserID}, {K: "name", V: s.Name}, {K: "picture", V: s.Picture}, {K: "role", V: s.Role}, {K: "created", V: s.Created}, {K: "updated", V: s.Updated}}
	return util.NewOrderedMap(false, 4, pairs...)
}
