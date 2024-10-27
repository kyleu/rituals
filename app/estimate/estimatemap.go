package estimate

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (e *Estimate) ToMap() util.ValueMap {
	return util.ValueMap{"id": e.ID, "slug": e.Slug, "title": e.Title, "icon": e.Icon, "status": e.Status, "teamID": e.TeamID, "sprintID": e.SprintID, "choices": e.Choices, "created": e.Created, "updated": e.Updated}
}

func EstimateFromMap(m util.ValueMap, setPK bool) (*Estimate, util.ValueMap, error) {
	ret := &Estimate{}
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
		case "slug":
			ret.Slug, err = m.ParseString(k, true, true)
		case "title":
			ret.Title, err = m.ParseString(k, true, true)
		case "icon":
			ret.Icon, err = m.ParseString(k, true, true)
		case "status":
			retStatus, err := m.ParseString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
			ret.Status = enum.AllSessionStatuses.Get(retStatus, nil)
		case "teamID":
			ret.TeamID, err = m.ParseUUID(k, true, true)
		case "sprintID":
			ret.SprintID, err = m.ParseUUID(k, true, true)
		case "choices":
			ret.Choices, err = m.ParseArrayString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
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
