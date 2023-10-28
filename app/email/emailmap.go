// Package email - Content managed by Project Forge, see [projectforge.md] for details.
package email

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Email, error) {
	ret := &Email{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Recipients, err = m.ParseArrayString("recipients", true, true)
	if err != nil {
		return nil, err
	}
	ret.Subject, err = m.ParseString("subject", true, true)
	if err != nil {
		return nil, err
	}
	ret.Data, err = m.ParseMap("data", true, true)
	if err != nil {
		return nil, err
	}
	ret.Plain, err = m.ParseString("plain", true, true)
	if err != nil {
		return nil, err
	}
	ret.HTML, err = m.ParseString("html", true, true)
	if err != nil {
		return nil, err
	}
	retUserID, e := m.ParseUUID("userID", true, true)
	if e != nil {
		return nil, e
	}
	if retUserID != nil {
		ret.UserID = *retUserID
	}
	ret.Status, err = m.ParseString("status", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
