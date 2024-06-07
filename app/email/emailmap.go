// Package email - Content managed by Project Forge, see [projectforge.md] for details.
package email

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Email, util.ValueMap, error) {
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
