// Package rhistory - Content managed by Project Forge, see [projectforge.md] for details.
package rhistory

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*RetroHistory, error) {
	ret := &RetroHistory{}
	var err error
	if setPK {
		ret.Slug, err = m.ParseString("slug", true, true)
		if err != nil {
			return nil, err
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
	ret.RetroName, err = m.ParseString("retroName", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
