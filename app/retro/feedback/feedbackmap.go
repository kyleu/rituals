package feedback

import "github.com/kyleu/rituals/app/util"

//nolint:lll
func (f *Feedback) ToMap() util.ValueMap {
	return util.ValueMap{"id": f.ID, "retroID": f.RetroID, "idx": f.Idx, "userID": f.UserID, "category": f.Category, "content": f.Content, "html": f.HTML, "created": f.Created, "updated": f.Updated}
}

func FeedbackFromMap(m util.ValueMap, setPK bool) (*Feedback, util.ValueMap, error) {
	ret := &Feedback{}
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
		case "retroID":
			retRetroID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retRetroID != nil {
				ret.RetroID = *retRetroID
			}
		case "idx":
			ret.Idx, err = m.ParseInt(k, true, true)
		case "userID":
			retUserID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retUserID != nil {
				ret.UserID = *retUserID
			}
		case "category":
			ret.Category, err = m.ParseString(k, true, true)
		case "content":
			ret.Content, err = m.ParseString(k, true, true)
		case "html":
			ret.HTML, err = m.ParseString(k, true, true)
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
