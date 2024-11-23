package action

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (a *Action) ToMap() util.ValueMap {
	return util.ValueMap{"id": a.ID, "svc": a.Svc, "modelID": a.ModelID, "userID": a.UserID, "act": a.Act, "content": a.Content, "note": a.Note, "created": a.Created}
}

//nolint:gocognit
func ActionFromMap(m util.ValueMap, setPK bool) (*Action, util.ValueMap, error) {
	ret := &Action{}
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
		case "svc":
			retSvc, err := m.ParseString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
			ret.Svc = enum.AllModelServices.Get(retSvc, nil)
		case "modelID":
			retModelID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retModelID != nil {
				ret.ModelID = *retModelID
			}
		case "userID":
			retUserID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retUserID != nil {
				ret.UserID = *retUserID
			}
		case "act":
			ret.Act, err = m.ParseString(k, true, true)
		case "content":
			ret.Content, err = m.ParseMap(k, true, true)
		case "note":
			ret.Note, err = m.ParseString(k, true, true)
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
func (a *Action) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: a.ID}, {K: "svc", V: a.Svc}, {K: "modelID", V: a.ModelID}, {K: "userID", V: a.UserID}, {K: "act", V: a.Act}, {K: "content", V: a.Content}, {K: "note", V: a.Note}, {K: "created", V: a.Created}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
