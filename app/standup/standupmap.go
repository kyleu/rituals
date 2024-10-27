package standup

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (s *Standup) ToMap() util.ValueMap {
	return util.ValueMap{"id": s.ID, "slug": s.Slug, "title": s.Title, "icon": s.Icon, "status": s.Status, "teamID": s.TeamID, "sprintID": s.SprintID, "created": s.Created, "updated": s.Updated}
}

func StandupFromMap(m util.ValueMap, setPK bool) (*Standup, util.ValueMap, error) {
	ret := &Standup{}
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
