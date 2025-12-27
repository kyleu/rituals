package email

import "github.com/kyleu/rituals/app/util"

//nolint:lll
func (e *Email) ToMap() util.ValueMap {
	return util.ValueMap{"id": e.ID, "recipients": e.Recipients, "subject": e.Subject, "data": e.Data, "plain": e.Plain, "html": e.HTML, "userID": e.UserID, "status": e.Status, "created": e.Created}
}

func EmailFromMap(m util.ValueMap, setPK bool) (*Email, util.ValueMap, error) {
	ret := &Email{}
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
		case "recipients":
			ret.Recipients, err = m.ParseArrayString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
		case "subject":
			ret.Subject, err = m.ParseString(k, true, true)
		case "data":
			ret.Data, err = m.ParseMap(k, true, true)
		case "plain":
			ret.Plain, err = m.ParseString(k, true, true)
		case "html":
			ret.HTML, err = m.ParseString(k, true, true)
		case "userID":
			retUserID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retUserID != nil {
				ret.UserID = *retUserID
			}
		case "status":
			ret.Status, err = m.ParseString(k, true, true)
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
func (e *Email) ToOrderedMap() *util.OrderedMap[any] {
	if e == nil {
		return nil
	}
	pairs := util.OrderedPairs[any]{{K: "id", V: e.ID}, {K: "recipients", V: e.Recipients}, {K: "subject", V: e.Subject}, {K: "data", V: e.Data}, {K: "plain", V: e.Plain}, {K: "html", V: e.HTML}, {K: "userID", V: e.UserID}, {K: "status", V: e.Status}, {K: "created", V: e.Created}}
	return util.NewOrderedMap(false, 4, pairs...)
}
