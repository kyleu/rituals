// Package vote - Content managed by Project Forge, see [projectforge.md] for details.
package vote

import "github.com/kyleu/rituals/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Vote, error) {
	ret := &Vote{}
	var err error
	if setPK {
		retStoryID, e := m.ParseUUID("storyID", true, true)
		if e != nil {
			return nil, e
		}
		if retStoryID != nil {
			ret.StoryID = *retStoryID
		}
		retUserID, e := m.ParseUUID("userID", true, true)
		if e != nil {
			return nil, e
		}
		if retUserID != nil {
			ret.UserID = *retUserID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Choice, err = m.ParseString("choice", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
