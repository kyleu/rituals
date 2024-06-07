// Package story - Content managed by Project Forge, see [projectforge.md] for details.
package story

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

//nolint:gocognit
func FromMap(m util.ValueMap, setPK bool) (*Story, util.ValueMap, error) {
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
