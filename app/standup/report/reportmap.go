// Package report - Content managed by Project Forge, see [projectforge.md] for details.
package report

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Report, error) {
	ret := &Report{}
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
	retStandupID, e := m.ParseUUID("standupID", true, true)
	if e != nil {
		return nil, e
	}
	if retStandupID != nil {
		ret.StandupID = *retStandupID
	}
	retDay, e := m.ParseTime("day", true, true)
	if e != nil {
		return nil, e
	}
	if retDay != nil {
		ret.Day = *retDay
	}
	retUserID, e := m.ParseUUID("userID", true, true)
	if e != nil {
		return nil, e
	}
	if retUserID != nil {
		ret.UserID = *retUserID
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
