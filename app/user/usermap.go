package user

import "github.com/kyleu/rituals/app/util"

func (u *User) ToMap() util.ValueMap {
	return util.ValueMap{"id": u.ID, "name": u.Name, "picture": u.Picture, "created": u.Created, "updated": u.Updated}
}

func UserFromMap(m util.ValueMap, setPK bool) (*User, util.ValueMap, error) {
	ret := &User{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "id":
			if setPK {
				retID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retID != nil {
					ret.ID = *retID
				}
			}
		case "name":
			ret.Name, err = m.ParseString(k, true, true)
		case "picture":
			ret.Picture, err = m.ParseString(k, true, true)
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
func (u *User) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: u.ID}, {K: "name", V: u.Name}, {K: "picture", V: u.Picture}, {K: "created", V: u.Created}, {K: "updated", V: u.Updated}}
	return util.NewOrderedMap(false, 4, pairs...)
}
