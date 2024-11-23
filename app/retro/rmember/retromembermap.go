package rmember

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (r *RetroMember) ToMap() util.ValueMap {
	return util.ValueMap{"retroID": r.RetroID, "userID": r.UserID, "name": r.Name, "picture": r.Picture, "role": r.Role, "created": r.Created, "updated": r.Updated}
}

func RetroMemberFromMap(m util.ValueMap, setPK bool) (*RetroMember, util.ValueMap, error) {
	ret := &RetroMember{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "retroID":
			if setPK {
				retRetroID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retRetroID != nil {
					ret.RetroID = *retRetroID
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
func (r *RetroMember) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "retroID", V: r.RetroID}, {K: "userID", V: r.UserID}, {K: "name", V: r.Name}, {K: "picture", V: r.Picture}, {K: "role", V: r.Role}, {K: "created", V: r.Created}, {K: "updated", V: r.Updated}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
