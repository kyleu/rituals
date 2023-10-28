// Package action - Content managed by Project Forge, see [projectforge.md] for details.
package action

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func FromMap(m util.ValueMap, setPK bool) (*Action, error) {
	ret := &Action{}
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
	retSvc, err := m.ParseString("svc", true, true)
	if err != nil {
		return nil, err
	}
	ret.Svc = enum.AllModelServices.Get(retSvc, nil)
	retModelID, e := m.ParseUUID("modelID", true, true)
	if e != nil {
		return nil, e
	}
	if retModelID != nil {
		ret.ModelID = *retModelID
	}
	retUserID, e := m.ParseUUID("userID", true, true)
	if e != nil {
		return nil, e
	}
	if retUserID != nil {
		ret.UserID = *retUserID
	}
	ret.Act, err = m.ParseString("act", true, true)
	if err != nil {
		return nil, err
	}
	ret.Content, err = m.ParseMap("content", true, true)
	if err != nil {
		return nil, err
	}
	ret.Note, err = m.ParseString("note", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
