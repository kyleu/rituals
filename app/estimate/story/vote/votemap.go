package vote

import "github.com/kyleu/rituals/app/util"

func (v *Vote) ToMap() util.ValueMap {
	return util.ValueMap{"storyID": v.StoryID, "userID": v.UserID, "choice": v.Choice, "created": v.Created, "updated": v.Updated}
}

func VoteFromMap(m util.ValueMap, setPK bool) (*Vote, util.ValueMap, error) {
	ret := &Vote{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "storyID":
			if setPK {
				retStoryID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retStoryID != nil {
					ret.StoryID = *retStoryID
				}
			}
		case "userID":
			if setPK {
				retUserID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retUserID != nil {
					ret.UserID = *retUserID
				}
			}
		case "choice":
			ret.Choice, err = m.ParseString(k, true, true)
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
func (v *Vote) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "storyID", V: v.StoryID}, {K: "userID", V: v.UserID}, {K: "choice", V: v.Choice}, {K: "created", V: v.Created}, {K: "updated", V: v.Updated}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
