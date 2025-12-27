package story

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (s *Story) ToMap() util.ValueMap {
	return util.ValueMap{"id": s.ID, "estimateID": s.EstimateID, "idx": s.Idx, "userID": s.UserID, "title": s.Title, "status": s.Status, "finalVote": s.FinalVote, "created": s.Created, "updated": s.Updated}
}

//nolint:gocognit
func StoryFromMap(m util.ValueMap, setPK bool) (*Story, util.ValueMap, error) {
	ret := &Story{}
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
		case "estimateID":
			retEstimateID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retEstimateID != nil {
				ret.EstimateID = *retEstimateID
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
		case "title":
			ret.Title, err = m.ParseString(k, true, true)
		case "status":
			retStatus, err := m.ParseString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
			ret.Status = enum.AllSessionStatuses.Get(retStatus, nil)
		case "finalVote":
			ret.FinalVote, err = m.ParseString(k, true, true)
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
func (s *Story) ToOrderedMap() *util.OrderedMap[any] {
	if s == nil {
		return nil
	}
	pairs := util.OrderedPairs[any]{{K: "id", V: s.ID}, {K: "estimateID", V: s.EstimateID}, {K: "idx", V: s.Idx}, {K: "userID", V: s.UserID}, {K: "title", V: s.Title}, {K: "status", V: s.Status}, {K: "finalVote", V: s.FinalVote}, {K: "created", V: s.Created}, {K: "updated", V: s.Updated}}
	return util.NewOrderedMap(false, 4, pairs...)
}
