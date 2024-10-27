package emember

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (e *EstimateMember) ToMap() util.ValueMap {
	return util.ValueMap{"estimateID": e.EstimateID, "userID": e.UserID, "name": e.Name, "picture": e.Picture, "role": e.Role, "created": e.Created, "updated": e.Updated}
}

func EstimateMemberFromMap(m util.ValueMap, setPK bool) (*EstimateMember, util.ValueMap, error) {
	ret := &EstimateMember{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "estimateID":
			if setPK {
				retEstimateID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retEstimateID != nil {
					ret.EstimateID = *retEstimateID
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
