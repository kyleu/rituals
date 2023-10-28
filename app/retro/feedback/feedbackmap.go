// Package feedback - Content managed by Project Forge, see [projectforge.md] for details.
package feedback

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Feedback, error) {
	ret := &Feedback{}
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
	retRetroID, e := m.ParseUUID("retroID", true, true)
	if e != nil {
		return nil, e
	}
	if retRetroID != nil {
		ret.RetroID = *retRetroID
	}
	ret.Idx, err = m.ParseInt("idx", true, true)
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
	ret.Category, err = m.ParseString("category", true, true)
	if err != nil {
		return nil, err
	}
	ret.Content, err = m.ParseString("content", true, true)
	if err != nil {
		return nil, err
	}
	ret.HTML, err = m.ParseString("html", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
